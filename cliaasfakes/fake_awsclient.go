// This file was generated by counterfeiter
package cliaasfakes

import (
	"sync"

	"github.com/pivotal-cf/cliaas"
)

type FakeAWSClient struct {
	CreateVMStub        func(ami, instanceType, name, keyName, subnetID, securityGroupID string) (string, error)
	createVMMutex       sync.RWMutex
	createVMArgsForCall []struct {
		ami             string
		instanceType    string
		name            string
		keyName         string
		subnetID        string
		securityGroupID string
	}
	createVMReturns struct {
		result1 string
		result2 error
	}
	DeleteVMStub        func(instanceID string) error
	deleteVMMutex       sync.RWMutex
	deleteVMArgsForCall []struct {
		instanceID string
	}
	deleteVMReturns struct {
		result1 error
	}
	GetVMInfoStub        func(name string) (cliaas.VMInfo, error)
	getVMInfoMutex       sync.RWMutex
	getVMInfoArgsForCall []struct {
		name string
	}
	getVMInfoReturns struct {
		result1 cliaas.VMInfo
		result2 error
	}
	StopVMStub        func(instanceID string) error
	stopVMMutex       sync.RWMutex
	stopVMArgsForCall []struct {
		instanceID string
	}
	stopVMReturns struct {
		result1 error
	}
	AssignPublicIPStub        func(instance, ip string) error
	assignPublicIPMutex       sync.RWMutex
	assignPublicIPArgsForCall []struct {
		instance string
		ip       string
	}
	assignPublicIPReturns struct {
		result1 error
	}
	WaitForStatusStub        func(instanceID string, status string) error
	waitForStatusMutex       sync.RWMutex
	waitForStatusArgsForCall []struct {
		instanceID string
		status     string
	}
	waitForStatusReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAWSClient) CreateVM(ami string, instanceType string, name string, keyName string, subnetID string, securityGroupID string) (string, error) {
	fake.createVMMutex.Lock()
	fake.createVMArgsForCall = append(fake.createVMArgsForCall, struct {
		ami             string
		instanceType    string
		name            string
		keyName         string
		subnetID        string
		securityGroupID string
	}{ami, instanceType, name, keyName, subnetID, securityGroupID})
	fake.recordInvocation("CreateVM", []interface{}{ami, instanceType, name, keyName, subnetID, securityGroupID})
	fake.createVMMutex.Unlock()
	if fake.CreateVMStub != nil {
		return fake.CreateVMStub(ami, instanceType, name, keyName, subnetID, securityGroupID)
	} else {
		return fake.createVMReturns.result1, fake.createVMReturns.result2
	}
}

func (fake *FakeAWSClient) CreateVMCallCount() int {
	fake.createVMMutex.RLock()
	defer fake.createVMMutex.RUnlock()
	return len(fake.createVMArgsForCall)
}

func (fake *FakeAWSClient) CreateVMArgsForCall(i int) (string, string, string, string, string, string) {
	fake.createVMMutex.RLock()
	defer fake.createVMMutex.RUnlock()
	return fake.createVMArgsForCall[i].ami, fake.createVMArgsForCall[i].instanceType, fake.createVMArgsForCall[i].name, fake.createVMArgsForCall[i].keyName, fake.createVMArgsForCall[i].subnetID, fake.createVMArgsForCall[i].securityGroupID
}

func (fake *FakeAWSClient) CreateVMReturns(result1 string, result2 error) {
	fake.CreateVMStub = nil
	fake.createVMReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClient) DeleteVM(instanceID string) error {
	fake.deleteVMMutex.Lock()
	fake.deleteVMArgsForCall = append(fake.deleteVMArgsForCall, struct {
		instanceID string
	}{instanceID})
	fake.recordInvocation("DeleteVM", []interface{}{instanceID})
	fake.deleteVMMutex.Unlock()
	if fake.DeleteVMStub != nil {
		return fake.DeleteVMStub(instanceID)
	} else {
		return fake.deleteVMReturns.result1
	}
}

func (fake *FakeAWSClient) DeleteVMCallCount() int {
	fake.deleteVMMutex.RLock()
	defer fake.deleteVMMutex.RUnlock()
	return len(fake.deleteVMArgsForCall)
}

func (fake *FakeAWSClient) DeleteVMArgsForCall(i int) string {
	fake.deleteVMMutex.RLock()
	defer fake.deleteVMMutex.RUnlock()
	return fake.deleteVMArgsForCall[i].instanceID
}

