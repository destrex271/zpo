package internal

import (
	"context"
	"fmt"

	"sigs.k8s.io/yaml"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DescribePostgresqlCluster(namespace, name, output string, istest bool) error {
	client, err := GetDynamicClient(false)
	if err != nil {
		return err
	}

	resource, err := client.Resource(postgresGVR).Namespace(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	jsonBytes, err := resource.MarshalJSON()
	if err != nil {
		return err
	}

	switch output {
	case "json":
		fmt.Println(string(jsonBytes))
	case "yaml":
		yamlBytes, err := yaml.JSONToYAML(jsonBytes)
		if err != nil {
			return err
		}
		fmt.Println(string(yamlBytes))
	default:
		return fmt.Errorf("unsupported output format: %s", output)
	}
	return nil
}

