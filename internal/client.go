package internal

import (
	"os"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/tools/clientcmd"
)

var postgresGVR = schema.GroupVersionResource{
	Group:    "acid.zalan.do",
	Version:  "v1",
	Resource: "postgresqls",
}

func GetDynamicClient(isTest bool) (dynamic.Interface, error) {
	if isTest{
		return getDynamicClient_Test()
	}
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = os.ExpandEnv("$HOME/.kube/config")
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	return dynamic.NewForConfig(config)
}

func getDynamicClient_Test() (dynamic.Interface, error) {
	scheme := runtime.NewScheme()

	cluster1 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "postgres-operator.crunchydata.com/v1beta1",
			"kind":       "PostgresCluster",
			"metadata": map[string]interface{}{
				"name":      "pg-cluster-1",
				"namespace": "default",
			},
			"status": map[string]interface{}{
				"PostgresClusterStatus": "Running",
			},
		},
	}

	cluster2 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "postgres-operator.crunchydata.com/v1beta1",
			"kind":       "PostgresCluster",
			"metadata": map[string]interface{}{
				"name":      "pg-cluster-2",
				"namespace": "dev",
			},
			"status": map[string]interface{}{
				"PostgresClusterStatus": "Pending",
			},
		},
	}

	client := fake.NewSimpleDynamicClient(scheme, cluster1, cluster2)
	return client, nil
}
