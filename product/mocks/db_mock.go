// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/ports/db.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=./internal/ports/db.go -destination=./mocks/db_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	domain "product/internal/application/core/domain"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockDBPort is a mock of DBPort interface.
type MockDBPort struct {
	ctrl     *gomock.Controller
	recorder *MockDBPortMockRecorder
}

// MockDBPortMockRecorder is the mock recorder for MockDBPort.
type MockDBPortMockRecorder struct {
	mock *MockDBPort
}

// NewMockDBPort creates a new mock instance.
func NewMockDBPort(ctrl *gomock.Controller) *MockDBPort {
	mock := &MockDBPort{ctrl: ctrl}
	mock.recorder = &MockDBPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBPort) EXPECT() *MockDBPortMockRecorder {
	return m.recorder
}

// AddAuthor mocks base method.
func (m *MockDBPort) AddAuthor(ctx context.Context, author *domain.Author) (*domain.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAuthor", ctx, author)
	ret0, _ := ret[0].(*domain.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAuthor indicates an expected call of AddAuthor.
func (mr *MockDBPortMockRecorder) AddAuthor(ctx, author any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAuthor", reflect.TypeOf((*MockDBPort)(nil).AddAuthor), ctx, author)
}

// AddBook mocks base method.
func (m *MockDBPort) AddBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBook", ctx, book)
	ret0, _ := ret[0].(*domain.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBook indicates an expected call of AddBook.
func (mr *MockDBPortMockRecorder) AddBook(ctx, book any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBook", reflect.TypeOf((*MockDBPort)(nil).AddBook), ctx, book)
}

// AddGenre mocks base method.
func (m *MockDBPort) AddGenre(ctx context.Context, genre *domain.Genre) (*domain.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddGenre", ctx, genre)
	ret0, _ := ret[0].(*domain.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddGenre indicates an expected call of AddGenre.
func (mr *MockDBPortMockRecorder) AddGenre(ctx, genre any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGenre", reflect.TypeOf((*MockDBPort)(nil).AddGenre), ctx, genre)
}

// DeleteBookByID mocks base method.
func (m *MockDBPort) DeleteBookByID(ctx context.Context, ID domain.BookID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBookByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBookByID indicates an expected call of DeleteBookByID.
func (mr *MockDBPortMockRecorder) DeleteBookByID(ctx, ID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBookByID", reflect.TypeOf((*MockDBPort)(nil).DeleteBookByID), ctx, ID)
}

// GetAllAuthors mocks base method.
func (m *MockDBPort) GetAllAuthors(ctx context.Context) ([]domain.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAuthors", ctx)
	ret0, _ := ret[0].([]domain.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAuthors indicates an expected call of GetAllAuthors.
func (mr *MockDBPortMockRecorder) GetAllAuthors(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAuthors", reflect.TypeOf((*MockDBPort)(nil).GetAllAuthors), ctx)
}

// GetAllBooks mocks base method.
func (m *MockDBPort) GetAllBooks(ctx context.Context) ([]domain.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBooks", ctx)
	ret0, _ := ret[0].([]domain.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBooks indicates an expected call of GetAllBooks.
func (mr *MockDBPortMockRecorder) GetAllBooks(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBooks", reflect.TypeOf((*MockDBPort)(nil).GetAllBooks), ctx)
}

// GetAllGenres mocks base method.
func (m *MockDBPort) GetAllGenres(ctx context.Context) ([]domain.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllGenres", ctx)
	ret0, _ := ret[0].([]domain.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllGenres indicates an expected call of GetAllGenres.
func (mr *MockDBPortMockRecorder) GetAllGenres(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllGenres", reflect.TypeOf((*MockDBPort)(nil).GetAllGenres), ctx)
}

// GetBookByID mocks base method.
func (m *MockDBPort) GetBookByID(ctx context.Context, ID domain.BookID) (*domain.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookByID", ctx, ID)
	ret0, _ := ret[0].(*domain.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByID indicates an expected call of GetBookByID.
func (mr *MockDBPortMockRecorder) GetBookByID(ctx, ID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByID", reflect.TypeOf((*MockDBPort)(nil).GetBookByID), ctx, ID)
}

// GetBooksByAuthor mocks base method.
func (m *MockDBPort) GetBooksByAuthor(ctx context.Context, authorName string) ([]domain.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooksByAuthor", ctx, authorName)
	ret0, _ := ret[0].([]domain.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooksByAuthor indicates an expected call of GetBooksByAuthor.
func (mr *MockDBPortMockRecorder) GetBooksByAuthor(ctx, authorName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooksByAuthor", reflect.TypeOf((*MockDBPort)(nil).GetBooksByAuthor), ctx, authorName)
}

// GetBooksByGenre mocks base method.
func (m *MockDBPort) GetBooksByGenre(ctx context.Context, genreName string) ([]domain.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooksByGenre", ctx, genreName)
	ret0, _ := ret[0].([]domain.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooksByGenre indicates an expected call of GetBooksByGenre.
func (mr *MockDBPortMockRecorder) GetBooksByGenre(ctx, genreName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooksByGenre", reflect.TypeOf((*MockDBPort)(nil).GetBooksByGenre), ctx, genreName)
}

// GetBooksByTitle mocks base method.
func (m *MockDBPort) GetBooksByTitle(ctx context.Context, title string) ([]domain.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooksByTitle", ctx, title)
	ret0, _ := ret[0].([]domain.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooksByTitle indicates an expected call of GetBooksByTitle.
func (mr *MockDBPortMockRecorder) GetBooksByTitle(ctx, title any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooksByTitle", reflect.TypeOf((*MockDBPort)(nil).GetBooksByTitle), ctx, title)
}

// ModifyBookByID mocks base method.
func (m *MockDBPort) ModifyBookByID(ctx context.Context, ID domain.BookID, book *domain.Book) (*domain.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyBookByID", ctx, ID, book)
	ret0, _ := ret[0].(*domain.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyBookByID indicates an expected call of ModifyBookByID.
func (mr *MockDBPortMockRecorder) ModifyBookByID(ctx, ID, book any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyBookByID", reflect.TypeOf((*MockDBPort)(nil).ModifyBookByID), ctx, ID, book)
}
