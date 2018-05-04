// Code generated by counterfeiter. DO NOT EDIT.
package managerfakes

import (
	"bosh-dns/dns/manager"
	"sync"
)

type FakeDNSManager struct {
	SetPrimaryStub        func() error
	setPrimaryMutex       sync.RWMutex
	setPrimaryArgsForCall []struct{}
	setPrimaryReturns     struct {
		result1 error
	}
	setPrimaryReturnsOnCall map[int]struct {
		result1 error
	}
	ReadStub        func() ([]string, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct{}
	readReturns     struct {
		result1 []string
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDNSManager) SetPrimary() error {
	fake.setPrimaryMutex.Lock()
	ret, specificReturn := fake.setPrimaryReturnsOnCall[len(fake.setPrimaryArgsForCall)]
	fake.setPrimaryArgsForCall = append(fake.setPrimaryArgsForCall, struct{}{})
	fake.recordInvocation("SetPrimary", []interface{}{})
	fake.setPrimaryMutex.Unlock()
	if fake.SetPrimaryStub != nil {
		return fake.SetPrimaryStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setPrimaryReturns.result1
}

func (fake *FakeDNSManager) SetPrimaryCallCount() int {
	fake.setPrimaryMutex.RLock()
	defer fake.setPrimaryMutex.RUnlock()
	return len(fake.setPrimaryArgsForCall)
}

func (fake *FakeDNSManager) SetPrimaryReturns(result1 error) {
	fake.SetPrimaryStub = nil
	fake.setPrimaryReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDNSManager) SetPrimaryReturnsOnCall(i int, result1 error) {
	fake.SetPrimaryStub = nil
	if fake.setPrimaryReturnsOnCall == nil {
		fake.setPrimaryReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPrimaryReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDNSManager) Read() ([]string, error) {
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct{}{})
	fake.recordInvocation("Read", []interface{}{})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.readReturns.result1, fake.readReturns.result2
}

func (fake *FakeDNSManager) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeDNSManager) ReadReturns(result1 []string, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeDNSManager) ReadReturnsOnCall(i int, result1 []string, result2 error) {
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeDNSManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setPrimaryMutex.RLock()
	defer fake.setPrimaryMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDNSManager) recordInvocation(key string, args []interface{}) {
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

var _ manager.DNSManager = new(FakeDNSManager)
