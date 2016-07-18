// This file was generated by counterfeiter
package infakes

import (
	"sync"

	"github.com/pivotal-cf-experimental/go-pivnet"
)

type FakePivnetClient struct {
	GetReleaseStub        func(productSlug string, productVersion string) (pivnet.Release, error)
	getReleaseMutex       sync.RWMutex
	getReleaseArgsForCall []struct {
		productSlug    string
		productVersion string
	}
	getReleaseReturns struct {
		result1 pivnet.Release
		result2 error
	}
	AcceptEULAStub        func(productSlug string, releaseID int) error
	acceptEULAMutex       sync.RWMutex
	acceptEULAArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	acceptEULAReturns struct {
		result1 error
	}
	GetProductFilesStub        func(productSlug string, releaseID int) ([]pivnet.ProductFile, error)
	getProductFilesMutex       sync.RWMutex
	getProductFilesArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	getProductFilesReturns struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	GetProductFileStub        func(productSlug string, releaseID int, productFileID int) (pivnet.ProductFile, error)
	getProductFileMutex       sync.RWMutex
	getProductFileArgsForCall []struct {
		productSlug   string
		releaseID     int
		productFileID int
	}
	getProductFileReturns struct {
		result1 pivnet.ProductFile
		result2 error
	}
	ReleaseDependenciesStub        func(productSlug string, releaseID int) ([]pivnet.ReleaseDependency, error)
	releaseDependenciesMutex       sync.RWMutex
	releaseDependenciesArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	releaseDependenciesReturns struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) GetRelease(productSlug string, productVersion string) (pivnet.Release, error) {
	fake.getReleaseMutex.Lock()
	fake.getReleaseArgsForCall = append(fake.getReleaseArgsForCall, struct {
		productSlug    string
		productVersion string
	}{productSlug, productVersion})
	fake.recordInvocation("GetRelease", []interface{}{productSlug, productVersion})
	fake.getReleaseMutex.Unlock()
	if fake.GetReleaseStub != nil {
		return fake.GetReleaseStub(productSlug, productVersion)
	} else {
		return fake.getReleaseReturns.result1, fake.getReleaseReturns.result2
	}
}

func (fake *FakePivnetClient) GetReleaseCallCount() int {
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	return len(fake.getReleaseArgsForCall)
}

func (fake *FakePivnetClient) GetReleaseArgsForCall(i int) (string, string) {
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	return fake.getReleaseArgsForCall[i].productSlug, fake.getReleaseArgsForCall[i].productVersion
}

func (fake *FakePivnetClient) GetReleaseReturns(result1 pivnet.Release, result2 error) {
	fake.GetReleaseStub = nil
	fake.getReleaseReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) AcceptEULA(productSlug string, releaseID int) error {
	fake.acceptEULAMutex.Lock()
	fake.acceptEULAArgsForCall = append(fake.acceptEULAArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("AcceptEULA", []interface{}{productSlug, releaseID})
	fake.acceptEULAMutex.Unlock()
	if fake.AcceptEULAStub != nil {
		return fake.AcceptEULAStub(productSlug, releaseID)
	} else {
		return fake.acceptEULAReturns.result1
	}
}

func (fake *FakePivnetClient) AcceptEULACallCount() int {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return len(fake.acceptEULAArgsForCall)
}

func (fake *FakePivnetClient) AcceptEULAArgsForCall(i int) (string, int) {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return fake.acceptEULAArgsForCall[i].productSlug, fake.acceptEULAArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) AcceptEULAReturns(result1 error) {
	fake.AcceptEULAStub = nil
	fake.acceptEULAReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) GetProductFiles(productSlug string, releaseID int) ([]pivnet.ProductFile, error) {
	fake.getProductFilesMutex.Lock()
	fake.getProductFilesArgsForCall = append(fake.getProductFilesArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("GetProductFiles", []interface{}{productSlug, releaseID})
	fake.getProductFilesMutex.Unlock()
	if fake.GetProductFilesStub != nil {
		return fake.GetProductFilesStub(productSlug, releaseID)
	} else {
		return fake.getProductFilesReturns.result1, fake.getProductFilesReturns.result2
	}
}

func (fake *FakePivnetClient) GetProductFilesCallCount() int {
	fake.getProductFilesMutex.RLock()
	defer fake.getProductFilesMutex.RUnlock()
	return len(fake.getProductFilesArgsForCall)
}

func (fake *FakePivnetClient) GetProductFilesArgsForCall(i int) (string, int) {
	fake.getProductFilesMutex.RLock()
	defer fake.getProductFilesMutex.RUnlock()
	return fake.getProductFilesArgsForCall[i].productSlug, fake.getProductFilesArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) GetProductFilesReturns(result1 []pivnet.ProductFile, result2 error) {
	fake.GetProductFilesStub = nil
	fake.getProductFilesReturns = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) GetProductFile(productSlug string, releaseID int, productFileID int) (pivnet.ProductFile, error) {
	fake.getProductFileMutex.Lock()
	fake.getProductFileArgsForCall = append(fake.getProductFileArgsForCall, struct {
		productSlug   string
		releaseID     int
		productFileID int
	}{productSlug, releaseID, productFileID})
	fake.recordInvocation("GetProductFile", []interface{}{productSlug, releaseID, productFileID})
	fake.getProductFileMutex.Unlock()
	if fake.GetProductFileStub != nil {
		return fake.GetProductFileStub(productSlug, releaseID, productFileID)
	} else {
		return fake.getProductFileReturns.result1, fake.getProductFileReturns.result2
	}
}

func (fake *FakePivnetClient) GetProductFileCallCount() int {
	fake.getProductFileMutex.RLock()
	defer fake.getProductFileMutex.RUnlock()
	return len(fake.getProductFileArgsForCall)
}

func (fake *FakePivnetClient) GetProductFileArgsForCall(i int) (string, int, int) {
	fake.getProductFileMutex.RLock()
	defer fake.getProductFileMutex.RUnlock()
	return fake.getProductFileArgsForCall[i].productSlug, fake.getProductFileArgsForCall[i].releaseID, fake.getProductFileArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) GetProductFileReturns(result1 pivnet.ProductFile, result2 error) {
	fake.GetProductFileStub = nil
	fake.getProductFileReturns = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseDependencies(productSlug string, releaseID int) ([]pivnet.ReleaseDependency, error) {
	fake.releaseDependenciesMutex.Lock()
	fake.releaseDependenciesArgsForCall = append(fake.releaseDependenciesArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("ReleaseDependencies", []interface{}{productSlug, releaseID})
	fake.releaseDependenciesMutex.Unlock()
	if fake.ReleaseDependenciesStub != nil {
		return fake.ReleaseDependenciesStub(productSlug, releaseID)
	} else {
		return fake.releaseDependenciesReturns.result1, fake.releaseDependenciesReturns.result2
	}
}

func (fake *FakePivnetClient) ReleaseDependenciesCallCount() int {
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	return len(fake.releaseDependenciesArgsForCall)
}

func (fake *FakePivnetClient) ReleaseDependenciesArgsForCall(i int) (string, int) {
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	return fake.releaseDependenciesArgsForCall[i].productSlug, fake.releaseDependenciesArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) ReleaseDependenciesReturns(result1 []pivnet.ReleaseDependency, result2 error) {
	fake.ReleaseDependenciesStub = nil
	fake.releaseDependenciesReturns = struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	fake.getProductFilesMutex.RLock()
	defer fake.getProductFilesMutex.RUnlock()
	fake.getProductFileMutex.RLock()
	defer fake.getProductFileMutex.RUnlock()
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
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
