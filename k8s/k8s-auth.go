package k8s

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

var kubeconfig *string

func K8SAuth() *kubernetes.Clientset {

	kuberconfig := getKubeconfig()

	config, err := clientcmd.BuildConfigFromFlags("", *kuberconfig)
	if err != nil {
		clientset := inClusterAuth()
		return clientset
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		clientset := inClusterAuth()
		return clientset
	}
	return clientset

}

func init() {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
}

func getKubeconfig() *string {
	return kubeconfig
}

func inClusterAuth() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}
