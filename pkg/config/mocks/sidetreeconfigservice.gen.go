// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/trustbloc/sidetree-core-go/pkg/api/protocol"
	"github.com/trustbloc/sidetree-fabric/pkg/config"
	"github.com/trustbloc/sidetree-fabric/pkg/rest/blockchainhandler"
	"github.com/trustbloc/sidetree-fabric/pkg/rest/dcashandler"
	"github.com/trustbloc/sidetree-fabric/pkg/rest/discoveryhandler"
	"github.com/trustbloc/sidetree-fabric/pkg/rest/filehandler"
	"github.com/trustbloc/sidetree-fabric/pkg/rest/sidetreehandler"
)

type SidetreeConfigService struct {
	LoadBlockchainHandlersStub        func(string, string) ([]blockchainhandler.Config, error)
	loadBlockchainHandlersMutex       sync.RWMutex
	loadBlockchainHandlersArgsForCall []struct {
		arg1 string
		arg2 string
	}
	loadBlockchainHandlersReturns struct {
		result1 []blockchainhandler.Config
		result2 error
	}
	loadBlockchainHandlersReturnsOnCall map[int]struct {
		result1 []blockchainhandler.Config
		result2 error
	}
	LoadDCASStub        func() (config.DCAS, error)
	loadDCASMutex       sync.RWMutex
	loadDCASArgsForCall []struct {
	}
	loadDCASReturns struct {
		result1 config.DCAS
		result2 error
	}
	loadDCASReturnsOnCall map[int]struct {
		result1 config.DCAS
		result2 error
	}
	LoadDCASHandlersStub        func(string, string) ([]dcashandler.Config, error)
	loadDCASHandlersMutex       sync.RWMutex
	loadDCASHandlersArgsForCall []struct {
		arg1 string
		arg2 string
	}
	loadDCASHandlersReturns struct {
		result1 []dcashandler.Config
		result2 error
	}
	loadDCASHandlersReturnsOnCall map[int]struct {
		result1 []dcashandler.Config
		result2 error
	}
	LoadDiscoveryHandlersStub        func(string, string) ([]discoveryhandler.Config, error)
	loadDiscoveryHandlersMutex       sync.RWMutex
	loadDiscoveryHandlersArgsForCall []struct {
		arg1 string
		arg2 string
	}
	loadDiscoveryHandlersReturns struct {
		result1 []discoveryhandler.Config
		result2 error
	}
	loadDiscoveryHandlersReturnsOnCall map[int]struct {
		result1 []discoveryhandler.Config
		result2 error
	}
	LoadFileHandlersStub        func(string, string) ([]filehandler.Config, error)
	loadFileHandlersMutex       sync.RWMutex
	loadFileHandlersArgsForCall []struct {
		arg1 string
		arg2 string
	}
	loadFileHandlersReturns struct {
		result1 []filehandler.Config
		result2 error
	}
	loadFileHandlersReturnsOnCall map[int]struct {
		result1 []filehandler.Config
		result2 error
	}
	LoadProtocolsStub        func(string) (map[string]protocol.Protocol, error)
	loadProtocolsMutex       sync.RWMutex
	loadProtocolsArgsForCall []struct {
		arg1 string
	}
	loadProtocolsReturns struct {
		result1 map[string]protocol.Protocol
		result2 error
	}
	loadProtocolsReturnsOnCall map[int]struct {
		result1 map[string]protocol.Protocol
		result2 error
	}
	LoadSidetreeStub        func(string) (config.Sidetree, error)
	loadSidetreeMutex       sync.RWMutex
	loadSidetreeArgsForCall []struct {
		arg1 string
	}
	loadSidetreeReturns struct {
		result1 config.Sidetree
		result2 error
	}
	loadSidetreeReturnsOnCall map[int]struct {
		result1 config.Sidetree
		result2 error
	}
	LoadSidetreeHandlersStub        func(string, string) ([]sidetreehandler.Config, error)
	loadSidetreeHandlersMutex       sync.RWMutex
	loadSidetreeHandlersArgsForCall []struct {
		arg1 string
		arg2 string
	}
	loadSidetreeHandlersReturns struct {
		result1 []sidetreehandler.Config
		result2 error
	}
	loadSidetreeHandlersReturnsOnCall map[int]struct {
		result1 []sidetreehandler.Config
		result2 error
	}
	LoadSidetreePeerStub        func(string, string) (config.SidetreePeer, error)
	loadSidetreePeerMutex       sync.RWMutex
	loadSidetreePeerArgsForCall []struct {
		arg1 string
		arg2 string
	}
	loadSidetreePeerReturns struct {
		result1 config.SidetreePeer
		result2 error
	}
	loadSidetreePeerReturnsOnCall map[int]struct {
		result1 config.SidetreePeer
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SidetreeConfigService) LoadBlockchainHandlers(arg1 string, arg2 string) ([]blockchainhandler.Config, error) {
	fake.loadBlockchainHandlersMutex.Lock()
	ret, specificReturn := fake.loadBlockchainHandlersReturnsOnCall[len(fake.loadBlockchainHandlersArgsForCall)]
	fake.loadBlockchainHandlersArgsForCall = append(fake.loadBlockchainHandlersArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LoadBlockchainHandlers", []interface{}{arg1, arg2})
	fake.loadBlockchainHandlersMutex.Unlock()
	if fake.LoadBlockchainHandlersStub != nil {
		return fake.LoadBlockchainHandlersStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.loadBlockchainHandlersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SidetreeConfigService) LoadBlockchainHandlersCallCount() int {
	fake.loadBlockchainHandlersMutex.RLock()
	defer fake.loadBlockchainHandlersMutex.RUnlock()
	return len(fake.loadBlockchainHandlersArgsForCall)
}

func (fake *SidetreeConfigService) LoadBlockchainHandlersCalls(stub func(string, string) ([]blockchainhandler.Config, error)) {
	fake.loadBlockchainHandlersMutex.Lock()
	defer fake.loadBlockchainHandlersMutex.Unlock()
	fake.LoadBlockchainHandlersStub = stub
}

func (fake *SidetreeConfigService) LoadBlockchainHandlersArgsForCall(i int) (string, string) {
	fake.loadBlockchainHandlersMutex.RLock()
	defer fake.loadBlockchainHandlersMutex.RUnlock()
	argsForCall := fake.loadBlockchainHandlersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *SidetreeConfigService) LoadBlockchainHandlersReturns(result1 []blockchainhandler.Config, result2 error) {
	fake.loadBlockchainHandlersMutex.Lock()
	defer fake.loadBlockchainHandlersMutex.Unlock()
	fake.LoadBlockchainHandlersStub = nil
	fake.loadBlockchainHandlersReturns = struct {
		result1 []blockchainhandler.Config
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadBlockchainHandlersReturnsOnCall(i int, result1 []blockchainhandler.Config, result2 error) {
	fake.loadBlockchainHandlersMutex.Lock()
	defer fake.loadBlockchainHandlersMutex.Unlock()
	fake.LoadBlockchainHandlersStub = nil
	if fake.loadBlockchainHandlersReturnsOnCall == nil {
		fake.loadBlockchainHandlersReturnsOnCall = make(map[int]struct {
			result1 []blockchainhandler.Config
			result2 error
		})
	}
	fake.loadBlockchainHandlersReturnsOnCall[i] = struct {
		result1 []blockchainhandler.Config
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadDCAS() (config.DCAS, error) {
	fake.loadDCASMutex.Lock()
	ret, specificReturn := fake.loadDCASReturnsOnCall[len(fake.loadDCASArgsForCall)]
	fake.loadDCASArgsForCall = append(fake.loadDCASArgsForCall, struct {
	}{})
	fake.recordInvocation("LoadDCAS", []interface{}{})
	fake.loadDCASMutex.Unlock()
	if fake.LoadDCASStub != nil {
		return fake.LoadDCASStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.loadDCASReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SidetreeConfigService) LoadDCASCallCount() int {
	fake.loadDCASMutex.RLock()
	defer fake.loadDCASMutex.RUnlock()
	return len(fake.loadDCASArgsForCall)
}

func (fake *SidetreeConfigService) LoadDCASCalls(stub func() (config.DCAS, error)) {
	fake.loadDCASMutex.Lock()
	defer fake.loadDCASMutex.Unlock()
	fake.LoadDCASStub = stub
}

func (fake *SidetreeConfigService) LoadDCASReturns(result1 config.DCAS, result2 error) {
	fake.loadDCASMutex.Lock()
	defer fake.loadDCASMutex.Unlock()
	fake.LoadDCASStub = nil
	fake.loadDCASReturns = struct {
		result1 config.DCAS
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadDCASReturnsOnCall(i int, result1 config.DCAS, result2 error) {
	fake.loadDCASMutex.Lock()
	defer fake.loadDCASMutex.Unlock()
	fake.LoadDCASStub = nil
	if fake.loadDCASReturnsOnCall == nil {
		fake.loadDCASReturnsOnCall = make(map[int]struct {
			result1 config.DCAS
			result2 error
		})
	}
	fake.loadDCASReturnsOnCall[i] = struct {
		result1 config.DCAS
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadDCASHandlers(arg1 string, arg2 string) ([]dcashandler.Config, error) {
	fake.loadDCASHandlersMutex.Lock()
	ret, specificReturn := fake.loadDCASHandlersReturnsOnCall[len(fake.loadDCASHandlersArgsForCall)]
	fake.loadDCASHandlersArgsForCall = append(fake.loadDCASHandlersArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LoadDCASHandlers", []interface{}{arg1, arg2})
	fake.loadDCASHandlersMutex.Unlock()
	if fake.LoadDCASHandlersStub != nil {
		return fake.LoadDCASHandlersStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.loadDCASHandlersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SidetreeConfigService) LoadDCASHandlersCallCount() int {
	fake.loadDCASHandlersMutex.RLock()
	defer fake.loadDCASHandlersMutex.RUnlock()
	return len(fake.loadDCASHandlersArgsForCall)
}

func (fake *SidetreeConfigService) LoadDCASHandlersCalls(stub func(string, string) ([]dcashandler.Config, error)) {
	fake.loadDCASHandlersMutex.Lock()
	defer fake.loadDCASHandlersMutex.Unlock()
	fake.LoadDCASHandlersStub = stub
}

func (fake *SidetreeConfigService) LoadDCASHandlersArgsForCall(i int) (string, string) {
	fake.loadDCASHandlersMutex.RLock()
	defer fake.loadDCASHandlersMutex.RUnlock()
	argsForCall := fake.loadDCASHandlersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *SidetreeConfigService) LoadDCASHandlersReturns(result1 []dcashandler.Config, result2 error) {
	fake.loadDCASHandlersMutex.Lock()
	defer fake.loadDCASHandlersMutex.Unlock()
	fake.LoadDCASHandlersStub = nil
	fake.loadDCASHandlersReturns = struct {
		result1 []dcashandler.Config
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadDCASHandlersReturnsOnCall(i int, result1 []dcashandler.Config, result2 error) {
	fake.loadDCASHandlersMutex.Lock()
	defer fake.loadDCASHandlersMutex.Unlock()
	fake.LoadDCASHandlersStub = nil
	if fake.loadDCASHandlersReturnsOnCall == nil {
		fake.loadDCASHandlersReturnsOnCall = make(map[int]struct {
			result1 []dcashandler.Config
			result2 error
		})
	}
	fake.loadDCASHandlersReturnsOnCall[i] = struct {
		result1 []dcashandler.Config
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadDiscoveryHandlers(arg1 string, arg2 string) ([]discoveryhandler.Config, error) {
	fake.loadDiscoveryHandlersMutex.Lock()
	ret, specificReturn := fake.loadDiscoveryHandlersReturnsOnCall[len(fake.loadDiscoveryHandlersArgsForCall)]
	fake.loadDiscoveryHandlersArgsForCall = append(fake.loadDiscoveryHandlersArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LoadDiscoveryHandlers", []interface{}{arg1, arg2})
	fake.loadDiscoveryHandlersMutex.Unlock()
	if fake.LoadDiscoveryHandlersStub != nil {
		return fake.LoadDiscoveryHandlersStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.loadDiscoveryHandlersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SidetreeConfigService) LoadDiscoveryHandlersCallCount() int {
	fake.loadDiscoveryHandlersMutex.RLock()
	defer fake.loadDiscoveryHandlersMutex.RUnlock()
	return len(fake.loadDiscoveryHandlersArgsForCall)
}

func (fake *SidetreeConfigService) LoadDiscoveryHandlersCalls(stub func(string, string) ([]discoveryhandler.Config, error)) {
	fake.loadDiscoveryHandlersMutex.Lock()
	defer fake.loadDiscoveryHandlersMutex.Unlock()
	fake.LoadDiscoveryHandlersStub = stub
}

func (fake *SidetreeConfigService) LoadDiscoveryHandlersArgsForCall(i int) (string, string) {
	fake.loadDiscoveryHandlersMutex.RLock()
	defer fake.loadDiscoveryHandlersMutex.RUnlock()
	argsForCall := fake.loadDiscoveryHandlersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *SidetreeConfigService) LoadDiscoveryHandlersReturns(result1 []discoveryhandler.Config, result2 error) {
	fake.loadDiscoveryHandlersMutex.Lock()
	defer fake.loadDiscoveryHandlersMutex.Unlock()
	fake.LoadDiscoveryHandlersStub = nil
	fake.loadDiscoveryHandlersReturns = struct {
		result1 []discoveryhandler.Config
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadDiscoveryHandlersReturnsOnCall(i int, result1 []discoveryhandler.Config, result2 error) {
	fake.loadDiscoveryHandlersMutex.Lock()
	defer fake.loadDiscoveryHandlersMutex.Unlock()
	fake.LoadDiscoveryHandlersStub = nil
	if fake.loadDiscoveryHandlersReturnsOnCall == nil {
		fake.loadDiscoveryHandlersReturnsOnCall = make(map[int]struct {
			result1 []discoveryhandler.Config
			result2 error
		})
	}
	fake.loadDiscoveryHandlersReturnsOnCall[i] = struct {
		result1 []discoveryhandler.Config
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadFileHandlers(arg1 string, arg2 string) ([]filehandler.Config, error) {
	fake.loadFileHandlersMutex.Lock()
	ret, specificReturn := fake.loadFileHandlersReturnsOnCall[len(fake.loadFileHandlersArgsForCall)]
	fake.loadFileHandlersArgsForCall = append(fake.loadFileHandlersArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LoadFileHandlers", []interface{}{arg1, arg2})
	fake.loadFileHandlersMutex.Unlock()
	if fake.LoadFileHandlersStub != nil {
		return fake.LoadFileHandlersStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.loadFileHandlersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SidetreeConfigService) LoadFileHandlersCallCount() int {
	fake.loadFileHandlersMutex.RLock()
	defer fake.loadFileHandlersMutex.RUnlock()
	return len(fake.loadFileHandlersArgsForCall)
}

func (fake *SidetreeConfigService) LoadFileHandlersCalls(stub func(string, string) ([]filehandler.Config, error)) {
	fake.loadFileHandlersMutex.Lock()
	defer fake.loadFileHandlersMutex.Unlock()
	fake.LoadFileHandlersStub = stub
}

func (fake *SidetreeConfigService) LoadFileHandlersArgsForCall(i int) (string, string) {
	fake.loadFileHandlersMutex.RLock()
	defer fake.loadFileHandlersMutex.RUnlock()
	argsForCall := fake.loadFileHandlersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *SidetreeConfigService) LoadFileHandlersReturns(result1 []filehandler.Config, result2 error) {
	fake.loadFileHandlersMutex.Lock()
	defer fake.loadFileHandlersMutex.Unlock()
	fake.LoadFileHandlersStub = nil
	fake.loadFileHandlersReturns = struct {
		result1 []filehandler.Config
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadFileHandlersReturnsOnCall(i int, result1 []filehandler.Config, result2 error) {
	fake.loadFileHandlersMutex.Lock()
	defer fake.loadFileHandlersMutex.Unlock()
	fake.LoadFileHandlersStub = nil
	if fake.loadFileHandlersReturnsOnCall == nil {
		fake.loadFileHandlersReturnsOnCall = make(map[int]struct {
			result1 []filehandler.Config
			result2 error
		})
	}
	fake.loadFileHandlersReturnsOnCall[i] = struct {
		result1 []filehandler.Config
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadProtocols(arg1 string) (map[string]protocol.Protocol, error) {
	fake.loadProtocolsMutex.Lock()
	ret, specificReturn := fake.loadProtocolsReturnsOnCall[len(fake.loadProtocolsArgsForCall)]
	fake.loadProtocolsArgsForCall = append(fake.loadProtocolsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("LoadProtocols", []interface{}{arg1})
	fake.loadProtocolsMutex.Unlock()
	if fake.LoadProtocolsStub != nil {
		return fake.LoadProtocolsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.loadProtocolsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SidetreeConfigService) LoadProtocolsCallCount() int {
	fake.loadProtocolsMutex.RLock()
	defer fake.loadProtocolsMutex.RUnlock()
	return len(fake.loadProtocolsArgsForCall)
}

func (fake *SidetreeConfigService) LoadProtocolsCalls(stub func(string) (map[string]protocol.Protocol, error)) {
	fake.loadProtocolsMutex.Lock()
	defer fake.loadProtocolsMutex.Unlock()
	fake.LoadProtocolsStub = stub
}

func (fake *SidetreeConfigService) LoadProtocolsArgsForCall(i int) string {
	fake.loadProtocolsMutex.RLock()
	defer fake.loadProtocolsMutex.RUnlock()
	argsForCall := fake.loadProtocolsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *SidetreeConfigService) LoadProtocolsReturns(result1 map[string]protocol.Protocol, result2 error) {
	fake.loadProtocolsMutex.Lock()
	defer fake.loadProtocolsMutex.Unlock()
	fake.LoadProtocolsStub = nil
	fake.loadProtocolsReturns = struct {
		result1 map[string]protocol.Protocol
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadProtocolsReturnsOnCall(i int, result1 map[string]protocol.Protocol, result2 error) {
	fake.loadProtocolsMutex.Lock()
	defer fake.loadProtocolsMutex.Unlock()
	fake.LoadProtocolsStub = nil
	if fake.loadProtocolsReturnsOnCall == nil {
		fake.loadProtocolsReturnsOnCall = make(map[int]struct {
			result1 map[string]protocol.Protocol
			result2 error
		})
	}
	fake.loadProtocolsReturnsOnCall[i] = struct {
		result1 map[string]protocol.Protocol
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadSidetree(arg1 string) (config.Sidetree, error) {
	fake.loadSidetreeMutex.Lock()
	ret, specificReturn := fake.loadSidetreeReturnsOnCall[len(fake.loadSidetreeArgsForCall)]
	fake.loadSidetreeArgsForCall = append(fake.loadSidetreeArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("LoadSidetree", []interface{}{arg1})
	fake.loadSidetreeMutex.Unlock()
	if fake.LoadSidetreeStub != nil {
		return fake.LoadSidetreeStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.loadSidetreeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SidetreeConfigService) LoadSidetreeCallCount() int {
	fake.loadSidetreeMutex.RLock()
	defer fake.loadSidetreeMutex.RUnlock()
	return len(fake.loadSidetreeArgsForCall)
}

func (fake *SidetreeConfigService) LoadSidetreeCalls(stub func(string) (config.Sidetree, error)) {
	fake.loadSidetreeMutex.Lock()
	defer fake.loadSidetreeMutex.Unlock()
	fake.LoadSidetreeStub = stub
}

func (fake *SidetreeConfigService) LoadSidetreeArgsForCall(i int) string {
	fake.loadSidetreeMutex.RLock()
	defer fake.loadSidetreeMutex.RUnlock()
	argsForCall := fake.loadSidetreeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *SidetreeConfigService) LoadSidetreeReturns(result1 config.Sidetree, result2 error) {
	fake.loadSidetreeMutex.Lock()
	defer fake.loadSidetreeMutex.Unlock()
	fake.LoadSidetreeStub = nil
	fake.loadSidetreeReturns = struct {
		result1 config.Sidetree
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadSidetreeReturnsOnCall(i int, result1 config.Sidetree, result2 error) {
	fake.loadSidetreeMutex.Lock()
	defer fake.loadSidetreeMutex.Unlock()
	fake.LoadSidetreeStub = nil
	if fake.loadSidetreeReturnsOnCall == nil {
		fake.loadSidetreeReturnsOnCall = make(map[int]struct {
			result1 config.Sidetree
			result2 error
		})
	}
	fake.loadSidetreeReturnsOnCall[i] = struct {
		result1 config.Sidetree
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadSidetreeHandlers(arg1 string, arg2 string) ([]sidetreehandler.Config, error) {
	fake.loadSidetreeHandlersMutex.Lock()
	ret, specificReturn := fake.loadSidetreeHandlersReturnsOnCall[len(fake.loadSidetreeHandlersArgsForCall)]
	fake.loadSidetreeHandlersArgsForCall = append(fake.loadSidetreeHandlersArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LoadSidetreeHandlers", []interface{}{arg1, arg2})
	fake.loadSidetreeHandlersMutex.Unlock()
	if fake.LoadSidetreeHandlersStub != nil {
		return fake.LoadSidetreeHandlersStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.loadSidetreeHandlersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SidetreeConfigService) LoadSidetreeHandlersCallCount() int {
	fake.loadSidetreeHandlersMutex.RLock()
	defer fake.loadSidetreeHandlersMutex.RUnlock()
	return len(fake.loadSidetreeHandlersArgsForCall)
}

func (fake *SidetreeConfigService) LoadSidetreeHandlersCalls(stub func(string, string) ([]sidetreehandler.Config, error)) {
	fake.loadSidetreeHandlersMutex.Lock()
	defer fake.loadSidetreeHandlersMutex.Unlock()
	fake.LoadSidetreeHandlersStub = stub
}

func (fake *SidetreeConfigService) LoadSidetreeHandlersArgsForCall(i int) (string, string) {
	fake.loadSidetreeHandlersMutex.RLock()
	defer fake.loadSidetreeHandlersMutex.RUnlock()
	argsForCall := fake.loadSidetreeHandlersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *SidetreeConfigService) LoadSidetreeHandlersReturns(result1 []sidetreehandler.Config, result2 error) {
	fake.loadSidetreeHandlersMutex.Lock()
	defer fake.loadSidetreeHandlersMutex.Unlock()
	fake.LoadSidetreeHandlersStub = nil
	fake.loadSidetreeHandlersReturns = struct {
		result1 []sidetreehandler.Config
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadSidetreeHandlersReturnsOnCall(i int, result1 []sidetreehandler.Config, result2 error) {
	fake.loadSidetreeHandlersMutex.Lock()
	defer fake.loadSidetreeHandlersMutex.Unlock()
	fake.LoadSidetreeHandlersStub = nil
	if fake.loadSidetreeHandlersReturnsOnCall == nil {
		fake.loadSidetreeHandlersReturnsOnCall = make(map[int]struct {
			result1 []sidetreehandler.Config
			result2 error
		})
	}
	fake.loadSidetreeHandlersReturnsOnCall[i] = struct {
		result1 []sidetreehandler.Config
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadSidetreePeer(arg1 string, arg2 string) (config.SidetreePeer, error) {
	fake.loadSidetreePeerMutex.Lock()
	ret, specificReturn := fake.loadSidetreePeerReturnsOnCall[len(fake.loadSidetreePeerArgsForCall)]
	fake.loadSidetreePeerArgsForCall = append(fake.loadSidetreePeerArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LoadSidetreePeer", []interface{}{arg1, arg2})
	fake.loadSidetreePeerMutex.Unlock()
	if fake.LoadSidetreePeerStub != nil {
		return fake.LoadSidetreePeerStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.loadSidetreePeerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SidetreeConfigService) LoadSidetreePeerCallCount() int {
	fake.loadSidetreePeerMutex.RLock()
	defer fake.loadSidetreePeerMutex.RUnlock()
	return len(fake.loadSidetreePeerArgsForCall)
}

func (fake *SidetreeConfigService) LoadSidetreePeerCalls(stub func(string, string) (config.SidetreePeer, error)) {
	fake.loadSidetreePeerMutex.Lock()
	defer fake.loadSidetreePeerMutex.Unlock()
	fake.LoadSidetreePeerStub = stub
}

func (fake *SidetreeConfigService) LoadSidetreePeerArgsForCall(i int) (string, string) {
	fake.loadSidetreePeerMutex.RLock()
	defer fake.loadSidetreePeerMutex.RUnlock()
	argsForCall := fake.loadSidetreePeerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *SidetreeConfigService) LoadSidetreePeerReturns(result1 config.SidetreePeer, result2 error) {
	fake.loadSidetreePeerMutex.Lock()
	defer fake.loadSidetreePeerMutex.Unlock()
	fake.LoadSidetreePeerStub = nil
	fake.loadSidetreePeerReturns = struct {
		result1 config.SidetreePeer
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) LoadSidetreePeerReturnsOnCall(i int, result1 config.SidetreePeer, result2 error) {
	fake.loadSidetreePeerMutex.Lock()
	defer fake.loadSidetreePeerMutex.Unlock()
	fake.LoadSidetreePeerStub = nil
	if fake.loadSidetreePeerReturnsOnCall == nil {
		fake.loadSidetreePeerReturnsOnCall = make(map[int]struct {
			result1 config.SidetreePeer
			result2 error
		})
	}
	fake.loadSidetreePeerReturnsOnCall[i] = struct {
		result1 config.SidetreePeer
		result2 error
	}{result1, result2}
}

func (fake *SidetreeConfigService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loadBlockchainHandlersMutex.RLock()
	defer fake.loadBlockchainHandlersMutex.RUnlock()
	fake.loadDCASMutex.RLock()
	defer fake.loadDCASMutex.RUnlock()
	fake.loadDCASHandlersMutex.RLock()
	defer fake.loadDCASHandlersMutex.RUnlock()
	fake.loadDiscoveryHandlersMutex.RLock()
	defer fake.loadDiscoveryHandlersMutex.RUnlock()
	fake.loadFileHandlersMutex.RLock()
	defer fake.loadFileHandlersMutex.RUnlock()
	fake.loadProtocolsMutex.RLock()
	defer fake.loadProtocolsMutex.RUnlock()
	fake.loadSidetreeMutex.RLock()
	defer fake.loadSidetreeMutex.RUnlock()
	fake.loadSidetreeHandlersMutex.RLock()
	defer fake.loadSidetreeHandlersMutex.RUnlock()
	fake.loadSidetreePeerMutex.RLock()
	defer fake.loadSidetreePeerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SidetreeConfigService) recordInvocation(key string, args []interface{}) {
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

var _ config.SidetreeService = new(SidetreeConfigService)