func (fake *FakeAWSClient) DeleteVMReturns(result1 error) {
	fake.DeleteVMStub = nil
	fake.deleteVMReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAWSClient) GetVMInfo(name string) (cliaas.VMInfo, error) {
	fake.getVMInfoMutex.Lock()
	fake.getVMInfoArgsForCall = append(fake.getVMInfoArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("GetVMInfo", []interface{}{name})
	fake.getVMInfoMutex.Unlock()
	if fake.GetVMInfoStub != nil {
		return fake.GetVMInfoStub(name)
	} else {
		return fake.getVMInfoReturns.result1, fake.getVMInfoReturns.result2
	}
}

func (fake *FakeAWSClient) GetVMInfoCallCount() int {
	fake.getVMInfoMutex.RLock()
	defer fake.getVMInfoMutex.RUnlock()
	return len(fake.getVMInfoArgsForCall)
}

func (fake *FakeAWSClient) GetVMInfoArgsForCall(i int) string {
	fake.getVMInfoMutex.RLock()
	defer fake.getVMInfoMutex.RUnlock()
	return fake.getVMInfoArgsForCall[i].name
}

func (fake *FakeAWSClient) GetVMInfoReturns(result1 cliaas.VMInfo, result2 error) {
	fake.GetVMInfoStub = nil
	fake.getVMInfoReturns = struct {
		result1 cliaas.VMInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClient) StopVM(instanceID string) error {
	fake.stopVMMutex.Lock()
	fake.stopVMArgsForCall = append(fake.stopVMArgsForCall, struct {
		instanceID string
	}{instanceID})
	fake.recordInvocation("StopVM", []interface{}{instanceID})
	fake.stopVMMutex.Unlock()
	if fake.StopVMStub != nil {
		return fake.StopVMStub(instanceID)
	} else {
		return fake.stopVMReturns.result1
	}
}

func (fake *FakeAWSClient) StopVMCallCount() int {
	fake.stopVMMutex.RLock()
	defer fake.stopVMMutex.RUnlock()
	return len(fake.stopVMArgsForCall)
}

func (fake *FakeAWSClient) StopVMArgsForCall(i int) string {
	fake.stopVMMutex.RLock()
	defer fake.stopVMMutex.RUnlock()
	return fake.stopVMArgsForCall[i].instanceID
}

func (fake *FakeAWSClient) StopVMReturns(result1 error) {
	fake.StopVMStub = nil
	fake.stopVMReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAWSClient) AssignPublicIP(instance string, ip string) error {
	fake.assignPublicIPMutex.Lock()
	fake.assignPublicIPArgsForCall = append(fake.assignPublicIPArgsForCall, struct {
		instance string
		ip       string
	}{instance, ip})
	fake.recordInvocation("AssignPublicIP", []interface{}{instance, ip})
	fake.assignPublicIPMutex.Unlock()
	if fake.AssignPublicIPStub != nil {
		return fake.AssignPublicIPStub(instance, ip)
	} else {
		return fake.assignPublicIPReturns.result1
	}
}

func (fake *FakeAWSClient) AssignPublicIPCallCount() int {
	fake.assignPublicIPMutex.RLock()
	defer fake.assignPublicIPMutex.RUnlock()
	return len(fake.assignPublicIPArgsForCall)
}

func (fake *FakeAWSClient) AssignPublicIPArgsForCall(i int) (string, string) {
	fake.assignPublicIPMutex.RLock()
	defer fake.assignPublicIPMutex.RUnlock()
	return fake.assignPublicIPArgsForCall[i].instance, fake.assignPublicIPArgsForCall[i].ip
}

func (fake *FakeAWSClient) AssignPublicIPReturns(result1 error) {
	fake.AssignPublicIPStub = nil
	fake.assignPublicIPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAWSClient) WaitForStatus(instanceID string, status string) error {
	fake.waitForStatusMutex.Lock()
	fake.waitForStatusArgsForCall = append(fake.waitForStatusArgsForCall, struct {
		instanceID string
		status     string
	}{instanceID, status})
	fake.recordInvocation("WaitForStatus", []interface{}{instanceID, status})
	fake.waitForStatusMutex.Unlock()
	if fake.WaitForStatusStub != nil {
		return fake.WaitForStatusStub(instanceID, status)
	} else {
		return fake.waitForStatusReturns.result1
	}
}

func (fake *FakeAWSClient) WaitForStatusCallCount() int {
	fake.waitForStatusMutex.RLock()
	defer fake.waitForStatusMutex.RUnlock()
	return len(fake.waitForStatusArgsForCall)
}

func (fake *FakeAWSClient) WaitForStatusArgsForCall(i int) (string, string) {
	fake.waitForStatusMutex.RLock()
	defer fake.waitForStatusMutex.RUnlock()
	return fake.waitForStatusArgsForCall[i].instanceID, fake.waitForStatusArgsForCall[i].status
}

func (fake *FakeAWSClient) WaitForStatusReturns(result1 error) {
	fake.WaitForStatusStub = nil
	fake.waitForStatusReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAWSClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createVMMutex.RLock()
	defer fake.createVMMutex.RUnlock()
	fake.deleteVMMutex.RLock()
	defer fake.deleteVMMutex.RUnlock()
	fake.getVMInfoMutex.RLock()
	defer fake.getVMInfoMutex.RUnlock()
	fake.stopVMMutex.RLock()
	defer fake.stopVMMutex.RUnlock()
	fake.assignPublicIPMutex.RLock()
	defer fake.assignPublicIPMutex.RUnlock()
	fake.waitForStatusMutex.RLock()
	defer fake.waitForStatusMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeAWSClient) recordInvocation(key string, args []interface{}) {
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

var _ cliaas.AWSClient = new(FakeAWSClient)
