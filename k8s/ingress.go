package k8s

import (
	"fmt"

	"github.com/julz/cube/opi"
	ext "k8s.io/api/extensions/v1beta1"
	typed "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"

	v1beta1 "k8s.io/api/extensions/v1beta1"
	av1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	INGRESS_NAME        = "eirini"
	INGRESS_API_VERSION = "extensions/v1beta1"
)

//go:generate counterfeiter . IngressClient
type IngressClient interface {
	Get(namespace, name string, options av1.GetOptions) (*v1beta1.Ingress, error)
	Update(namespace string, rule *v1beta1.Ingress) (*v1beta1.Ingress, error)
}

type ExtensionsIngressClient struct {
	Client typed.ExtensionsV1beta1Interface
}

func (i *ExtensionsIngressClient) Get(namespace, name string, options av1.GetOptions) (*v1beta1.Ingress, error) {
	ingresses := i.Client.Ingresses(namespace)
	return ingresses.Get(name, options)
}

func (i *ExtensionsIngressClient) Update(namespace string, rule *v1beta1.Ingress) (*v1beta1.Ingress, error) {
	ingresses := i.Client.Ingresses(namespace)
	return ingresses.Update(rule)
}

type IngressController struct {
	client       IngressClient
	kubeEndpoint string
}

func NewIngressController(client IngressClient, kubeEndpoint string) *IngressController {
	return &IngressController{
		client:       client,
		kubeEndpoint: kubeEndpoint,
	}
}

// func (i *IngressController) UpdateIngress(lrp opi.LRP, vcap VcapApp, namespace string) {
// 	//TODO
// }

// //func (i *IngressController) updateIngressNamespace(namespace string) {
// //i.client = i.clientset.v1beta1().Ingresses(namespace)
// //}

func (i *IngressController) updateIngress(lrp opi.LRP, vcap VcapApp, namespace string) error {
	ingress, err := i.client.Get(namespace, "eririni", av1.GetOptions{})
	if err != nil {
		return err
	}

	//"cube-kube.uk-south.containers.mybluemix.net"
	ingress.Spec.TLS[0].Hosts = append(ingress.Spec.TLS[0].Hosts, fmt.Sprintf("%s.%s", vcap.AppName, i.kubeEndpoint))
	rule := createIngressRule(lrp, vcap, i.kubeEndpoint)
	ingress.Spec.Rules = append(ingress.Spec.Rules, rule)

	if _, err = i.client.Update(namespace, ingress); err != nil {
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

func (i *IngressController) CreateIngress(namespace string) *ext.Ingress {
	ingress := &ext.Ingress{}

	ingress.APIVersion = INGRESS_NAME
	ingress.Kind = "Ingress"
	ingress.Name = INGRESS_API_VERSION
	ingress.Namespace = namespace

	return ingress
}
