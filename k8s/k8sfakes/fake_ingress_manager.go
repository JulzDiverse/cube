// Code generated by counterfeiter. DO NOT EDIT.
package k8sfakes

import (
	"sync"

	"github.com/julz/cube/k8s"
	"github.com/julz/cube/opi"
	ext "k8s.io/api/extensions/v1beta1"
)

type FakeIngressManager struct {
	CreateIngressStub        func(namespace string) (*ext.Ingress, error)
	createIngressMutex       sync.RWMutex
	createIngressArgsForCall []struct {
		namespace string
	}
	createIngressReturns struct {
		result1 *ext.Ingress
		result2 error
	}
	createIngressReturnsOnCall map[int]struct {
		result1 *ext.Ingress
		result2 error
	}
	UpdateIngressStub        func(namespace string, lrp opi.LRP, vcap k8s.VcapApp) error
	updateIngressMutex       sync.RWMutex
	updateIngressArgsForCall []struct {
		namespace string
		lrp       opi.LRP
		vcap      k8s.VcapApp
	}
	updateIngressReturns struct {
		result1 error
	}
	updateIngressReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIngressManager) CreateIngress(namespace string) (*ext.Ingress, error) {
	fake.createIngressMutex.Lock()
	ret, specificReturn := fake.createIngressReturnsOnCall[len(fake.createIngressArgsForCall)]
	fake.createIngressArgsForCall = append(fake.createIngressArgsForCall, struct {
		namespace string
	}{namespace})
	fake.recordInvocation("CreateIngress", []interface{}{namespace})
	fake.createIngressMutex.Unlock()
	if fake.CreateIngressStub != nil {
		return fake.CreateIngressStub(namespace)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createIngressReturns.result1, fake.createIngressReturns.result2
}

func (fake *FakeIngressManager) CreateIngressCallCount() int {
	fake.createIngressMutex.RLock()
	defer fake.createIngressMutex.RUnlock()
	return len(fake.createIngressArgsForCall)
}

func (fake *FakeIngressManager) CreateIngressArgsForCall(i int) string {
	fake.createIngressMutex.RLock()
	defer fake.createIngressMutex.RUnlock()
	return fake.createIngressArgsForCall[i].namespace
}

func (fake *FakeIngressManager) CreateIngressReturns(result1 *ext.Ingress, result2 error) {
	fake.CreateIngressStub = nil
	fake.createIngressReturns = struct {
		result1 *ext.Ingress
		result2 error
	}{result1, result2}
}

func (fake *FakeIngressManager) CreateIngressReturnsOnCall(i int, result1 *ext.Ingress, result2 error) {
	fake.CreateIngressStub = nil
	if fake.createIngressReturnsOnCall == nil {
		fake.createIngressReturnsOnCall = make(map[int]struct {
			result1 *ext.Ingress
			result2 error
		})
	}
	fake.createIngressReturnsOnCall[i] = struct {
		result1 *ext.Ingress
		result2 error
	}{result1, result2}
}

func (fake *FakeIngressManager) UpdateIngress(namespace string, lrp opi.LRP, vcap k8s.VcapApp) error {
	fake.updateIngressMutex.Lock()
	ret, specificReturn := fake.updateIngressReturnsOnCall[len(fake.updateIngressArgsForCall)]
	fake.updateIngressArgsForCall = append(fake.updateIngressArgsForCall, struct {
		namespace string
		lrp       opi.LRP
		vcap      k8s.VcapApp
	}{namespace, lrp, vcap})
	fake.recordInvocation("UpdateIngress", []interface{}{namespace, lrp, vcap})
	fake.updateIngressMutex.Unlock()
	if fake.UpdateIngressStub != nil {
		return fake.UpdateIngressStub(namespace, lrp, vcap)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateIngressReturns.result1
}

func (fake *FakeIngressManager) UpdateIngressCallCount() int {
	fake.updateIngressMutex.RLock()
	defer fake.updateIngressMutex.RUnlock()
	return len(fake.updateIngressArgsForCall)
}

func (fake *FakeIngressManager) UpdateIngressArgsForCall(i int) (string, opi.LRP, k8s.VcapApp) {
	fake.updateIngressMutex.RLock()
	defer fake.updateIngressMutex.RUnlock()
	return fake.updateIngressArgsForCall[i].namespace, fake.updateIngressArgsForCall[i].lrp, fake.updateIngressArgsForCall[i].vcap
}

func (fake *FakeIngressManager) UpdateIngressReturns(result1 error) {
	fake.UpdateIngressStub = nil
	fake.updateIngressReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressManager) UpdateIngressReturnsOnCall(i int, result1 error) {
	fake.UpdateIngressStub = nil
	if fake.updateIngressReturnsOnCall == nil {
		fake.updateIngressReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateIngressReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createIngressMutex.RLock()
	defer fake.createIngressMutex.RUnlock()
	fake.updateIngressMutex.RLock()
	defer fake.updateIngressMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIngressManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ k8s.IngressManager = new(FakeIngressManager)
