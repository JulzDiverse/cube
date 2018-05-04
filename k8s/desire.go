package k8s

import (
	"context"

	"github.com/julz/cube/launcher"
	"github.com/julz/cube/opi"
	"k8s.io/api/apps/v1beta1"
	"k8s.io/api/core/v1"
	av1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Desirer struct {
	KubeNamespace string
	Client        *kubernetes.Clientset
}

func (d *Desirer) Desire(ctx context.Context, lrps []opi.LRP) error {
	deployments, err := d.Client.AppsV1beta1().Deployments(d.KubeNamespace).List(av1.ListOptions{})
	if err != nil {
		return err
	}

	dByName := make(map[string]struct{})
	for _, d := range deployments.Items {
		dByName[d.Name] = struct{}{}
	}

	for _, lrp := range lrps {
		if _, ok := dByName[lrp.Name]; ok {
			continue
		}

		if _, err := d.Client.AppsV1beta1().Deployments(d.KubeNamespace).Create(toDeployment(lrp)); err != nil {
			// fixme: this should be a multi-error and deferred
			return err
		}

		if _, err = d.Client.CoreV1().Services(d.KubeNamespace).Create(exposeDeployment(lrp, d.KubeNamespace)); err != nil {
			return err
		}
	}

	return nil
}

func toDeployment(lrp opi.LRP) *v1beta1.Deployment {
	environment := launcher.SetupEnv(lrp.Command[0])
	deployment := &v1beta1.Deployment{
		Spec: v1beta1.DeploymentSpec{
			Replicas: int32ptr(lrp.TargetInstances),
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{{
						Name:  "web",
						Image: lrp.Image,
						Command: []string{
							launcher.Launch,
						},
						Env: mapToEnvVar(mergeMaps(lrp.Env, environment)),
						Ports: []v1.ContainerPort{
							v1.ContainerPort{
								Name:          "expose",
								ContainerPort: 8080,
							},
						},
					}},
				},
			},
		},
	}

	deployment.Name = lrp.Name
	deployment.Spec.Template.Labels = map[string]string{
		"name": lrp.Name,
	}

	deployment.Labels = map[string]string{
		"cube": "cube",
		"name": lrp.Name,
	}

	return deployment
}

func exposeDeployment(lrp opi.LRP, namespace string) *v1.Service {
	service := &v1.Service{
		Spec: v1.ServiceSpec{
			ExternalTrafficPolicy: "Cluster",
			Ports: []v1.ServicePort{
				v1.ServicePort{
					Port:     8080,
					Protocol: v1.ProtocolTCP,
				},
			},
			Selector: map[string]string{
				"name": lrp.Name,
			},
			SessionAffinity: "None",
			Type:            "NodePort",
		},
		Status: v1.ServiceStatus{
			LoadBalancer: v1.LoadBalancerStatus{},
		},
	}

	service.APIVersion = "v1"
	service.Kind = "Service"
	service.Name = lrp.Name
	service.Namespace = namespace
	service.Labels = map[string]string{
		"cube": "cube",
		"name": lrp.Name,
	}

	return service
}

func mergeMaps(maps ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func int32ptr(i int) *int32 {
	u := int32(i)
	return &u
}

func int64ptr(i int) *int64 {
	u := int64(i)
	return &u
}
