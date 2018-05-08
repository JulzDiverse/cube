package k8s

import (
	"fmt"

	"github.com/julz/cube/opi"
	ext "k8s.io/api/extensions/v1beta1"

	av1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	//"k8s.io/client-go/kubernetes"
	kubernetes "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

const (
	INGRESS_NAME        = "eirini"
	INGRESS_API_VERSION = "extensions/v1beta1"
)

type IngressController struct {
	client kubernetes.IngressInterface
	//client       kubernetes.Interface
	kubeEndpoint string
}

func NewIngressController(client kubernetes.IngressInterface, kubeEndpoint string) *IngressController {
	return &IngressController{
		client:       client,
		kubeEndpoint: kubeEndpoint,
	}
}

func (i *IngressController) UpdateIngress(lrp opi.LRP, vcap VcapApp, namespace string) {
	//TODO
}

//func (i *IngressController) updateIngressNamespace(namespace string) {
//i.client = i.clientset.v1beta1().Ingresses(namespace)
//}

func (i *IngressController) updateIngressRules(lrp opi.LRP, vcap VcapApp, namespace string) error {
	ingress, err := i.client.Get("eririni", av1.GetOptions{})
	if err != nil {
		return err
	}

	//"cube-kube.uk-south.containers.mybluemix.net"
	ingress.Spec.TLS[0].Hosts = append(ingress.Spec.TLS[0].Hosts, fmt.Sprintf("%s.%s", vcap.AppName, i.kubeEndpoint))
	rule := createIngressRule(lrp, vcap, i.kubeEndpoint)
	ingress.Spec.Rules = append(ingress.Spec.Rules, rule)

	if _, err = i.client.Update(ingress); err != nil {
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
