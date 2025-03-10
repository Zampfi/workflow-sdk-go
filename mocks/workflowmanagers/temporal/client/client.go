// Code generated by mockery v2.50.4. DO NOT EDIT.

package mock_client

import (
	context "context"

	models "github.com/Zampfi/workflow-sdk-go/workflowmanagers/temporal/models"
	mock "github.com/stretchr/testify/mock"
)

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

type MockClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClient) EXPECT() *MockClient_Expecter {
	return &MockClient_Expecter{mock: &_m.Mock}
}

// CancelWorkflow provides a mock function with given fields: ctx, params
func (_m *MockClient) CancelWorkflow(ctx context.Context, params models.CancelWorkflowParams) error {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for CancelWorkflow")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.CancelWorkflowParams) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockClient_CancelWorkflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CancelWorkflow'
type MockClient_CancelWorkflow_Call struct {
	*mock.Call
}

// CancelWorkflow is a helper method to define mock.On call
//   - ctx context.Context
//   - params models.CancelWorkflowParams
func (_e *MockClient_Expecter) CancelWorkflow(ctx interface{}, params interface{}) *MockClient_CancelWorkflow_Call {
	return &MockClient_CancelWorkflow_Call{Call: _e.mock.On("CancelWorkflow", ctx, params)}
}

func (_c *MockClient_CancelWorkflow_Call) Run(run func(ctx context.Context, params models.CancelWorkflowParams)) *MockClient_CancelWorkflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.CancelWorkflowParams))
	})
	return _c
}

func (_c *MockClient_CancelWorkflow_Call) Return(_a0 error) *MockClient_CancelWorkflow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClient_CancelWorkflow_Call) RunAndReturn(run func(context.Context, models.CancelWorkflowParams) error) *MockClient_CancelWorkflow_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields: ctx
func (_m *MockClient) Close(ctx context.Context) {
	_m.Called(ctx)
}

// MockClient_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockClient_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockClient_Expecter) Close(ctx interface{}) *MockClient_Close_Call {
	return &MockClient_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *MockClient_Close_Call) Run(run func(ctx context.Context)) *MockClient_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockClient_Close_Call) Return() *MockClient_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockClient_Close_Call) RunAndReturn(run func(context.Context)) *MockClient_Close_Call {
	_c.Run(run)
	return _c
}

// Connect provides a mock function with given fields: ctx, params
func (_m *MockClient) Connect(ctx context.Context, params models.ConnectClientParams) error {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for Connect")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ConnectClientParams) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockClient_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type MockClient_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
//   - ctx context.Context
//   - params models.ConnectClientParams
func (_e *MockClient_Expecter) Connect(ctx interface{}, params interface{}) *MockClient_Connect_Call {
	return &MockClient_Connect_Call{Call: _e.mock.On("Connect", ctx, params)}
}

func (_c *MockClient_Connect_Call) Run(run func(ctx context.Context, params models.ConnectClientParams)) *MockClient_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.ConnectClientParams))
	})
	return _c
}

func (_c *MockClient_Connect_Call) Return(_a0 error) *MockClient_Connect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClient_Connect_Call) RunAndReturn(run func(context.Context, models.ConnectClientParams) error) *MockClient_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteAsyncWorkflow provides a mock function with given fields: ctx, params
func (_m *MockClient) ExecuteAsyncWorkflow(ctx context.Context, params models.ExecuteWorkflowParams) (models.WorkflowResponse, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteAsyncWorkflow")
	}

	var r0 models.WorkflowResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ExecuteWorkflowParams) (models.WorkflowResponse, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.ExecuteWorkflowParams) models.WorkflowResponse); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.WorkflowResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.ExecuteWorkflowParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_ExecuteAsyncWorkflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteAsyncWorkflow'
type MockClient_ExecuteAsyncWorkflow_Call struct {
	*mock.Call
}

// ExecuteAsyncWorkflow is a helper method to define mock.On call
//   - ctx context.Context
//   - params models.ExecuteWorkflowParams
func (_e *MockClient_Expecter) ExecuteAsyncWorkflow(ctx interface{}, params interface{}) *MockClient_ExecuteAsyncWorkflow_Call {
	return &MockClient_ExecuteAsyncWorkflow_Call{Call: _e.mock.On("ExecuteAsyncWorkflow", ctx, params)}
}

func (_c *MockClient_ExecuteAsyncWorkflow_Call) Run(run func(ctx context.Context, params models.ExecuteWorkflowParams)) *MockClient_ExecuteAsyncWorkflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.ExecuteWorkflowParams))
	})
	return _c
}

