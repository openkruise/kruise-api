package main

import (
	"context"
	"flag"
	"fmt"
	appsv1alpha1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	kruiseclient "github.com/openkruise/kruise-api/client/clientset/versioned"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kruiseclient.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	client := clientset.AppsV1alpha1().CloneSets("default")

	rollout := appsv1alpha1.CloneSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-rollout",
		},

		Spec: appsv1alpha1.CloneSetSpec{
			Replicas: int32Ptr(2),
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
	result, err := client.Create(context.TODO(), &rollout, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created rollout %q.\n", result.GetObjectMeta().GetName())

}

func int32Ptr(i int32) *int32 { return &i }
