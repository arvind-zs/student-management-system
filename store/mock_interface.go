// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package store is a generated GoMock package.
package store

import (
	context "context"
	reflect "reflect"
	models "student-management-system/models"

	gomock "github.com/golang/mock/gomock"
)

// MockStudent is a mock of Student interface.
type MockStudent struct {
	ctrl     *gomock.Controller
	recorder *MockStudentMockRecorder
}

// MockStudentMockRecorder is the mock recorder for MockStudent.
type MockStudentMockRecorder struct {
	mock *MockStudent
}

// NewMockStudent creates a new mock instance.
func NewMockStudent(ctrl *gomock.Controller) *MockStudent {
	mock := &MockStudent{ctrl: ctrl}
	mock.recorder = &MockStudentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudent) EXPECT() *MockStudentMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockStudent) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStudentMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStudent)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockStudent) Get(ctx context.Context) ([]models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].([]models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStudentMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStudent)(nil).Get), ctx)
}

// GetByFirstAndLastName mocks base method.
func (m *MockStudent) GetByFirstAndLastName(ctx context.Context, firstName, lastName string) ([]models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByFirstAndLastName", ctx, firstName, lastName)
	ret0, _ := ret[0].([]models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByFirstAndLastName indicates an expected call of GetByFirstAndLastName.
func (mr *MockStudentMockRecorder) GetByFirstAndLastName(ctx, firstName, lastName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByFirstAndLastName", reflect.TypeOf((*MockStudent)(nil).GetByFirstAndLastName), ctx, firstName, lastName)
}

// GetByFirstName mocks base method.
func (m *MockStudent) GetByFirstName(ctx context.Context, firstName string) ([]models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByFirstName", ctx, firstName)
	ret0, _ := ret[0].([]models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByFirstName indicates an expected call of GetByFirstName.
func (mr *MockStudentMockRecorder) GetByFirstName(ctx, firstName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByFirstName", reflect.TypeOf((*MockStudent)(nil).GetByFirstName), ctx, firstName)
}

// GetByID mocks base method.
func (m *MockStudent) GetByID(ctx context.Context, id int) (models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockStudentMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockStudent)(nil).GetByID), ctx, id)
}

// GetByLastName mocks base method.
func (m *MockStudent) GetByLastName(ctx context.Context, lastName string) ([]models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByLastName", ctx, lastName)
	ret0, _ := ret[0].([]models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByLastName indicates an expected call of GetByLastName.
func (mr *MockStudentMockRecorder) GetByLastName(ctx, lastName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByLastName", reflect.TypeOf((*MockStudent)(nil).GetByLastName), ctx, lastName)
}

// Post mocks base method.
func (m *MockStudent) Post(ctx context.Context, student *models.Student) (models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", ctx, student)
	ret0, _ := ret[0].(models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockStudentMockRecorder) Post(ctx, student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockStudent)(nil).Post), ctx, student)
}

// Put mocks base method.
func (m *MockStudent) Put(ctx context.Context, id int, student *models.Student) (models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", ctx, id, student)
	ret0, _ := ret[0].(models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockStudentMockRecorder) Put(ctx, id, student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockStudent)(nil).Put), ctx, id, student)
}
