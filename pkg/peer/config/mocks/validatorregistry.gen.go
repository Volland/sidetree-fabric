// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	ledgerconfig "github.com/trustbloc/fabric-peer-ext/pkg/config/ledgerconfig/config"
)

type ValidatorRegistry struct {
	RegisterStub        func(v ledgerconfig.Validator)
	registerMutex       sync.RWMutex
	registerArgsForCall []struct {
		v ledgerconfig.Validator
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ValidatorRegistry) Register(v ledgerconfig.Validator) {
	fake.registerMutex.Lock()
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		v ledgerconfig.Validator
	}{v})
	fake.recordInvocation("Register", []interface{}{v})
	fake.registerMutex.Unlock()
	if fake.RegisterStub != nil {
		fake.RegisterStub(v)
	}
}

func (fake *ValidatorRegistry) RegisterCallCount() int {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return len(fake.registerArgsForCall)
}

func (fake *ValidatorRegistry) RegisterArgsForCall(i int) ledgerconfig.Validator {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return fake.registerArgsForCall[i].v
}

func (fake *ValidatorRegistry) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ValidatorRegistry) recordInvocation(key string, args []interface{}) {
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
