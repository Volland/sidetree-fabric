// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/trustbloc/sidetree-core-go/pkg/api/cas"
	"github.com/trustbloc/sidetree-core-go/pkg/api/protocol"
	common2 "github.com/trustbloc/sidetree-fabric/pkg/common"
	ctxcommon "github.com/trustbloc/sidetree-fabric/pkg/context/common"
)

type ProtocolFactory struct {
	CreateStub        func(version string, p protocol.Protocol, casClient cas.Client, opStore ctxcommon.OperationStore, docType common2.DocumentType) (protocol.Version, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		version   string
		p         protocol.Protocol
		casClient cas.Client
		opStore   ctxcommon.OperationStore
		docType   common2.DocumentType
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

func (fake *ProtocolFactory) Create(version string, p protocol.Protocol, casClient cas.Client, opStore ctxcommon.OperationStore, docType common2.DocumentType) (protocol.Version, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		version   string
		p         protocol.Protocol
		casClient cas.Client
		opStore   ctxcommon.OperationStore
		docType   common2.DocumentType
	}{version, p, casClient, opStore, docType})
	fake.recordInvocation("Create", []interface{}{version, p, casClient, opStore, docType})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(version, p, casClient, opStore, docType)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *ProtocolFactory) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *ProtocolFactory) CreateArgsForCall(i int) (string, protocol.Protocol, cas.Client, ctxcommon.OperationStore, common2.DocumentType) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].version, fake.createArgsForCall[i].p, fake.createArgsForCall[i].casClient, fake.createArgsForCall[i].opStore, fake.createArgsForCall[i].docType
}

func (fake *ProtocolFactory) CreateReturns(result1 protocol.Version, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 protocol.Version
		result2 error
	}{result1, result2}
}

func (fake *ProtocolFactory) CreateReturnsOnCall(i int, result1 protocol.Version, result2 error) {
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
