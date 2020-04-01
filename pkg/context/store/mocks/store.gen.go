// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	commonledger "github.com/hyperledger/fabric/common/ledger"
)

type Store struct {
	QueryStub        func(query string) (commonledger.ResultsIterator, error)
	queryMutex       sync.RWMutex
	queryArgsForCall []struct {
		query string
	}
	queryReturns struct {
		result1 commonledger.ResultsIterator
		result2 error
	}
	queryReturnsOnCall map[int]struct {
		result1 commonledger.ResultsIterator
		result2 error
	}
	PutStub        func(value []byte) error
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		value []byte
	}
	putReturns struct {
		result1 error
	}
	putReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Store) Query(query string) (commonledger.ResultsIterator, error) {
	fake.queryMutex.Lock()
	ret, specificReturn := fake.queryReturnsOnCall[len(fake.queryArgsForCall)]
	fake.queryArgsForCall = append(fake.queryArgsForCall, struct {
		query string
	}{query})
	fake.recordInvocation("Query", []interface{}{query})
	fake.queryMutex.Unlock()
	if fake.QueryStub != nil {
		return fake.QueryStub(query)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.queryReturns.result1, fake.queryReturns.result2
}

func (fake *Store) QueryCallCount() int {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return len(fake.queryArgsForCall)
}

func (fake *Store) QueryArgsForCall(i int) string {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return fake.queryArgsForCall[i].query
}

func (fake *Store) QueryReturns(result1 commonledger.ResultsIterator, result2 error) {
	fake.QueryStub = nil
	fake.queryReturns = struct {
		result1 commonledger.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *Store) QueryReturnsOnCall(i int, result1 commonledger.ResultsIterator, result2 error) {
	fake.QueryStub = nil
	if fake.queryReturnsOnCall == nil {
		fake.queryReturnsOnCall = make(map[int]struct {
			result1 commonledger.ResultsIterator
			result2 error
		})
	}
	fake.queryReturnsOnCall[i] = struct {
		result1 commonledger.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *Store) Put(value []byte) error {
	var valueCopy []byte
	if value != nil {
		valueCopy = make([]byte, len(value))
		copy(valueCopy, value)
	}
	fake.putMutex.Lock()
	ret, specificReturn := fake.putReturnsOnCall[len(fake.putArgsForCall)]
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		value []byte
	}{valueCopy})
	fake.recordInvocation("Put", []interface{}{valueCopy})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(value)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.putReturns.result1
}

func (fake *Store) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *Store) PutArgsForCall(i int) []byte {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return fake.putArgsForCall[i].value
}

func (fake *Store) PutReturns(result1 error) {
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 error
	}{result1}
}

func (fake *Store) PutReturnsOnCall(i int, result1 error) {
	fake.PutStub = nil
	if fake.putReturnsOnCall == nil {
		fake.putReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.putReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Store) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Store) recordInvocation(key string, args []interface{}) {
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
