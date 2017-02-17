// This file was generated by counterfeiter
package gcpfakes

import (
	"sync"

	"github.com/c0-ops/cliaas/iaas/gcp"
	compute "google.golang.org/api/compute/v1"
)

type FakeGoogleComputeClient struct {
	ListStub        func(project string, zone string) (*compute.InstanceList, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		project string
		zone    string
	}
	listReturns struct {
		result1 *compute.InstanceList
		result2 error
	}
	DeleteStub        func(project string, zone string, instanceName string) (*compute.Operation, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		project      string
		zone         string
		instanceName string
	}
	deleteReturns struct {
		result1 *compute.Operation
		result2 error
	}
	InsertStub        func(project string, zone string, instance *compute.Instance) (*compute.Operation, error)
	insertMutex       sync.RWMutex
	insertArgsForCall []struct {
		project  string
		zone     string
		instance *compute.Instance
	}
	insertReturns struct {
		result1 *compute.Operation
		result2 error
	}
	StopStub        func(project string, zone string, instanceName string) (*compute.Operation, error)
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
		project      string
		zone         string
		instanceName string
	}
	stopReturns struct {
		result1 *compute.Operation
		result2 error
	}
}

func (fake *FakeGoogleComputeClient) List(project string, zone string) (*compute.InstanceList, error) {
	fake.listMutex.Lock()
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		project string
		zone    string
	}{project, zone})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(project, zone)
	} else {
		return fake.listReturns.result1, fake.listReturns.result2
	}
}

func (fake *FakeGoogleComputeClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeGoogleComputeClient) ListArgsForCall(i int) (string, string) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].project, fake.listArgsForCall[i].zone
}

func (fake *FakeGoogleComputeClient) ListReturns(result1 *compute.InstanceList, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *compute.InstanceList
		result2 error
	}{result1, result2}
}

func (fake *FakeGoogleComputeClient) Delete(project string, zone string, instanceName string) (*compute.Operation, error) {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		project      string
		zone         string
		instanceName string
	}{project, zone, instanceName})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(project, zone, instanceName)
	} else {
		return fake.deleteReturns.result1, fake.deleteReturns.result2
	}
}

func (fake *FakeGoogleComputeClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeGoogleComputeClient) DeleteArgsForCall(i int) (string, string, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].project, fake.deleteArgsForCall[i].zone, fake.deleteArgsForCall[i].instanceName
}

func (fake *FakeGoogleComputeClient) DeleteReturns(result1 *compute.Operation, result2 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 *compute.Operation
		result2 error
	}{result1, result2}
}

func (fake *FakeGoogleComputeClient) Insert(project string, zone string, instance *compute.Instance) (*compute.Operation, error) {
	fake.insertMutex.Lock()
	fake.insertArgsForCall = append(fake.insertArgsForCall, struct {
		project  string
		zone     string
		instance *compute.Instance
	}{project, zone, instance})
	fake.insertMutex.Unlock()
	if fake.InsertStub != nil {
		return fake.InsertStub(project, zone, instance)
	} else {
		return fake.insertReturns.result1, fake.insertReturns.result2
	}
}

func (fake *FakeGoogleComputeClient) InsertCallCount() int {
	fake.insertMutex.RLock()
	defer fake.insertMutex.RUnlock()
	return len(fake.insertArgsForCall)
}

func (fake *FakeGoogleComputeClient) InsertArgsForCall(i int) (string, string, *compute.Instance) {
	fake.insertMutex.RLock()
	defer fake.insertMutex.RUnlock()
	return fake.insertArgsForCall[i].project, fake.insertArgsForCall[i].zone, fake.insertArgsForCall[i].instance
}

func (fake *FakeGoogleComputeClient) InsertReturns(result1 *compute.Operation, result2 error) {
	fake.InsertStub = nil
	fake.insertReturns = struct {
		result1 *compute.Operation
		result2 error
	}{result1, result2}
}

func (fake *FakeGoogleComputeClient) Stop(project string, zone string, instanceName string) (*compute.Operation, error) {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
		project      string
		zone         string
		instanceName string
	}{project, zone, instanceName})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub(project, zone, instanceName)
	} else {
		return fake.stopReturns.result1, fake.stopReturns.result2
	}
}

func (fake *FakeGoogleComputeClient) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeGoogleComputeClient) StopArgsForCall(i int) (string, string, string) {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return fake.stopArgsForCall[i].project, fake.stopArgsForCall[i].zone, fake.stopArgsForCall[i].instanceName
}

func (fake *FakeGoogleComputeClient) StopReturns(result1 *compute.Operation, result2 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 *compute.Operation
		result2 error
	}{result1, result2}
}

var _ gcp.GoogleComputeClient = new(FakeGoogleComputeClient)
