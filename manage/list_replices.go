package manage

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ListReplicas use for listreplice
func ListReplicas(oneNameSpace string, clientset *kubernetes.Clientset) {

	// clientset := c.Getk8sclient()

	fmt.Printf("Listing deployments in namespace %s:\n", oneNameSpace)
	// deploymentNamespace := "liang-namespace01"
	deploymentsClient := clientset.AppsV1().Deployments(oneNameSpace)
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	}

}
