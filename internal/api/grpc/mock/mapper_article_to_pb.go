// Code generated by MockGen. DO NOT EDIT.
// Source: mapper_article_to_pb.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	internal "github.com/angver/employcitytestcase/internal"
	articlev1 "github.com/angver/employcitytestcase/internal/api/grpc/gen/employcity/microservice/article/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockArticleToPbMapper is a mock of ArticleToPbMapper interface.
type MockArticleToPbMapper struct {
	ctrl     *gomock.Controller
	recorder *MockArticleToPbMapperMockRecorder
}

// MockArticleToPbMapperMockRecorder is the mock recorder for MockArticleToPbMapper.
type MockArticleToPbMapperMockRecorder struct {
	mock *MockArticleToPbMapper
}

// NewMockArticleToPbMapper creates a new mock instance.
func NewMockArticleToPbMapper(ctrl *gomock.Controller) *MockArticleToPbMapper {
	mock := &MockArticleToPbMapper{ctrl: ctrl}
	mock.recorder = &MockArticleToPbMapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleToPbMapper) EXPECT() *MockArticleToPbMapperMockRecorder {
	return m.recorder
}

// MapArticle mocks base method.
func (m *MockArticleToPbMapper) MapArticle(article *internal.Article) *articlev1.Article {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapArticle", article)
	ret0, _ := ret[0].(*articlev1.Article)
	return ret0
}

// MapArticle indicates an expected call of MapArticle.
func (mr *MockArticleToPbMapperMockRecorder) MapArticle(article interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapArticle", reflect.TypeOf((*MockArticleToPbMapper)(nil).MapArticle), article)
}
