package main

import (
	"fmt"
        "flag"
        "k8s.io/client-go/tools/clientcmd"
        "k8s.io/client-go/kubernetes"
)

func main() {

	kubeconfig := flag.String("kubeconfig", "~/.kube/config", "Location of kubeconfig file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		fmt.Println("Got error %s from  clientcmd.BuildConfigFromFlags() ", log.Fatal(err))
		exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Println("Got error %s from  NewForConfig() ", log.Fatal(err))
		exit(1)
	}

	fmt.Println("got clientset as \n %s \n", clientset)

}
