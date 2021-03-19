package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/golang/glog"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
}
func kTest() {
	home := homedir.HomeDir()
	configPath := flag.String("kubeconfig", home+"/.kube/config", "kubeconfig path")
	flag.Parse()
	defer glog.Flush()
	config, e := clientcmd.BuildConfigFromFlags("", *configPath)
	if e != nil {
		glog.Fatal(e)
	}
	clientSet, e := kubernetes.NewForConfig(config)
	if e != nil {
		glog.Fatal(e)
	}
	nodeList, e := clientSet.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	if e != nil {
		glog.Fatal(e)
	}
	for _, n := range nodeList.Items {
		fmt.Println(n.Name)
	}
	podList, e := clientSet.CoreV1().Pods("kube-system").List(context.TODO(), v1.ListOptions{})
	if e != nil {
		glog.Fatal(e)
	}
	for _, p := range podList.Items {
		fmt.Println(p.Name)
	}
}
