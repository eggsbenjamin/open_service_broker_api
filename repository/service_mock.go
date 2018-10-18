// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package repository is a generated GoMock package.
package repository

import (
	models "github.com/eggsbenjamin/open_service_broker_api/models"
	uuid "github.com/eggsbenjamin/open_service_broker_api/uuid"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockServiceRepository is a mock of ServiceRepository interface
type MockServiceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockServiceRepositoryMockRecorder
}

// MockServiceRepositoryMockRecorder is the mock recorder for MockServiceRepository
type MockServiceRepositoryMockRecorder struct {
	mock *MockServiceRepository
}

// NewMockServiceRepository creates a new mock instance
func NewMockServiceRepository(ctrl *gomock.Controller) *MockServiceRepository {
	mock := &MockServiceRepository{ctrl: ctrl}
	mock.recorder = &MockServiceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceRepository) EXPECT() *MockServiceRepositoryMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockServiceRepository) GetAll() ([]*models.DBService, error) {
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*models.DBService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockServiceRepositoryMockRecorder) GetAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockServiceRepository)(nil).GetAll))
}

// GetByServiceID mocks base method
func (m *MockServiceRepository) GetByServiceID(arg0 uuid.UUID) (*models.DBService, error) {
	ret := m.ctrl.Call(m, "GetByServiceID", arg0)
	ret0, _ := ret[0].(*models.DBService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByServiceID indicates an expected call of GetByServiceID
func (mr *MockServiceRepositoryMockRecorder) GetByServiceID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByServiceID", reflect.TypeOf((*MockServiceRepository)(nil).GetByServiceID), arg0)
}

// Create mocks base method
func (m *MockServiceRepository) Create(arg0 *models.DBService) error {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockServiceRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServiceRepository)(nil).Create), arg0)
}
