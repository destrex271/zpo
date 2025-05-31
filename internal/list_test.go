package internal

import (
	"testing"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/fake"
)



func TestListPostgresqlClusters(t *testing.T) {
	err := ListPostgresqlClusters("", true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

