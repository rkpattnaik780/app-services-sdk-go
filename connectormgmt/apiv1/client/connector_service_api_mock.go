// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package connectormgmtclient

import (
	_context "context"
	_nethttp "net/http"
	"sync"
)

// Ensure, that ConnectorServiceApiMock does implement ConnectorServiceApi.
// If this is not the case, regenerate this file with moq.
var _ ConnectorServiceApi = &ConnectorServiceApiMock{}

// ConnectorServiceApiMock is a mock implementation of ConnectorServiceApi.
//
// 	func TestSomethingThatUsesConnectorServiceApi(t *testing.T) {
//
// 		// make and configure a mocked ConnectorServiceApi
// 		mockedConnectorServiceApi := &ConnectorServiceApiMock{
// 			GetVersionMetadataFunc: func(ctx _context.Context) ApiGetVersionMetadataRequest {
// 				panic("mock out the GetVersionMetadata method")
// 			},
// 			GetVersionMetadataExecuteFunc: func(r ApiGetVersionMetadataRequest) (VersionMetadata, *_nethttp.Response, error) {
// 				panic("mock out the GetVersionMetadataExecute method")
// 			},
// 		}
//
// 		// use mockedConnectorServiceApi in code that requires ConnectorServiceApi
// 		// and then make assertions.
//
// 	}
type ConnectorServiceApiMock struct {
	// GetVersionMetadataFunc mocks the GetVersionMetadata method.
	GetVersionMetadataFunc func(ctx _context.Context) ApiGetVersionMetadataRequest

	// GetVersionMetadataExecuteFunc mocks the GetVersionMetadataExecute method.
	GetVersionMetadataExecuteFunc func(r ApiGetVersionMetadataRequest) (VersionMetadata, *_nethttp.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetVersionMetadata holds details about calls to the GetVersionMetadata method.
		GetVersionMetadata []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
		}
		// GetVersionMetadataExecute holds details about calls to the GetVersionMetadataExecute method.
		GetVersionMetadataExecute []struct {
			// R is the r argument value.
			R ApiGetVersionMetadataRequest
		}
	}
	lockGetVersionMetadata        sync.RWMutex
	lockGetVersionMetadataExecute sync.RWMutex
}

// GetVersionMetadata calls GetVersionMetadataFunc.
func (mock *ConnectorServiceApiMock) GetVersionMetadata(ctx _context.Context) ApiGetVersionMetadataRequest {
	if mock.GetVersionMetadataFunc == nil {
		panic("ConnectorServiceApiMock.GetVersionMetadataFunc: method is nil but ConnectorServiceApi.GetVersionMetadata was just called")
	}
	callInfo := struct {
		Ctx _context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetVersionMetadata.Lock()
	mock.calls.GetVersionMetadata = append(mock.calls.GetVersionMetadata, callInfo)
	mock.lockGetVersionMetadata.Unlock()
	return mock.GetVersionMetadataFunc(ctx)
}

// GetVersionMetadataCalls gets all the calls that were made to GetVersionMetadata.
// Check the length with:
//     len(mockedConnectorServiceApi.GetVersionMetadataCalls())
func (mock *ConnectorServiceApiMock) GetVersionMetadataCalls() []struct {
	Ctx _context.Context
} {
	var calls []struct {
		Ctx _context.Context
	}
	mock.lockGetVersionMetadata.RLock()
	calls = mock.calls.GetVersionMetadata
	mock.lockGetVersionMetadata.RUnlock()
	return calls
}

// GetVersionMetadataExecute calls GetVersionMetadataExecuteFunc.
func (mock *ConnectorServiceApiMock) GetVersionMetadataExecute(r ApiGetVersionMetadataRequest) (VersionMetadata, *_nethttp.Response, error) {
	if mock.GetVersionMetadataExecuteFunc == nil {
		panic("ConnectorServiceApiMock.GetVersionMetadataExecuteFunc: method is nil but ConnectorServiceApi.GetVersionMetadataExecute was just called")
	}
	callInfo := struct {
		R ApiGetVersionMetadataRequest
	}{
		R: r,
	}
	mock.lockGetVersionMetadataExecute.Lock()
	mock.calls.GetVersionMetadataExecute = append(mock.calls.GetVersionMetadataExecute, callInfo)
	mock.lockGetVersionMetadataExecute.Unlock()
	return mock.GetVersionMetadataExecuteFunc(r)
}

// GetVersionMetadataExecuteCalls gets all the calls that were made to GetVersionMetadataExecute.
// Check the length with:
//     len(mockedConnectorServiceApi.GetVersionMetadataExecuteCalls())
func (mock *ConnectorServiceApiMock) GetVersionMetadataExecuteCalls() []struct {
	R ApiGetVersionMetadataRequest
} {
	var calls []struct {
		R ApiGetVersionMetadataRequest
	}
	mock.lockGetVersionMetadataExecute.RLock()
	calls = mock.calls.GetVersionMetadataExecute
	mock.lockGetVersionMetadataExecute.RUnlock()
	return calls
}
