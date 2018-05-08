package k8s

import (
	"fmt"

	"github.com/julz/cube/opi"
	ext "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	INGRESS_NAME        = "eirini"
	INGRESS_API_VERSION = "extensions/v1beta1"
)

type IngressController struct {
	Client *kubernetes.Clientset
}

func (i *IngressController) UpdateIngress(lrp opi.LRP, vcap VcapApp) error {
	if ingress, err = d.Client.ExtensionsV1beta1().Ingresses("default").Get("eririni", av1.GetOptions{}); err != nil {
		return err //TODO: if ingress not found create it
	}

	ingress.Spec.TLS[0].Hosts = append(ing.Spec.TLS[0].Hosts, fmt.Sprintf("%s.%s", vcap.AppName, "cube-kube.uk-south.containers.mybluemix.net")) //TODO parameterize and name TLS array
	rule := createIngressRule(vcap)
	ingress.Spec.Rules = append(ingress.Spec.Rules, rule)

	if _, err = d.Client.ExtensionsV1beta1().Ingresses("default").Update(ingress); err != nil {
		return err
	}

	return nil
}

func _createIngressRule(vcap VcapApp) ext.IngressRule {
	rule := ext.IngressRule{
		Host: fmt.Sprintf("%s.%s", vcap.AppName, "cube-kube.uk-south.containers.mybluemix.net"), //TODO parameterize
	}

	rule.HTTP = &ext.HTTPIngressRuleValue{
		Paths: []ext.HTTPIngressPath{
			ext.HTTPIngressPath{
				Path: "/",
				Backend: ext.IngressBackend{
					ServiceName: fmt.Sprintf("cf-%s", lrp.Name),
					ServicePort: intstr.FromInt(8080),
				},
			},
		},
	}

	return rule
}

func (i *IngressController) CreateIngress(namespace string) *ext.Ingress {
	ingress := &ext.Ingress{}

	ingress.APIVersion = INGRESS_NAME
	ingress.Kind = "Ingress"
	ingress.Name = INGRESS_API_VERSION
	ingress.Namespace = namespace

	return ingress
}
