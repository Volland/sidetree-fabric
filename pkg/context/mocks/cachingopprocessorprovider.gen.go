// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/trustbloc/sidetree-core-go/pkg/dochandler"
	"github.com/trustbloc/sidetree-fabric/pkg/rest/sidetreehandler"
)

type CachingOpProcessorProvider struct {
	CreateCachingOperationProcessorStub        func(channelID string, cfg sidetreehandler.Config, resolver dochandler.OperationProcessor) dochandler.OperationProcessor
	createCachingOperationProcessorMutex       sync.RWMutex
	createCachingOperationProcessorArgsForCall []struct {
		channelID string
		cfg       sidetreehandler.Config
		resolver  dochandler.OperationProcessor
	}
	createCachingOperationProcessorReturns struct {
		result1 dochandler.OperationProcessor
	}
	createCachingOperationProcessorReturnsOnCall map[int]struct {
		result1 dochandler.OperationProcessor
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CachingOpProcessorProvider) CreateCachingOperationProcessor(channelID string, cfg sidetreehandler.Config, resolver dochandler.OperationProcessor) dochandler.OperationProcessor {
	fake.createCachingOperationProcessorMutex.Lock()
	ret, specificReturn := fake.createCachingOperationProcessorReturnsOnCall[len(fake.createCachingOperationProcessorArgsForCall)]
	fake.createCachingOperationProcessorArgsForCall = append(fake.createCachingOperationProcessorArgsForCall, struct {
		channelID string
		cfg       sidetreehandler.Config
		resolver  dochandler.OperationProcessor
	}{channelID, cfg, resolver})
	fake.recordInvocation("CreateCachingOperationProcessor", []interface{}{channelID, cfg, resolver})
	fake.createCachingOperationProcessorMutex.Unlock()
	if fake.CreateCachingOperationProcessorStub != nil {
		return fake.CreateCachingOperationProcessorStub(channelID, cfg, resolver)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createCachingOperationProcessorReturns.result1
}

func (fake *CachingOpProcessorProvider) CreateCachingOperationProcessorCallCount() int {
	fake.createCachingOperationProcessorMutex.RLock()
	defer fake.createCachingOperationProcessorMutex.RUnlock()
	return len(fake.createCachingOperationProcessorArgsForCall)
}

func (fake *CachingOpProcessorProvider) CreateCachingOperationProcessorArgsForCall(i int) (string, sidetreehandler.Config, dochandler.OperationProcessor) {
	fake.createCachingOperationProcessorMutex.RLock()
	defer fake.createCachingOperationProcessorMutex.RUnlock()
	return fake.createCachingOperationProcessorArgsForCall[i].channelID, fake.createCachingOperationProcessorArgsForCall[i].cfg, fake.createCachingOperationProcessorArgsForCall[i].resolver
}

func (fake *CachingOpProcessorProvider) CreateCachingOperationProcessorReturns(result1 dochandler.OperationProcessor) {
	fake.CreateCachingOperationProcessorStub = nil
	fake.createCachingOperationProcessorReturns = struct {
		result1 dochandler.OperationProcessor
	}{result1}
}

func (fake *CachingOpProcessorProvider) CreateCachingOperationProcessorReturnsOnCall(i int, result1 dochandler.OperationProcessor) {
	fake.CreateCachingOperationProcessorStub = nil
	if fake.createCachingOperationProcessorReturnsOnCall == nil {
		fake.createCachingOperationProcessorReturnsOnCall = make(map[int]struct {
			result1 dochandler.OperationProcessor
		})
	}
	fake.createCachingOperationProcessorReturnsOnCall[i] = struct {
		result1 dochandler.OperationProcessor
	}{result1}
}

func (fake *CachingOpProcessorProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCachingOperationProcessorMutex.RLock()
	defer fake.createCachingOperationProcessorMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CachingOpProcessorProvider) recordInvocation(key string, args []interface{}) {
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