package k8s_test

import (
	"fmt"

	"github.com/julz/cube/k8s/k8sfakes"

	. "github.com/julz/cube/k8s"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ext "k8s.io/api/extensions/v1beta1"
	av1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes/fake"
	extensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

var _ = FDescribe("Ingress", func() {

	Context("Ingress Client", func() {
		var (
			client      IngressClient
			fakeClient  extensionsv1beta1.ExtensionsV1beta1Interface
			namespace   string
			ingressName string
		)

		createIngress := func(name, namespace string) *ext.Ingress {
			ingress := &ext.Ingress{}

			ingress.APIVersion = "eirini"
			ingress.Kind = "Ingress"
			ingress.Name = name
			ingress.Namespace = namespace

			rule := ext.IngressRule{
				Host: "app-name.kube-endpoint",
			}

			rule.HTTP = &ext.HTTPIngressRuleValue{
				Paths: []ext.HTTPIngressPath{
					ext.HTTPIngressPath{
						Path: "/",
						Backend: ext.IngressBackend{
							ServiceName: fmt.Sprintf("cf-appname"),
							ServicePort: intstr.FromInt(8080),
						},
					},
				},
			}

			ingress.Spec.Rules = append(ingress.Spec.Rules, rule)

			return ingress
		}

		BeforeEach(func() {
			fakeClient = fake.NewSimpleClientset().ExtensionsV1beta1()
			namespace = "test-namespace"
			ingressName = "eirini"
		})

		JustBeforeEach(func() {
			client = &ExtensionsIngressClient{
				Client: fakeClient,
			}
		})

		Context("Get ingress", func() {

			Context("when it already exists", func() {

				var (
					ingress *ext.Ingress
					result  *ext.Ingress
					err     error
				)

				BeforeEach(func() {
					ingress = createIngress(ingressName, namespace)

					_, e := fakeClient.Ingresses(namespace).Create(ingress)
					Expect(e).ToNot(HaveOccurred())
				})

				JustBeforeEach(func() {
					result, err = client.Get(namespace, ingressName, av1.GetOptions{})

				})

				It("should not error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("should return the expected result", func() {
					Expect(result).To(Equal(ingress))
				})
			})

			Context("when it does not exist", func() {

				It("should return a 'not found' error", func() {
					_, err := client.Get(namespace, ingressName, av1.GetOptions{})
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError(ContainSubstring("\"eirini\" not found")))
				})
			})

		})

		Context("Update ingress rule", func() {

			Context("when ingress already exists", func() {

				var (
					ingress *ext.Ingress
					result  *ext.Ingress
					err     error
				)

				BeforeEach(func() {
					ingress = createIngress(ingressName, namespace)

					_, e := fakeClient.Ingresses(namespace).Create(ingress)
					Expect(e).ToNot(HaveOccurred())
				})

				JustBeforeEach(func() {
					newRule := ext.IngressRule{
						Host: "another-app-name.kube-endpoint",
					}
					ingress.Spec.Rules = append(ingress.Spec.Rules, newRule)

					result, err = client.Update("test-namespace", ingress)
				})

				It("should not return an error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("should update the record", func() {
					Expect(result).To(Equal(ingress))
				})
			})

			Context("when that ingress does not exist", func() {

				It("should return a 'not found' error", func() {
					_, err := client.Update(namespace, createIngress(ingressName, namespace))
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError(ContainSubstring("\"eirini\" not found")))
				})
			})
		})
	})

	Context("Ingress Controller", func() {

		Context("Update ingress", func() {
			Context("is successful", func() {

				var (
					fakeClient *k8sfakes.FakeIngressClient
					controller *IngressController
					namespace  string
				)

				BeforeEach(func() {
					fakeClient = new(k8sfakes.FakeIngressClient)
					namespace = "testing"

					// mock away
					controller = NewIngressController(fakeClient, "test-endpoint")
				})
			})

		})

	})
})
