package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"k8s.io/client-go/1.5/pkg/util/homedir"
	"k8s.io/client-go/1.5/pkg/util/yaml"
)

func main() {
	s := homedir.HomeDir()
	fmt.Println(s)
	b, e := ioutil.ReadFile("test.yaml")
	if e != nil {
		log.Fatal(e)
	}
	if new, e := yaml.ToJSON(b); e != nil {
		log.Fatal(e)
	} else {
		fmt.Println(string(new))
	}
}
