package k8s

import (
	"fmt"

	"github.com/julz/cube/opi"
	ext "k8s.io/api/extensions/v1beta1"
	av1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

const (
	INGRESS_NAME        = "eirini"
	INGRESS_API_VERSION = "extensions/v1beta1"
)

type IngressManager struct {
	client   kubernetes.Interface
	endpoint string
}

func NewIngressManager(client kubernetes.Interface, kubeEndpoint string) *IngressManager {
	return &IngressManager{
		client:   client,
		endpoint: kubeEndpoint,
	}
}

func (i *IngressManager) UpdateIngress(lrp opi.LRP, vcap VcapApp) error {
	ingress, err := i.client.ExtensionsV1beta1().Ingresses("default").Get("eririni", av1.GetOptions{})
	if err != nil {
		//TODO: CreateIngress
		return err
	}
	//
	ingress.Spec.TLS[0].Hosts = append(ingress.Spec.TLS[0].Hosts, fmt.Sprintf("%s.%s", vcap.AppName, i.endpoint))
	rule := createIngressRule(lrp, vcap, i.endpoint)
	ingress.Spec.Rules = append(ingress.Spec.Rules, rule)

	if _, err = i.client.ExtensionsV1beta1().Ingresses("default").Update(ingress); err != nil {
		return err
	}

	return nil
}

func createIngressRule(lrp opi.LRP, vcap VcapApp, kubeEndpoint string) ext.IngressRule {
	rule := ext.IngressRule{
		Host: fmt.Sprintf("%s.%s", vcap.AppName, kubeEndpoint),
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

func (i *IngressManager) CreateIngress(namespace string) error {
	ingress := &ext.Ingress{}

	ingress.APIVersion = INGRESS_API_VERSION
	ingress.Kind = "Ingress"
	ingress.Name = INGRESS_NAME
	ingress.Namespace = namespace

	_, err := i.client.ExtensionsV1beta1().Ingresses(namespace).Create(ingress)
	if err != nil {
		return err
	}

	return nil
}
