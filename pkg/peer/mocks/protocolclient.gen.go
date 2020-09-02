// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/trustbloc/sidetree-core-go/pkg/api/protocol"
)

type ProtocolClient struct {
	CurrentStub        func() (protocol.Protocol, error)
	currentMutex       sync.RWMutex
	currentArgsForCall []struct{}
	currentReturns     struct {
		result1 protocol.Protocol
		result2 error
	}
	currentReturnsOnCall map[int]struct {
		result1 protocol.Protocol
		result2 error
	}
	GetStub        func(transactionTime uint64) (protocol.Protocol, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		transactionTime uint64
	}
	getReturns struct {
		result1 protocol.Protocol
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 protocol.Protocol
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ProtocolClient) Current() (protocol.Protocol, error) {
	fake.currentMutex.Lock()
	ret, specificReturn := fake.currentReturnsOnCall[len(fake.currentArgsForCall)]
	fake.currentArgsForCall = append(fake.currentArgsForCall, struct{}{})
	fake.recordInvocation("Current", []interface{}{})
	fake.currentMutex.Unlock()
	if fake.CurrentStub != nil {
		return fake.CurrentStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.currentReturns.result1, fake.currentReturns.result2
}

func (fake *ProtocolClient) CurrentCallCount() int {
	fake.currentMutex.RLock()
	defer fake.currentMutex.RUnlock()
	return len(fake.currentArgsForCall)
}

func (fake *ProtocolClient) CurrentReturns(result1 protocol.Protocol, result2 error) {
	fake.CurrentStub = nil
	fake.currentReturns = struct {
		result1 protocol.Protocol
		result2 error
	}{result1, result2}
}

func (fake *ProtocolClient) CurrentReturnsOnCall(i int, result1 protocol.Protocol, result2 error) {
	fake.CurrentStub = nil
	if fake.currentReturnsOnCall == nil {
		fake.currentReturnsOnCall = make(map[int]struct {
			result1 protocol.Protocol
			result2 error
		})
	}
	fake.currentReturnsOnCall[i] = struct {
		result1 protocol.Protocol
		result2 error
	}{result1, result2}
}

func (fake *ProtocolClient) Get(transactionTime uint64) (protocol.Protocol, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		transactionTime uint64
	}{transactionTime})
	fake.recordInvocation("Get", []interface{}{transactionTime})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(transactionTime)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getReturns.result1, fake.getReturns.result2
}

func (fake *ProtocolClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *ProtocolClient) GetArgsForCall(i int) uint64 {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].transactionTime
}

func (fake *ProtocolClient) GetReturns(result1 protocol.Protocol, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 protocol.Protocol
		result2 error
	}{result1, result2}
}

func (fake *ProtocolClient) GetReturnsOnCall(i int, result1 protocol.Protocol, result2 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 protocol.Protocol
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 protocol.Protocol
		result2 error
	}{result1, result2}
}

func (fake *ProtocolClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.currentMutex.RLock()
	defer fake.currentMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ProtocolClient) recordInvocation(key string, args []interface{}) {
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

var _ protocol.Client = new(ProtocolClient)
