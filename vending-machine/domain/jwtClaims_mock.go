// Code generated by MockGen. DO NOT EDIT.
// Source: ./vending-machine/domain/jwtClaims.go

// Package mock_domain is a generated GoMock package.
package domain

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockJwtClaimsInterface is a mock of JwtClaimsInterface interface.
type MockJwtClaimsInterface struct {
	ctrl     *gomock.Controller
	recorder *MockJwtClaimsInterfaceMockRecorder
}

// MockJwtClaimsInterfaceMockRecorder is the mock recorder for MockJwtClaimsInterface.
type MockJwtClaimsInterfaceMockRecorder struct {
	mock *MockJwtClaimsInterface
}

// NewMockJwtClaimsInterface creates a new mock instance.
func NewMockJwtClaimsInterface(ctrl *gomock.Controller) *MockJwtClaimsInterface {
	mock := &MockJwtClaimsInterface{ctrl: ctrl}
	mock.recorder = &MockJwtClaimsInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJwtClaimsInterface) EXPECT() *MockJwtClaimsInterfaceMockRecorder {
	return m.recorder
}

// CreateToken mocks base method.
func (m *MockJwtClaimsInterface) CreateToken(sub string, user User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToken", sub, user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateToken indicates an expected call of CreateToken.
func (mr *MockJwtClaimsInterfaceMockRecorder) CreateToken(sub, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockJwtClaimsInterface)(nil).CreateToken), sub, user)
}

// GetClaimsFromToken mocks base method.
func (m *MockJwtClaimsInterface) GetClaimsFromToken(tokenString string) (*JwtClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClaimsFromToken", tokenString)
	ret0, _ := ret[0].(*JwtClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClaimsFromToken indicates an expected call of GetClaimsFromToken.
func (mr *MockJwtClaimsInterfaceMockRecorder) GetClaimsFromToken(tokenString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClaimsFromToken", reflect.TypeOf((*MockJwtClaimsInterface)(nil).GetClaimsFromToken), tokenString)
}

// SetJWTClaimsContext mocks base method.
func (m *MockJwtClaimsInterface) SetJWTClaimsContext(ctx context.Context, claims JwtClaims) context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetJWTClaimsContext", ctx, claims)
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// SetJWTClaimsContext indicates an expected call of SetJWTClaimsContext.
func (mr *MockJwtClaimsInterfaceMockRecorder) SetJWTClaimsContext(ctx, claims interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetJWTClaimsContext", reflect.TypeOf((*MockJwtClaimsInterface)(nil).SetJWTClaimsContext), ctx, claims)
}
