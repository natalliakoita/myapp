// Code generated by MockGen. DO NOT EDIT.
// Source: service/book_service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	model "myapp/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBookServiceInterface is a mock of BookServiceInterface interface.
type MockBookServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockBookServiceInterfaceMockRecorder
}

// MockBookServiceInterfaceMockRecorder is the mock recorder for MockBookServiceInterface.
type MockBookServiceInterfaceMockRecorder struct {
	mock *MockBookServiceInterface
}

// NewMockBookServiceInterface creates a new mock instance.
func NewMockBookServiceInterface(ctrl *gomock.Controller) *MockBookServiceInterface {
	mock := &MockBookServiceInterface{ctrl: ctrl}
	mock.recorder = &MockBookServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookServiceInterface) EXPECT() *MockBookServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockBookServiceInterface) CreateBook(ctx context.Context, book *model.BookForm) (*model.BookDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", ctx, book)
	ret0, _ := ret[0].(*model.BookDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockBookServiceInterfaceMockRecorder) CreateBook(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockBookServiceInterface)(nil).CreateBook), ctx, book)
}

// DeleteBook mocks base method.
func (m *MockBookServiceInterface) DeleteBook(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockBookServiceInterfaceMockRecorder) DeleteBook(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockBookServiceInterface)(nil).DeleteBook), ctx, id)
}

// GetBookByID mocks base method.
func (m *MockBookServiceInterface) GetBookByID(ctx context.Context, id uint) (*model.BookDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookByID", ctx, id)
	ret0, _ := ret[0].(*model.BookDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByID indicates an expected call of GetBookByID.
func (mr *MockBookServiceInterfaceMockRecorder) GetBookByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByID", reflect.TypeOf((*MockBookServiceInterface)(nil).GetBookByID), ctx, id)
}

// GetListBook mocks base method.
func (m *MockBookServiceInterface) GetListBook(ctx context.Context) ([]model.BookDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListBook", ctx)
	ret0, _ := ret[0].([]model.BookDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListBook indicates an expected call of GetListBook.
func (mr *MockBookServiceInterfaceMockRecorder) GetListBook(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListBook", reflect.TypeOf((*MockBookServiceInterface)(nil).GetListBook), ctx)
}

// UpdateBook mocks base method.
func (m *MockBookServiceInterface) UpdateBook(ctx context.Context, id uint, book *model.BookForm) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", ctx, id, book)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockBookServiceInterfaceMockRecorder) UpdateBook(ctx, id, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockBookServiceInterface)(nil).UpdateBook), ctx, id, book)
}
