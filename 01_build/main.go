package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	buildv1 "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".crc", "cache", "crc_libvirt_4.1.9", "kubeconfig"), "(optional) absolute path to kubeconfig")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to kubeconfig")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	buildv1Client, err := buildv1.NewForConfig(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}

	ns := "test-namespace"
	builds, err := buildv1Client.Builds(ns).List(metav1.ListOptions{})

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("There are %d builds in project %s\n", len(builds.Items), ns)

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}
