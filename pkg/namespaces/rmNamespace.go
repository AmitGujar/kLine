package namespaces

import (
	"context"
	"fmt"

	"github.com/AmitGujar/kLine/pkg/callbacks"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func DeleteNamespace(clientset *kubernetes.Clientset, namespace string) {
	err := clientset.CoreV1().Namespaces().Delete(context.TODO(), namespace, metav1.DeleteOptions{})
	if err != nil {
		callbacks.PrintErrorMessage(err)
	}
	fmt.Printf("❌ Deleted namespace => %s\n", namespace)
}
