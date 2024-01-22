// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/bborbe/kv"
)

type Iterator struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	ItemStub        func() kv.Item
	itemMutex       sync.RWMutex
	itemArgsForCall []struct {
	}
	itemReturns struct {
		result1 kv.Item
	}
	itemReturnsOnCall map[int]struct {
		result1 kv.Item
	}
	NextStub        func()
	nextMutex       sync.RWMutex
	nextArgsForCall []struct {
	}
	RewindStub        func()
	rewindMutex       sync.RWMutex
	rewindArgsForCall []struct {
	}
	SeekStub        func([]byte)
	seekMutex       sync.RWMutex
	seekArgsForCall []struct {
		arg1 []byte
	}
	ValidStub        func() bool
	validMutex       sync.RWMutex
	validArgsForCall []struct {
	}
	validReturns struct {
		result1 bool
	}
	validReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Iterator) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		fake.CloseStub()
	}
}

func (fake *Iterator) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *Iterator) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *Iterator) Item() kv.Item {
	fake.itemMutex.Lock()
	ret, specificReturn := fake.itemReturnsOnCall[len(fake.itemArgsForCall)]
	fake.itemArgsForCall = append(fake.itemArgsForCall, struct {
	}{})
	stub := fake.ItemStub
	fakeReturns := fake.itemReturns
	fake.recordInvocation("Item", []interface{}{})
	fake.itemMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Iterator) ItemCallCount() int {
	fake.itemMutex.RLock()
	defer fake.itemMutex.RUnlock()
	return len(fake.itemArgsForCall)
}

func (fake *Iterator) ItemCalls(stub func() kv.Item) {
	fake.itemMutex.Lock()
	defer fake.itemMutex.Unlock()
	fake.ItemStub = stub
}

func (fake *Iterator) ItemReturns(result1 kv.Item) {
	fake.itemMutex.Lock()
	defer fake.itemMutex.Unlock()
	fake.ItemStub = nil
	fake.itemReturns = struct {
		result1 kv.Item
	}{result1}
}

func (fake *Iterator) ItemReturnsOnCall(i int, result1 kv.Item) {
	fake.itemMutex.Lock()
	defer fake.itemMutex.Unlock()
	fake.ItemStub = nil
	if fake.itemReturnsOnCall == nil {
		fake.itemReturnsOnCall = make(map[int]struct {
			result1 kv.Item
		})
	}
	fake.itemReturnsOnCall[i] = struct {
		result1 kv.Item
	}{result1}
}

func (fake *Iterator) Next() {
	fake.nextMutex.Lock()
	fake.nextArgsForCall = append(fake.nextArgsForCall, struct {
	}{})
	stub := fake.NextStub
	fake.recordInvocation("Next", []interface{}{})
	fake.nextMutex.Unlock()
	if stub != nil {
		fake.NextStub()
	}
}

func (fake *Iterator) NextCallCount() int {
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	return len(fake.nextArgsForCall)
}

func (fake *Iterator) NextCalls(stub func()) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = stub
}

func (fake *Iterator) Rewind() {
	fake.rewindMutex.Lock()
	fake.rewindArgsForCall = append(fake.rewindArgsForCall, struct {
	}{})
	stub := fake.RewindStub
	fake.recordInvocation("Rewind", []interface{}{})
	fake.rewindMutex.Unlock()
	if stub != nil {
		fake.RewindStub()
	}
}

func (fake *Iterator) RewindCallCount() int {
	fake.rewindMutex.RLock()
	defer fake.rewindMutex.RUnlock()
	return len(fake.rewindArgsForCall)
}

func (fake *Iterator) RewindCalls(stub func()) {
	fake.rewindMutex.Lock()
	defer fake.rewindMutex.Unlock()
	fake.RewindStub = stub
}

func (fake *Iterator) Seek(arg1 []byte) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.seekMutex.Lock()
	fake.seekArgsForCall = append(fake.seekArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.SeekStub
	fake.recordInvocation("Seek", []interface{}{arg1Copy})
	fake.seekMutex.Unlock()
	if stub != nil {
		fake.SeekStub(arg1)
	}
}

func (fake *Iterator) SeekCallCount() int {
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	return len(fake.seekArgsForCall)
}

func (fake *Iterator) SeekCalls(stub func([]byte)) {
	fake.seekMutex.Lock()
	defer fake.seekMutex.Unlock()
	fake.SeekStub = stub
}

func (fake *Iterator) SeekArgsForCall(i int) []byte {
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	argsForCall := fake.seekArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Iterator) Valid() bool {
	fake.validMutex.Lock()
	ret, specificReturn := fake.validReturnsOnCall[len(fake.validArgsForCall)]
	fake.validArgsForCall = append(fake.validArgsForCall, struct {
	}{})
	stub := fake.ValidStub
	fakeReturns := fake.validReturns
	fake.recordInvocation("Valid", []interface{}{})
	fake.validMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Iterator) ValidCallCount() int {
	fake.validMutex.RLock()
	defer fake.validMutex.RUnlock()
	return len(fake.validArgsForCall)
}

func (fake *Iterator) ValidCalls(stub func() bool) {
	fake.validMutex.Lock()
	defer fake.validMutex.Unlock()
	fake.ValidStub = stub
}

func (fake *Iterator) ValidReturns(result1 bool) {
	fake.validMutex.Lock()
	defer fake.validMutex.Unlock()
	fake.ValidStub = nil
	fake.validReturns = struct {
		result1 bool
	}{result1}
}

func (fake *Iterator) ValidReturnsOnCall(i int, result1 bool) {
	fake.validMutex.Lock()
	defer fake.validMutex.Unlock()
	fake.ValidStub = nil
	if fake.validReturnsOnCall == nil {
		fake.validReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.validReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *Iterator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.itemMutex.RLock()
	defer fake.itemMutex.RUnlock()
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	fake.rewindMutex.RLock()
	defer fake.rewindMutex.RUnlock()
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	fake.validMutex.RLock()
	defer fake.validMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Iterator) recordInvocation(key string, args []interface{}) {
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

var _ kv.Iterator = new(Iterator)