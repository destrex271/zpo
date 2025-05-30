package internal

import (
	"testing"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/fake"
)

func GetDynamicClient_Test() (dynamic.Interface, error) {
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

func TestListPostgresqlClusters(t *testing.T) {
	err := ListPostgresqlClusters("")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

