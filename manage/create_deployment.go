package manage

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"time"

	extendservice "kubernetes-client/service"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// CreateDeployment use for create deployment
func CreateDeployment(deploymentNamespace string) {

	// clientset := c.getk8sclient()
	// "kubernetes-client/Client"

	var kubeconfig = flag.String("kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) absolute path to the kubeconfig file")

	var config, _ = clientcmd.BuildConfigFromFlags("", *kubeconfig)

	var clientset, _ = kubernetes.NewForConfig(config)

	// deploymentNamespace := "liang-namespace01"

	deploymentsClient := clientset.AppsV1().Deployments(deploymentNamespace)
	// 这个过程中会把包含RESTClient配置信息、命名空间信息赋值到deploymentsClient中
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "use-client-go-demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: extendservice.Int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// Create Deployment
	fmt.Println("Creating deployment...")
	// spew.Dump(deploymentsClient)
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	time.Sleep(30 * time.Second)

}
