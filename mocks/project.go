// Code generated by MockGen. DO NOT EDIT.
// Source: project.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	circleci "github.com/grezar/go-circleci"
)

// MockProjects is a mock of Projects interface.
type MockProjects struct {
	ctrl     *gomock.Controller
	recorder *MockProjectsMockRecorder
}

// MockProjectsMockRecorder is the mock recorder for MockProjects.
type MockProjectsMockRecorder struct {
	mock *MockProjects
}

// NewMockProjects creates a new mock instance.
func NewMockProjects(ctrl *gomock.Controller) *MockProjects {
	mock := &MockProjects{ctrl: ctrl}
	mock.recorder = &MockProjectsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjects) EXPECT() *MockProjectsMockRecorder {
	return m.recorder
}

// CreateCheckoutKey mocks base method.
func (m *MockProjects) CreateCheckoutKey(ctx context.Context, projectSlug string, options circleci.ProjectCreateCheckoutKeyOptions) (*circleci.ProjectCheckoutKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCheckoutKey", ctx, projectSlug, options)
	ret0, _ := ret[0].(*circleci.ProjectCheckoutKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCheckoutKey indicates an expected call of CreateCheckoutKey.
func (mr *MockProjectsMockRecorder) CreateCheckoutKey(ctx, projectSlug, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCheckoutKey", reflect.TypeOf((*MockProjects)(nil).CreateCheckoutKey), ctx, projectSlug, options)
}

// CreateVariable mocks base method.
func (m *MockProjects) CreateVariable(ctx context.Context, projectSlug string, options circleci.ProjectCreateVariableOptions) (*circleci.ProjectVariable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVariable", ctx, projectSlug, options)
	ret0, _ := ret[0].(*circleci.ProjectVariable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVariable indicates an expected call of CreateVariable.
func (mr *MockProjectsMockRecorder) CreateVariable(ctx, projectSlug, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVariable", reflect.TypeOf((*MockProjects)(nil).CreateVariable), ctx, projectSlug, options)
}

// DeleteCheckoutKey mocks base method.
func (m *MockProjects) DeleteCheckoutKey(ctx context.Context, projectSlug, fingerprint string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCheckoutKey", ctx, projectSlug, fingerprint)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCheckoutKey indicates an expected call of DeleteCheckoutKey.
func (mr *MockProjectsMockRecorder) DeleteCheckoutKey(ctx, projectSlug, fingerprint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCheckoutKey", reflect.TypeOf((*MockProjects)(nil).DeleteCheckoutKey), ctx, projectSlug, fingerprint)
}

// DeleteVariable mocks base method.
func (m *MockProjects) DeleteVariable(ctx context.Context, projectSlug, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVariable", ctx, projectSlug, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVariable indicates an expected call of DeleteVariable.
func (mr *MockProjectsMockRecorder) DeleteVariable(ctx, projectSlug, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVariable", reflect.TypeOf((*MockProjects)(nil).DeleteVariable), ctx, projectSlug, name)
}

// Get mocks base method.
func (m *MockProjects) Get(ctx context.Context, projectSlug string) (*circleci.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, projectSlug)
	ret0, _ := ret[0].(*circleci.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProjectsMockRecorder) Get(ctx, projectSlug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProjects)(nil).Get), ctx, projectSlug)
}

// GetCheckoutKey mocks base method.
func (m *MockProjects) GetCheckoutKey(ctx context.Context, projectSlug, fingerprint string) (*circleci.ProjectCheckoutKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckoutKey", ctx, projectSlug, fingerprint)
	ret0, _ := ret[0].(*circleci.ProjectCheckoutKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckoutKey indicates an expected call of GetCheckoutKey.
func (mr *MockProjectsMockRecorder) GetCheckoutKey(ctx, projectSlug, fingerprint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckoutKey", reflect.TypeOf((*MockProjects)(nil).GetCheckoutKey), ctx, projectSlug, fingerprint)
}

// GetPipeline mocks base method.
func (m *MockProjects) GetPipeline(ctx context.Context, projectSlug, pipelineNumber string) (*circleci.Pipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPipeline", ctx, projectSlug, pipelineNumber)
	ret0, _ := ret[0].(*circleci.Pipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPipeline indicates an expected call of GetPipeline.
func (mr *MockProjectsMockRecorder) GetPipeline(ctx, projectSlug, pipelineNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipeline", reflect.TypeOf((*MockProjects)(nil).GetPipeline), ctx, projectSlug, pipelineNumber)
}

// GetVariable mocks base method.
func (m *MockProjects) GetVariable(ctx context.Context, projectSlug, name string) (*circleci.ProjectVariable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVariable", ctx, projectSlug, name)
	ret0, _ := ret[0].(*circleci.ProjectVariable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVariable indicates an expected call of GetVariable.
func (mr *MockProjectsMockRecorder) GetVariable(ctx, projectSlug, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVariable", reflect.TypeOf((*MockProjects)(nil).GetVariable), ctx, projectSlug, name)
}

// ListCheckoutKeys mocks base method.
func (m *MockProjects) ListCheckoutKeys(ctx context.Context, projectSlug string) (*circleci.ProjectCheckoutKeyList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCheckoutKeys", ctx, projectSlug)
	ret0, _ := ret[0].(*circleci.ProjectCheckoutKeyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCheckoutKeys indicates an expected call of ListCheckoutKeys.
func (mr *MockProjectsMockRecorder) ListCheckoutKeys(ctx, projectSlug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCheckoutKeys", reflect.TypeOf((*MockProjects)(nil).ListCheckoutKeys), ctx, projectSlug)
}

// ListMyPipelines mocks base method.
func (m *MockProjects) ListMyPipelines(ctx context.Context, projectSlug string, options circleci.ProjectListMyPipelinesOptions) (*circleci.PipelineList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMyPipelines", ctx, projectSlug, options)
	ret0, _ := ret[0].(*circleci.PipelineList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMyPipelines indicates an expected call of ListMyPipelines.
func (mr *MockProjectsMockRecorder) ListMyPipelines(ctx, projectSlug, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMyPipelines", reflect.TypeOf((*MockProjects)(nil).ListMyPipelines), ctx, projectSlug, options)
}

// ListPipelines mocks base method.
func (m *MockProjects) ListPipelines(ctx context.Context, projectSlug string, options circleci.ProjectListPipelinesOptions) (*circleci.PipelineList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPipelines", ctx, projectSlug, options)
	ret0, _ := ret[0].(*circleci.PipelineList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPipelines indicates an expected call of ListPipelines.
func (mr *MockProjectsMockRecorder) ListPipelines(ctx, projectSlug, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipelines", reflect.TypeOf((*MockProjects)(nil).ListPipelines), ctx, projectSlug, options)
}

// ListVariables mocks base method.
func (m *MockProjects) ListVariables(ctx context.Context, projectSlug string) (*circleci.ProjectVariableList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVariables", ctx, projectSlug)
	ret0, _ := ret[0].(*circleci.ProjectVariableList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVariables indicates an expected call of ListVariables.
func (mr *MockProjectsMockRecorder) ListVariables(ctx, projectSlug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVariables", reflect.TypeOf((*MockProjects)(nil).ListVariables), ctx, projectSlug)
}

// TriggerPipeline mocks base method.
func (m *MockProjects) TriggerPipeline(ctx context.Context, projectSlug string, options circleci.ProjectTriggerPipelineOptions) (*circleci.Pipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerPipeline", ctx, projectSlug, options)
	ret0, _ := ret[0].(*circleci.Pipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerPipeline indicates an expected call of TriggerPipeline.
func (mr *MockProjectsMockRecorder) TriggerPipeline(ctx, projectSlug, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerPipeline", reflect.TypeOf((*MockProjects)(nil).TriggerPipeline), ctx, projectSlug, options)
}
