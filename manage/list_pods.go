package manage

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ListPods use list pods
func ListPods(clientset *kubernetes.Clientset) {

	// clientset := c.Getk8sclient()

	// var kubeconfig = flag.String("kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) absolute path to the kubeconfig file")

	// var config, _ = clientcmd.BuildConfigFromFlags("", *kubeconfig)

	// var clientset, _ = kubernetes.NewForConfig(config)

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

}
