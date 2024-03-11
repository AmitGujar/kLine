package control

import (
	"flag"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	// "os"
)

// ...

func ParseFlags() (string, string, string) {
	flag.Parse()
	command := flag.Arg(0)
	namespace := flag.Arg(1)
	clusterName := flag.Arg(1)
	return command, namespace, clusterName
}

func CreateClientset() (*kubernetes.Clientset, *clientcmdapi.Config) {
	kubeconfig := clientcmd.NewDefaultPathOptions()

	// Load the kubeconfig file
	config, err := kubeconfig.GetStartingConfig()
	if err != nil {
		log.Fatalf("Failed to load kubeconfig: %v", err)
	}

	// Create a *rest.Config object
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig.GetDefaultFilename())
	if err != nil {
		log.Fatalf("Failed to build config: %v", err)
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("Failed to create clientset: %v", err)
	}

	return clientset, config
}
