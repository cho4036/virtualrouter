/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"os/exec"
	"time"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"

	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"github.com/cho4036/virtualrouter/executor/iptables"
	"github.com/cho4036/virtualrouter/iptablescontroller"
	clientset "github.com/cho4036/virtualrouter/pkg/client/clientset/versioned"
	informers "github.com/cho4036/virtualrouter/pkg/client/informers/externalversions"
	"github.com/cho4036/virtualrouter/pkg/signals"
	"github.com/vishvananda/netlink"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	// ToDo: Make it variable not constant
	var fwmark uint32 = 200
	// ToDo: Make it variable not constant
	intif := "ethint"
	// ToDo: Make it variable not constant
	extif := "ethext"

	// ToDo: Make it start after interface job done(CNI: intif, extif or signal)
	for {
		if _, err := netlink.LinkByName(extif); err == nil {
			if _, err := netlink.LinkByName(intif); err == nil {
				break
			}
		}
		time.Sleep(time.Second * 1)
	}

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	nfvClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %s", err.Error())
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	nfvInformerFactory := informers.NewSharedInformerFactory(nfvClient, time.Second*30)

	cmd := exec.Command("sysctl", "-p")
	if err := cmd.Run(); err != nil {
		klog.ErrorS(err, "Apply sysctl failed", "command", "sysctl -p")
		return
	}

	controller := NewController(kubeClient, nfvClient,
		nfvInformerFactory.Tmax().V1().NATRules(),
		nfvInformerFactory.Tmax().V1().FireWallRules(),
		nfvInformerFactory.Tmax().V1().LoadBalancerRules(),
		iptablescontroller.New(intif, extif, fwmark, iptables.NewIPV4(), 10*time.Second))

	// notice that there is no need to run Start methods in a separate goroutine. (i.e. go kubeInformerFactory.Start(stopCh)
	// Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	kubeInformerFactory.Start(stopCh)
	nfvInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
