// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type PortAllocator struct {
	AllocatePortStub        func(handle string, port int) (int, error)
	allocatePortMutex       sync.RWMutex
	allocatePortArgsForCall []struct {
		handle string
		port   int
	}
	allocatePortReturns struct {
		result1 int
		result2 error
	}
	allocatePortReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	ReleaseAllPortsStub        func(handle string) error
	releaseAllPortsMutex       sync.RWMutex
	releaseAllPortsArgsForCall []struct {
		handle string
	}
	releaseAllPortsReturns struct {
		result1 error
	}
	releaseAllPortsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PortAllocator) AllocatePort(handle string, port int) (int, error) {
	fake.allocatePortMutex.Lock()
	ret, specificReturn := fake.allocatePortReturnsOnCall[len(fake.allocatePortArgsForCall)]
	fake.allocatePortArgsForCall = append(fake.allocatePortArgsForCall, struct {
		handle string
		port   int
	}{handle, port})
	fake.recordInvocation("AllocatePort", []interface{}{handle, port})
	fake.allocatePortMutex.Unlock()
	if fake.AllocatePortStub != nil {
		return fake.AllocatePortStub(handle, port)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.allocatePortReturns.result1, fake.allocatePortReturns.result2
}

func (fake *PortAllocator) AllocatePortCallCount() int {
	fake.allocatePortMutex.RLock()
	defer fake.allocatePortMutex.RUnlock()
	return len(fake.allocatePortArgsForCall)
}

func (fake *PortAllocator) AllocatePortArgsForCall(i int) (string, int) {
	fake.allocatePortMutex.RLock()
	defer fake.allocatePortMutex.RUnlock()
	return fake.allocatePortArgsForCall[i].handle, fake.allocatePortArgsForCall[i].port
}

func (fake *PortAllocator) AllocatePortReturns(result1 int, result2 error) {
	fake.AllocatePortStub = nil
	fake.allocatePortReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *PortAllocator) AllocatePortReturnsOnCall(i int, result1 int, result2 error) {
	fake.AllocatePortStub = nil
	if fake.allocatePortReturnsOnCall == nil {
		fake.allocatePortReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.allocatePortReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *PortAllocator) ReleaseAllPorts(handle string) error {
	fake.releaseAllPortsMutex.Lock()
	ret, specificReturn := fake.releaseAllPortsReturnsOnCall[len(fake.releaseAllPortsArgsForCall)]
	fake.releaseAllPortsArgsForCall = append(fake.releaseAllPortsArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("ReleaseAllPorts", []interface{}{handle})
	fake.releaseAllPortsMutex.Unlock()
	if fake.ReleaseAllPortsStub != nil {
		return fake.ReleaseAllPortsStub(handle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.releaseAllPortsReturns.result1
}

func (fake *PortAllocator) ReleaseAllPortsCallCount() int {
	fake.releaseAllPortsMutex.RLock()
	defer fake.releaseAllPortsMutex.RUnlock()
	return len(fake.releaseAllPortsArgsForCall)
}

func (fake *PortAllocator) ReleaseAllPortsArgsForCall(i int) string {
	fake.releaseAllPortsMutex.RLock()
	defer fake.releaseAllPortsMutex.RUnlock()
	return fake.releaseAllPortsArgsForCall[i].handle
}

func (fake *PortAllocator) ReleaseAllPortsReturns(result1 error) {
	fake.ReleaseAllPortsStub = nil
	fake.releaseAllPortsReturns = struct {
		result1 error
	}{result1}
}

func (fake *PortAllocator) ReleaseAllPortsReturnsOnCall(i int, result1 error) {
	fake.ReleaseAllPortsStub = nil
	if fake.releaseAllPortsReturnsOnCall == nil {
		fake.releaseAllPortsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.releaseAllPortsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PortAllocator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allocatePortMutex.RLock()
	defer fake.allocatePortMutex.RUnlock()
	fake.releaseAllPortsMutex.RLock()
	defer fake.releaseAllPortsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PortAllocator) recordInvocation(key string, args []interface{}) {
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
