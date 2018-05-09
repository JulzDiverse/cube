package k8s_test

import (
	. "github.com/julz/cube/k8s"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	av1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
)

var _ = FDescribe("Ingress", func() {

	var (
		namespace    string
		kubeEndpoint string
		fakeClient   kubernetes.Interface
		ingCtrl      *IngressManager
	)

	BeforeEach(func() {
		namespace = "testing"
		kubeEndpoint = "alfheim"

		fakeClient = fake.NewSimpleClientset()
		ingCtrl = NewIngressManager(fakeClient, kubeEndpoint)
	})

	Context("Create Ingress", func() {
		Context("should be successful", func() {
			It("should succeed", func() {
				err := ingCtrl.CreateIngress(namespace)
				Expect(err).ToNot(HaveOccurred())

				ing, err := fakeClient.ExtensionsV1beta1().Ingresses(namespace).Get("eirini", av1.GetOptions{})
				Expect(err).ToNot(HaveOccurred())
				Expect(ing.Name).To(Equal("eirini"))
			})
		})
	})
})
