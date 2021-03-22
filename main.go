package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	//kTest()
	go runServer()
	select {
	
	}
}
func runServer() {
	flag.Parse()
	defer glog.Flush()
	glog.Info("pass")
	hostname, _ := os.Hostname()
	eng := gin.Default()
	eng.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":   "ok",
			"hostname": hostname,
		})
	})
	eng.Run(":80")
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

	podList, e := clientSet.CoreV1().Pods("default").List(context.TODO(), v1.ListOptions{})
	if e != nil {
		glog.Fatal(e)
	}
	for _, p := range podList.Items {
		fmt.Println(p.Name)
	}
}
