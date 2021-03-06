package manage

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// DeleteDeploy use for delete deployment
func DeleteDeploy(deploymentNamespace, deploymentName string) {

	var kubeconfig = flag.String("kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) absolute path to the kubeconfig file")

	var config, _ = clientcmd.BuildConfigFromFlags("", *kubeconfig)

	var clientset, _ = kubernetes.NewForConfig(config)

	fmt.Println("Deleting deployment...")
	deploymentsClient := clientset.AppsV1().Deployments(deploymentNamespace)
	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentsClient.Delete(context.TODO(), deploymentName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	fmt.Println("Deleted deployment.")
}
