// Code generated by counterfeiter. DO NOT EDIT.
package service

import (
	"sync"

	"github.com/nginx/agent/v3/api/grpc/instances"
	"github.com/nginx/agent/v3/internal/model/os"
)

type FakeInstanceServiceInterface struct {
	GetInstancesStub        func() ([]*instances.Instance, error)
	getInstancesMutex       sync.RWMutex
	getInstancesArgsForCall []struct {
	}
	getInstancesReturns struct {
		result1 []*instances.Instance
		result2 error
	}
	getInstancesReturnsOnCall map[int]struct {
		result1 []*instances.Instance
		result2 error
	}
	UpdateProcessesStub        func([]*os.Process)
	updateProcessesMutex       sync.RWMutex
	updateProcessesArgsForCall []struct {
		arg1 []*os.Process
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInstanceServiceInterface) GetInstances() ([]*instances.Instance, error) {
	fake.getInstancesMutex.Lock()
	ret, specificReturn := fake.getInstancesReturnsOnCall[len(fake.getInstancesArgsForCall)]
	fake.getInstancesArgsForCall = append(fake.getInstancesArgsForCall, struct {
	}{})
	stub := fake.GetInstancesStub
	fakeReturns := fake.getInstancesReturns
	fake.recordInvocation("GetInstances", []interface{}{})
	fake.getInstancesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstanceServiceInterface) GetInstancesCallCount() int {
	fake.getInstancesMutex.RLock()
	defer fake.getInstancesMutex.RUnlock()
	return len(fake.getInstancesArgsForCall)
}

func (fake *FakeInstanceServiceInterface) GetInstancesCalls(stub func() ([]*instances.Instance, error)) {
	fake.getInstancesMutex.Lock()
	defer fake.getInstancesMutex.Unlock()
	fake.GetInstancesStub = stub
}

func (fake *FakeInstanceServiceInterface) GetInstancesReturns(result1 []*instances.Instance, result2 error) {
	fake.getInstancesMutex.Lock()
	defer fake.getInstancesMutex.Unlock()
	fake.GetInstancesStub = nil
	fake.getInstancesReturns = struct {
		result1 []*instances.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeInstanceServiceInterface) GetInstancesReturnsOnCall(i int, result1 []*instances.Instance, result2 error) {
	fake.getInstancesMutex.Lock()
	defer fake.getInstancesMutex.Unlock()
	fake.GetInstancesStub = nil
	if fake.getInstancesReturnsOnCall == nil {
		fake.getInstancesReturnsOnCall = make(map[int]struct {
			result1 []*instances.Instance
			result2 error
		})
	}
	fake.getInstancesReturnsOnCall[i] = struct {
		result1 []*instances.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeInstanceServiceInterface) UpdateProcesses(arg1 []*os.Process) {
	var arg1Copy []*os.Process
	if arg1 != nil {
		arg1Copy = make([]*os.Process, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.updateProcessesMutex.Lock()
	fake.updateProcessesArgsForCall = append(fake.updateProcessesArgsForCall, struct {
		arg1 []*os.Process
	}{arg1Copy})
	stub := fake.UpdateProcessesStub
	fake.recordInvocation("UpdateProcesses", []interface{}{arg1Copy})
	fake.updateProcessesMutex.Unlock()
	if stub != nil {
		fake.UpdateProcessesStub(arg1)
	}
}

func (fake *FakeInstanceServiceInterface) UpdateProcessesCallCount() int {
	fake.updateProcessesMutex.RLock()
	defer fake.updateProcessesMutex.RUnlock()
	return len(fake.updateProcessesArgsForCall)
}

func (fake *FakeInstanceServiceInterface) UpdateProcessesCalls(stub func([]*os.Process)) {
	fake.updateProcessesMutex.Lock()
	defer fake.updateProcessesMutex.Unlock()
	fake.UpdateProcessesStub = stub
}

func (fake *FakeInstanceServiceInterface) UpdateProcessesArgsForCall(i int) []*os.Process {
	fake.updateProcessesMutex.RLock()
	defer fake.updateProcessesMutex.RUnlock()
	argsForCall := fake.updateProcessesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstanceServiceInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getInstancesMutex.RLock()
	defer fake.getInstancesMutex.RUnlock()
	fake.updateProcessesMutex.RLock()
	defer fake.updateProcessesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInstanceServiceInterface) recordInvocation(key string, args []interface{}) {
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

var _ InstanceServiceInterface = new(FakeInstanceServiceInterface)
