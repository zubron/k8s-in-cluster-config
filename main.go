package main

import (
	"fmt"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/restmapper"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetRESTConfig() (*rest.Config, error) {
	cfgLoadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(cfgLoadingRules, configOverrides)
	return kubeConfig.ClientConfig()
}

func main() {
	config, err := GetRESTConfig()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("config.Host", config.Host)

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("could not create discovery client", err)
		return
	}

	discover, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		fmt.Println("could not create discovery client", err)
		return
	}

	groupResources, err := restmapper.GetAPIGroupResources(discover)
	if err != nil {
		fmt.Println("could not get api group resources", err)
		return
	}

	fmt.Printf("found %v API group resources\n", len(groupResources))


	ns := os.Getenv("POD_NAMESPACE")
	pods, err := client.CoreV1().Pods(ns).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the %v namespace:\n", len(pods.Items), ns)
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
}
