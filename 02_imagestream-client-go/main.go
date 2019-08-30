package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	isV1 "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	home := os.Getenv("HOME")
	kubeconfigPath := filepath.Join(home, ".crc", "cache", "crc_libvirt_4.1.9", "kubeconfig")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatal(err)
	}
	isClient, err := isV1.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	imStrs, err := isClient.ImageStreams("openshift").List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for i, is := range imStrs.Items {
		fmt.Println(i+1, is.Name)
	}
}
