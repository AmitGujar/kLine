package namespaces

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func CheckoutNamespace(clientset *kubernetes.Clientset, config *clientcmdapi.Config, namespace string) {
	// Check if the namespace exists
	_, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
		// If the namespace doesn't exist, create it
		_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &v1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: namespace,
			},
		}, metav1.CreateOptions{})
		fmt.Println("Namespace created...")
		if err != nil {
			log.Fatalf("Failed to create namespace: %v", err)
		}
	}

	// Change the namespace of the current context
	context := config.Contexts[config.CurrentContext]
	if context == nil {
		context = clientcmdapi.NewContext()
		config.Contexts[config.CurrentContext] = context
	}
	context.Namespace = namespace

	// Save the changes back to the kubeconfig file
	kubeconfig := clientcmd.NewDefaultPathOptions()
	if err := clientcmd.WriteToFile(*config, kubeconfig.GetDefaultFilename()); err != nil {
		log.Fatalf("Failed to save kubeconfig: %v", err)
	}

	fmt.Printf("✅ Switched to namespace => %s\n", namespace)
}
