// Code generated by counterfeiter. DO NOT EDIT.
package fake_handlers

import (
	"sync"

	"code.cloudfoundry.org/diego-ssh/handlers"
	"code.cloudfoundry.org/lager"
	"golang.org/x/crypto/ssh"
)

type FakeNewChannelHandler struct {
	HandleNewChannelStub        func(logger lager.Logger, newChannel ssh.NewChannel)
	handleNewChannelMutex       sync.RWMutex
	handleNewChannelArgsForCall []struct {
		logger     lager.Logger
		newChannel ssh.NewChannel
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNewChannelHandler) HandleNewChannel(logger lager.Logger, newChannel ssh.NewChannel) {
	fake.handleNewChannelMutex.Lock()
	fake.handleNewChannelArgsForCall = append(fake.handleNewChannelArgsForCall, struct {
		logger     lager.Logger
		newChannel ssh.NewChannel
	}{logger, newChannel})
	fake.recordInvocation("HandleNewChannel", []interface{}{logger, newChannel})
	fake.handleNewChannelMutex.Unlock()
	if fake.HandleNewChannelStub != nil {
		fake.HandleNewChannelStub(logger, newChannel)
	}
}

func (fake *FakeNewChannelHandler) HandleNewChannelCallCount() int {
	fake.handleNewChannelMutex.RLock()
	defer fake.handleNewChannelMutex.RUnlock()
	return len(fake.handleNewChannelArgsForCall)
}

func (fake *FakeNewChannelHandler) HandleNewChannelArgsForCall(i int) (lager.Logger, ssh.NewChannel) {
	fake.handleNewChannelMutex.RLock()
	defer fake.handleNewChannelMutex.RUnlock()
	return fake.handleNewChannelArgsForCall[i].logger, fake.handleNewChannelArgsForCall[i].newChannel
}

func (fake *FakeNewChannelHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleNewChannelMutex.RLock()
	defer fake.handleNewChannelMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNewChannelHandler) recordInvocation(key string, args []interface{}) {
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

var _ handlers.NewChannelHandler = new(FakeNewChannelHandler)
