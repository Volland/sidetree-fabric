// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/trustbloc/sidetree-core-go/pkg/api/cas"
	"github.com/trustbloc/sidetree-core-go/pkg/api/protocol"
	"github.com/trustbloc/sidetree-fabric/pkg/config"
	"github.com/trustbloc/sidetree-fabric/pkg/context/common"
)

type ProtocolFactory struct {
	CreateStub        func(string, protocol.Protocol, cas.Client, common.OperationStore, string, config.Sidetree) (protocol.Version, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
		arg2 protocol.Protocol
		arg3 cas.Client
		arg4 common.OperationStore
		arg5 string
		arg6 config.Sidetree
	}
	createReturns struct {
		result1 protocol.Version
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 protocol.Version
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ProtocolFactory) Create(arg1 string, arg2 protocol.Protocol, arg3 cas.Client, arg4 common.OperationStore, arg5 string, arg6 config.Sidetree) (protocol.Version, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
		arg2 protocol.Protocol
		arg3 cas.Client
		arg4 common.OperationStore
		arg5 string
		arg6 config.Sidetree
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ProtocolFactory) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *ProtocolFactory) CreateCalls(stub func(string, protocol.Protocol, cas.Client, common.OperationStore, string, config.Sidetree) (protocol.Version, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *ProtocolFactory) CreateArgsForCall(i int) (string, protocol.Protocol, cas.Client, common.OperationStore, string, config.Sidetree) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *ProtocolFactory) CreateReturns(result1 protocol.Version, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 protocol.Version
		result2 error
	}{result1, result2}
}

func (fake *ProtocolFactory) CreateReturnsOnCall(i int, result1 protocol.Version, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 protocol.Version
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 protocol.Version
		result2 error
	}{result1, result2}
}

func (fake *ProtocolFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ProtocolFactory) recordInvocation(key string, args []interface{}) {
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