func (_c *MockClient_ExecuteAsyncWorkflow_Call) Return(_a0 models.WorkflowResponse, _a1 error) *MockClient_ExecuteAsyncWorkflow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_ExecuteAsyncWorkflow_Call) RunAndReturn(run func(context.Context, models.ExecuteWorkflowParams) (models.WorkflowResponse, error)) *MockClient_ExecuteAsyncWorkflow_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteScheduledWorkflow provides a mock function with given fields: ctx, params
func (_m *MockClient) ExecuteScheduledWorkflow(ctx context.Context, params models.ExecuteWorkflowWithScheduleParams) (models.ScheduledWorkflowResponse, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteScheduledWorkflow")
	}

	var r0 models.ScheduledWorkflowResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ExecuteWorkflowWithScheduleParams) (models.ScheduledWorkflowResponse, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.ExecuteWorkflowWithScheduleParams) models.ScheduledWorkflowResponse); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.ScheduledWorkflowResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.ExecuteWorkflowWithScheduleParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_ExecuteScheduledWorkflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteScheduledWorkflow'
type MockClient_ExecuteScheduledWorkflow_Call struct {
	*mock.Call
}

// ExecuteScheduledWorkflow is a helper method to define mock.On call
//   - ctx context.Context
//   - params models.ExecuteWorkflowWithScheduleParams
func (_e *MockClient_Expecter) ExecuteScheduledWorkflow(ctx interface{}, params interface{}) *MockClient_ExecuteScheduledWorkflow_Call {
	return &MockClient_ExecuteScheduledWorkflow_Call{Call: _e.mock.On("ExecuteScheduledWorkflow", ctx, params)}
}

func (_c *MockClient_ExecuteScheduledWorkflow_Call) Run(run func(ctx context.Context, params models.ExecuteWorkflowWithScheduleParams)) *MockClient_ExecuteScheduledWorkflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.ExecuteWorkflowWithScheduleParams))
	})
	return _c
}

func (_c *MockClient_ExecuteScheduledWorkflow_Call) Return(_a0 models.ScheduledWorkflowResponse, _a1 error) *MockClient_ExecuteScheduledWorkflow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_ExecuteScheduledWorkflow_Call) RunAndReturn(run func(context.Context, models.ExecuteWorkflowWithScheduleParams) (models.ScheduledWorkflowResponse, error)) *MockClient_ExecuteScheduledWorkflow_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteSyncWorkflow provides a mock function with given fields: ctx, params
func (_m *MockClient) ExecuteSyncWorkflow(ctx context.Context, params models.ExecuteWorkflowParams) (models.WorkflowResponse, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteSyncWorkflow")
	}

	var r0 models.WorkflowResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ExecuteWorkflowParams) (models.WorkflowResponse, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.ExecuteWorkflowParams) models.WorkflowResponse); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.WorkflowResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.ExecuteWorkflowParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_ExecuteSyncWorkflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteSyncWorkflow'
type MockClient_ExecuteSyncWorkflow_Call struct {
	*mock.Call
}

// ExecuteSyncWorkflow is a helper method to define mock.On call
//   - ctx context.Context
//   - params models.ExecuteWorkflowParams
func (_e *MockClient_Expecter) ExecuteSyncWorkflow(ctx interface{}, params interface{}) *MockClient_ExecuteSyncWorkflow_Call {
	return &MockClient_ExecuteSyncWorkflow_Call{Call: _e.mock.On("ExecuteSyncWorkflow", ctx, params)}
}

func (_c *MockClient_ExecuteSyncWorkflow_Call) Run(run func(ctx context.Context, params models.ExecuteWorkflowParams)) *MockClient_ExecuteSyncWorkflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.ExecuteWorkflowParams))
	})
	return _c
}

func (_c *MockClient_ExecuteSyncWorkflow_Call) Return(_a0 models.WorkflowResponse, _a1 error) *MockClient_ExecuteSyncWorkflow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_ExecuteSyncWorkflow_Call) RunAndReturn(run func(context.Context, models.ExecuteWorkflowParams) (models.WorkflowResponse, error)) *MockClient_ExecuteSyncWorkflow_Call {
	_c.Call.Return(run)
	return _c
}

