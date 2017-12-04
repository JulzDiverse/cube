// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net"
	"sync"

	"code.cloudfoundry.org/diego-ssh/handlers"
)

type FakeDialer struct {
	DialStub        func(net, addr string) (net.Conn, error)
	dialMutex       sync.RWMutex
	dialArgsForCall []struct {
		net  string
		addr string
	}
	dialReturns struct {
		result1 net.Conn
		result2 error
	}
	dialReturnsOnCall map[int]struct {
		result1 net.Conn
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDialer) Dial(net string, addr string) (net.Conn, error) {
	fake.dialMutex.Lock()
	ret, specificReturn := fake.dialReturnsOnCall[len(fake.dialArgsForCall)]
	fake.dialArgsForCall = append(fake.dialArgsForCall, struct {
		net  string
		addr string
	}{net, addr})
	fake.recordInvocation("Dial", []interface{}{net, addr})
	fake.dialMutex.Unlock()
	if fake.DialStub != nil {
		return fake.DialStub(net, addr)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.dialReturns.result1, fake.dialReturns.result2
}

func (fake *FakeDialer) DialCallCount() int {
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	return len(fake.dialArgsForCall)
}

func (fake *FakeDialer) DialArgsForCall(i int) (string, string) {
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	return fake.dialArgsForCall[i].net, fake.dialArgsForCall[i].addr
}

func (fake *FakeDialer) DialReturns(result1 net.Conn, result2 error) {
	fake.DialStub = nil
	fake.dialReturns = struct {
		result1 net.Conn
		result2 error
	}{result1, result2}
}

func (fake *FakeDialer) DialReturnsOnCall(i int, result1 net.Conn, result2 error) {
	fake.DialStub = nil
	if fake.dialReturnsOnCall == nil {
		fake.dialReturnsOnCall = make(map[int]struct {
			result1 net.Conn
			result2 error
		})
	}
	fake.dialReturnsOnCall[i] = struct {
		result1 net.Conn
		result2 error
	}{result1, result2}
}

func (fake *FakeDialer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDialer) recordInvocation(key string, args []interface{}) {
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

var _ handlers.Dialer = new(FakeDialer)
