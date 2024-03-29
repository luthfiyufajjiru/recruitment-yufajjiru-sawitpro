// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	v5 "github.com/golang-jwt/jwt/v5"
	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// ComparePassword mocks base method.
func (m *MockRepositoryInterface) ComparePassword(ctx context.Context, phone_number, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComparePassword", ctx, phone_number, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// ComparePassword indicates an expected call of ComparePassword.
func (mr *MockRepositoryInterfaceMockRecorder) ComparePassword(ctx, phone_number, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComparePassword", reflect.TypeOf((*MockRepositoryInterface)(nil).ComparePassword), ctx, phone_number, password)
}

// GenJWTTokens mocks base method.
func (m *MockRepositoryInterface) GenJWTTokens(user_id int, name string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenJWTTokens", user_id, name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenJWTTokens indicates an expected call of GenJWTTokens.
func (mr *MockRepositoryInterfaceMockRecorder) GenJWTTokens(user_id, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenJWTTokens", reflect.TypeOf((*MockRepositoryInterface)(nil).GenJWTTokens), user_id, name)
}

// GetClaims mocks base method.
func (m *MockRepositoryInterface) GetClaims(tokenStr string, key []byte) (v5.MapClaims, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClaims", tokenStr, key)
	ret0, _ := ret[0].(v5.MapClaims)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetClaims indicates an expected call of GetClaims.
func (mr *MockRepositoryInterfaceMockRecorder) GetClaims(tokenStr, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClaims", reflect.TypeOf((*MockRepositoryInterface)(nil).GetClaims), tokenStr, key)
}

// GetProfile mocks base method.
func (m *MockRepositoryInterface) GetProfile(ctx context.Context, user_id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", ctx, user_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockRepositoryInterfaceMockRecorder) GetProfile(ctx, user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockRepositoryInterface)(nil).GetProfile), ctx, user_id)
}

// GetTestById mocks base method.
func (m *MockRepositoryInterface) GetTestById(ctx context.Context, input GetTestByIdInput) (GetTestByIdOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestById", ctx, input)
	ret0, _ := ret[0].(GetTestByIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestById indicates an expected call of GetTestById.
func (mr *MockRepositoryInterfaceMockRecorder) GetTestById(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetTestById), ctx, input)
}

// RefreshToken mocks base method.
func (m *MockRepositoryInterface) RefreshToken(token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockRepositoryInterfaceMockRecorder) RefreshToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockRepositoryInterface)(nil).RefreshToken), token)
}

// SignJWTToken mocks base method.
func (m *MockRepositoryInterface) SignJWTToken(cl v5.MapClaims, key []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignJWTToken", cl, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignJWTToken indicates an expected call of SignJWTToken.
func (mr *MockRepositoryInterfaceMockRecorder) SignJWTToken(cl, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignJWTToken", reflect.TypeOf((*MockRepositoryInterface)(nil).SignJWTToken), cl, key)
}

// UpdateProfile mocks base method.
func (m *MockRepositoryInterface) UpdateProfile(ctx context.Context, user_id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", ctx, user_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateProfile(ctx, user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateProfile), ctx, user_id)
}
