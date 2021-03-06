package client

import (
	"flag"
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func Getk8sclient() *kubernetes.Clientset {

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		//flag包提供了一系列解析命令行参数的功能接口,flag.String()该方式返回一个相应的指针
		//下面的意思是： 获取kubeconfig的flag的值，默认的值是filepath.Join(home, ".kube", "config")
		// Join 将任意数量的路径元素连接到单个路径中，如有必要添加分隔符。加入调用清理结果; 特别是，所有空串都被忽略。。有时，我们需要分割 PATH 或 GOPATH 之类的环境变量（这些路径被特定于 OS 的列表分隔符连接起来），filepath.SplitList 就是这个用途
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		fmt.Printf("use kubeconf :%v to init kubernetes client file  ", *kubeconfig)
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}
