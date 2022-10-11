// Code generated by MockGen. DO NOT EDIT.
// Source: ./vending-machine/domain/product_crud.go

// Package mock_domain is a generated GoMock package.
package domain

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockProductCrud is a mock of ProductCrud interface.
type MockProductCrud struct {
	ctrl     *gomock.Controller
	recorder *MockProductCrudMockRecorder
}

// MockProductCrudMockRecorder is the mock recorder for MockProductCrud.
type MockProductCrudMockRecorder struct {
	mock *MockProductCrud
}

// NewMockProductCrud creates a new mock instance.
func NewMockProductCrud(ctrl *gomock.Controller) *MockProductCrud {
	mock := &MockProductCrud{ctrl: ctrl}
	mock.recorder = &MockProductCrudMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductCrud) EXPECT() *MockProductCrudMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockProductCrud) CreateProduct(product Product, tx *sqlx.Tx) (*Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", product, tx)
	ret0, _ := ret[0].(*Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductCrudMockRecorder) CreateProduct(product, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProductCrud)(nil).CreateProduct), product, tx)
}

// DeleteProduct mocks base method.
func (m *MockProductCrud) DeleteProduct(product Product, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", product, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockProductCrudMockRecorder) DeleteProduct(product, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockProductCrud)(nil).DeleteProduct), product, tx)
}

// GetAllProducts mocks base method.
func (m *MockProductCrud) GetAllProducts(tx *sqlx.Tx) ([]Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProducts", tx)
	ret0, _ := ret[0].([]Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProducts indicates an expected call of GetAllProducts.
func (mr *MockProductCrudMockRecorder) GetAllProducts(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProducts", reflect.TypeOf((*MockProductCrud)(nil).GetAllProducts), tx)
}

// GetProductByID mocks base method.
func (m *MockProductCrud) GetProductByID(product Product, tx *sqlx.Tx) (*Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductByID", product, tx)
	ret0, _ := ret[0].(*Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductByID indicates an expected call of GetProductByID.
func (mr *MockProductCrudMockRecorder) GetProductByID(product, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductByID", reflect.TypeOf((*MockProductCrud)(nil).GetProductByID), product, tx)
}

// UpdateProduct mocks base method.
func (m *MockProductCrud) UpdateProduct(product Product, tx *sqlx.Tx) (*Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", product, tx)
	ret0, _ := ret[0].(*Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockProductCrudMockRecorder) UpdateProduct(product, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProductCrud)(nil).UpdateProduct), product, tx)
}