// GetWorkflowDetails provides a mock function with given fields: ctx, params
func (_m *MockClient) GetWorkflowDetails(ctx context.Context, params models.GetWorkflowDetailsParams) (models.WorkflowDetailsResponse, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for GetWorkflowDetails")
	}

	var r0 models.WorkflowDetailsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.GetWorkflowDetailsParams) (models.WorkflowDetailsResponse, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.GetWorkflowDetailsParams) models.WorkflowDetailsResponse); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.WorkflowDetailsResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.GetWorkflowDetailsParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_GetWorkflowDetails_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWorkflowDetails'
type MockClient_GetWorkflowDetails_Call struct {
	*mock.Call
}

// GetWorkflowDetails is a helper method to define mock.On call
//   - ctx context.Context
//   - params models.GetWorkflowDetailsParams
func (_e *MockClient_Expecter) GetWorkflowDetails(ctx interface{}, params interface{}) *MockClient_GetWorkflowDetails_Call {
	return &MockClient_GetWorkflowDetails_Call{Call: _e.mock.On("GetWorkflowDetails", ctx, params)}
}

func (_c *MockClient_GetWorkflowDetails_Call) Run(run func(ctx context.Context, params models.GetWorkflowDetailsParams)) *MockClient_GetWorkflowDetails_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.GetWorkflowDetailsParams))
	})
	return _c
}

func (_c *MockClient_GetWorkflowDetails_Call) Return(_a0 models.WorkflowDetailsResponse, _a1 error) *MockClient_GetWorkflowDetails_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_GetWorkflowDetails_Call) RunAndReturn(run func(context.Context, models.GetWorkflowDetailsParams) (models.WorkflowDetailsResponse, error)) *MockClient_GetWorkflowDetails_Call {
	_c.Call.Return(run)
	return _c
}

// ListWorkflows provides a mock function with given fields: ctx, params
func (_m *MockClient) ListWorkflows(ctx context.Context, params models.ListWorkflowsParams) (models.ListWorkflowsResponse, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for ListWorkflows")
	}

	var r0 models.ListWorkflowsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ListWorkflowsParams) (models.ListWorkflowsResponse, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.ListWorkflowsParams) models.ListWorkflowsResponse); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.ListWorkflowsResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.ListWorkflowsParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_ListWorkflows_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListWorkflows'
type MockClient_ListWorkflows_Call struct {
	*mock.Call
}

// ListWorkflows is a helper method to define mock.On call
//   - ctx context.Context
//   - params models.ListWorkflowsParams
func (_e *MockClient_Expecter) ListWorkflows(ctx interface{}, params interface{}) *MockClient_ListWorkflows_Call {
	return &MockClient_ListWorkflows_Call{Call: _e.mock.On("ListWorkflows", ctx, params)}
}

func (_c *MockClient_ListWorkflows_Call) Run(run func(ctx context.Context, params models.ListWorkflowsParams)) *MockClient_ListWorkflows_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.ListWorkflowsParams))
	})
	return _c
}

func (_c *MockClient_ListWorkflows_Call) Return(_a0 models.ListWorkflowsResponse, _a1 error) *MockClient_ListWorkflows_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_ListWorkflows_Call) RunAndReturn(run func(context.Context, models.ListWorkflowsParams) (models.ListWorkflowsResponse, error)) *MockClient_ListWorkflows_Call {
	_c.Call.Return(run)
	return _c
}

// QuerySchedule provides a mock function with given fields: ctx, params
func (_m *MockClient) QuerySchedule(ctx context.Context, params models.QueryScheduleParams) (models.QueryScheduleResponse, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for QuerySchedule")
	}

	var r0 models.QueryScheduleResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.QueryScheduleParams) (models.QueryScheduleResponse, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.QueryScheduleParams) models.QueryScheduleResponse); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.QueryScheduleResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.QueryScheduleParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_QuerySchedule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QuerySchedule'
type MockClient_QuerySchedule_Call struct {
	*mock.Call
}

// QuerySchedule is a helper method to define mock.On call
//   - ctx context.Context
//   - params models.QueryScheduleParams
func (_e *MockClient_Expecter) QuerySchedule(ctx interface{}, params interface{}) *MockClient_QuerySchedule_Call {
	return &MockClient_QuerySchedule_Call{Call: _e.mock.On("QuerySchedule", ctx, params)}
}

func (_c *MockClient_QuerySchedule_Call) Run(run func(ctx context.Context, params models.QueryScheduleParams)) *MockClient_QuerySchedule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.QueryScheduleParams))
	})
	return _c
}

