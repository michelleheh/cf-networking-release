// This file was generated by counterfeiter
package fakes

import "sync"

type WarrantClient struct {
	GetTokenStub        func(clientName, clientSecret string) (string, error)
	getTokenMutex       sync.RWMutex
	getTokenArgsForCall []struct {
		clientName   string
		clientSecret string
	}
	getTokenReturns struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *WarrantClient) GetToken(clientName string, clientSecret string) (string, error) {
	fake.getTokenMutex.Lock()
	fake.getTokenArgsForCall = append(fake.getTokenArgsForCall, struct {
		clientName   string
		clientSecret string
	}{clientName, clientSecret})
	fake.recordInvocation("GetToken", []interface{}{clientName, clientSecret})
	fake.getTokenMutex.Unlock()
	if fake.GetTokenStub != nil {
		return fake.GetTokenStub(clientName, clientSecret)
	}
	return fake.getTokenReturns.result1, fake.getTokenReturns.result2
}

func (fake *WarrantClient) GetTokenCallCount() int {
	fake.getTokenMutex.RLock()
	defer fake.getTokenMutex.RUnlock()
	return len(fake.getTokenArgsForCall)
}

func (fake *WarrantClient) GetTokenArgsForCall(i int) (string, string) {
	fake.getTokenMutex.RLock()
	defer fake.getTokenMutex.RUnlock()
	return fake.getTokenArgsForCall[i].clientName, fake.getTokenArgsForCall[i].clientSecret
}

func (fake *WarrantClient) GetTokenReturns(result1 string, result2 error) {
	fake.GetTokenStub = nil
	fake.getTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *WarrantClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTokenMutex.RLock()
	defer fake.getTokenMutex.RUnlock()
	return fake.invocations
}

func (fake *WarrantClient) recordInvocation(key string, args []interface{}) {
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