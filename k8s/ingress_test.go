package k8s_test

import (
	. "github.com/julz/cube/k8s"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	av1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fakes "k8s.io/client-go/kubernetes/fake"
)

var _ = FDescribe("Ingress", func() {

	Context("Update ingress", func() {
		Context("is successful", func() {

			var (
				fakeClient *fakes.Clientset
				controller *IngressController
				namespace  string
			)

			BeforeEach(func() {
				fakeClient = fakes.NewSimpleClientset()
				namespace = "testing"

				ingress, err := fakeClient.ExtensionsV1beta1().Ingresses(namespace).Gt("eirini", av1.GetOptions{})
				controller = NewIngressController(fakeClient, "test-endpoint")
			})
		})

	})
})