func (_c *MockClient_QuerySchedule_Call) Return(_a0 models.QueryScheduleResponse, _a1 error) *MockClient_QuerySchedule_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_QuerySchedule_Call) RunAndReturn(run func(context.Context, models.QueryScheduleParams) (models.QueryScheduleResponse, error)) *MockClient_QuerySchedule_Call {
	_c.Call.Return(run)
	return _c
}

// QueryWorkflow provides a mock function with given fields: ctx, params
func (_m *MockClient) QueryWorkflow(ctx context.Context, params models.QueryWorkflowParams) (models.QueryWorkflowResponse, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for QueryWorkflow")
	}

	var r0 models.QueryWorkflowResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.QueryWorkflowParams) (models.QueryWorkflowResponse, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.QueryWorkflowParams) models.QueryWorkflowResponse); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.QueryWorkflowResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.QueryWorkflowParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_QueryWorkflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryWorkflow'
type MockClient_QueryWorkflow_Call struct {
	*mock.Call
}

// QueryWorkflow is a helper method to define mock.On call
//   - ctx context.Context
//   - params models.QueryWorkflowParams
func (_e *MockClient_Expecter) QueryWorkflow(ctx interface{}, params interface{}) *MockClient_QueryWorkflow_Call {
	return &MockClient_QueryWorkflow_Call{Call: _e.mock.On("QueryWorkflow", ctx, params)}
}

func (_c *MockClient_QueryWorkflow_Call) Run(run func(ctx context.Context, params models.QueryWorkflowParams)) *MockClient_QueryWorkflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.QueryWorkflowParams))
	})
	return _c
}

func (_c *MockClient_QueryWorkflow_Call) Return(_a0 models.QueryWorkflowResponse, _a1 error) *MockClient_QueryWorkflow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_QueryWorkflow_Call) RunAndReturn(run func(context.Context, models.QueryWorkflowParams) (models.QueryWorkflowResponse, error)) *MockClient_QueryWorkflow_Call {
	_c.Call.Return(run)
	return _c
}

// SignalWorkflow provides a mock function with given fields: ctx, params
func (_m *MockClient) SignalWorkflow(ctx context.Context, params models.SignalWorkflowParams) error {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for SignalWorkflow")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SignalWorkflowParams) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockClient_SignalWorkflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignalWorkflow'
type MockClient_SignalWorkflow_Call struct {
	*mock.Call
}

// SignalWorkflow is a helper method to define mock.On call
//   - ctx context.Context
//   - params models.SignalWorkflowParams
func (_e *MockClient_Expecter) SignalWorkflow(ctx interface{}, params interface{}) *MockClient_SignalWorkflow_Call {
	return &MockClient_SignalWorkflow_Call{Call: _e.mock.On("SignalWorkflow", ctx, params)}
}

func (_c *MockClient_SignalWorkflow_Call) Run(run func(ctx context.Context, params models.SignalWorkflowParams)) *MockClient_SignalWorkflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.SignalWorkflowParams))
	})
	return _c
}

func (_c *MockClient_SignalWorkflow_Call) Return(_a0 error) *MockClient_SignalWorkflow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClient_SignalWorkflow_Call) RunAndReturn(run func(context.Context, models.SignalWorkflowParams) error) *MockClient_SignalWorkflow_Call {
	_c.Call.Return(run)
	return _c
}

// TerminateWorkflow provides a mock function with given fields: ctx, params
func (_m *MockClient) TerminateWorkflow(ctx context.Context, params models.TerminateWorkflowParams) error {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for TerminateWorkflow")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.TerminateWorkflowParams) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockClient_TerminateWorkflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TerminateWorkflow'
type MockClient_TerminateWorkflow_Call struct {
	*mock.Call
}

// TerminateWorkflow is a helper method to define mock.On call
//   - ctx context.Context
//   - params models.TerminateWorkflowParams
func (_e *MockClient_Expecter) TerminateWorkflow(ctx interface{}, params interface{}) *MockClient_TerminateWorkflow_Call {
	return &MockClient_TerminateWorkflow_Call{Call: _e.mock.On("TerminateWorkflow", ctx, params)}
}

func (_c *MockClient_TerminateWorkflow_Call) Run(run func(ctx context.Context, params models.TerminateWorkflowParams)) *MockClient_TerminateWorkflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.TerminateWorkflowParams))
	})
	return _c
}

func (_c *MockClient_TerminateWorkflow_Call) Return(_a0 error) *MockClient_TerminateWorkflow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClient_TerminateWorkflow_Call) RunAndReturn(run func(context.Context, models.TerminateWorkflowParams) error) *MockClient_TerminateWorkflow_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
