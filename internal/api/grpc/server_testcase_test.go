package grpc

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/angver/employcitytestcase/internal"
	articlev1 "github.com/angver/employcitytestcase/internal/api/grpc/gen/employcity/microservice/article/v1"
	"github.com/angver/employcitytestcase/internal/api/grpc/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestNewServerTestCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	mocks := struct {
		storage  *internal.MockArticleStorage
		pbMapper *mock.MockArticleToPbMapper
	}{
		storage:  internal.NewMockArticleStorage(ctrl),
		pbMapper: mock.NewMockArticleToPbMapper(ctrl),
	}
	tests := []struct {
		name string
		want *ServerTestCase
	}{
		{
			name: "correct",
			want: &ServerTestCase{
				storage:  mocks.storage,
				pbMapper: mocks.pbMapper,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServerTestCase(mocks.storage, mocks.pbMapper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServerTestCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServerTestCase_Create(t *testing.T) {
	type mocks struct {
		storage  *internal.MockArticleStorage
		pbMapper *mock.MockArticleToPbMapper
	}
	type args struct {
		in0 context.Context
		r   *articlev1.CreateRequest
	}
	tests := []struct {
		name    string
		args    *args
		prepare func(a *args, m *mocks) (*articlev1.CreateResponse, error)
	}{
		{
			name: "correct",
			args: &args{
				in0: context.Background(),
				r: &articlev1.CreateRequest{
					Title:   "title",
					Content: "content",
				},
			},
			prepare: func(a *args, m *mocks) (*articlev1.CreateResponse, error) {
				article := &internal.Article{
					Id:      1,
					Title:   "title",
					Content: "content",
				}
				pbArticle := &articlev1.Article{
					Id:      1,
					Title:   "title",
					Content: "content",
				}
				m.storage.EXPECT().Set(internal.ArticleId(0), a.r.GetTitle(), a.r.GetContent()).Return(article, nil)
				m.pbMapper.EXPECT().MapArticle(article).Return(pbArticle)

				return &articlev1.CreateResponse{Article: pbArticle}, nil
			},
		},
		{
			name: "error",
			args: &args{
				in0: context.Background(),
				r: &articlev1.CreateRequest{
					Title:   "title",
					Content: "content",
				},
			},
			prepare: func(a *args, m *mocks) (*articlev1.CreateResponse, error) {
				m.storage.EXPECT().Set(internal.ArticleId(0), a.r.GetTitle(), a.r.GetContent()).Return(nil, errors.New("some error"))

				return nil, errors.New("rpc error: code = Internal desc = can't create article: some error")
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mocks := &mocks{
				storage:  internal.NewMockArticleStorage(ctrl),
				pbMapper: mock.NewMockArticleToPbMapper(ctrl),
			}
			want, wantErr := tt.prepare(tt.args, mocks)
			c := ServerTestCase{
				storage:  mocks.storage,
				pbMapper: mocks.pbMapper,
			}
			got, err := c.Create(tt.args.in0, tt.args.r)
			if wantErr == nil {
				assert.Equal(t, want, got)
				assert.Nil(t, err)
			} else {
				assert.Nil(t, got)
				assert.EqualError(t, err, wantErr.Error())
			}
		})
	}
}

func TestServerTestCase_Delete(t *testing.T) {
	type mocks struct {
		storage *internal.MockArticleStorage
	}
	type args struct {
		in0 context.Context
		r   *articlev1.DeleteRequest
	}
	tests := []struct {
		name    string
		args    *args
		prepare func(a *args, m *mocks) (*articlev1.DeleteResponse, error)
	}{
		{
			name: "correct",
			args: &args{
				in0: context.Background(),
				r:   &articlev1.DeleteRequest{Ids: []int32{1}},
			},
			prepare: func(a *args, m *mocks) (*articlev1.DeleteResponse, error) {
				m.storage.EXPECT().Delete(internal.ArticleId(1)).Return(nil)

				return &articlev1.DeleteResponse{}, nil
			},
		},
		{
			name: "error",
			args: &args{
				in0: context.Background(),
				r:   &articlev1.DeleteRequest{Ids: []int32{1}},
			},
			prepare: func(a *args, m *mocks) (*articlev1.DeleteResponse, error) {
				m.storage.EXPECT().Delete(internal.ArticleId(1)).Return(errors.New("some error"))

				return nil, errors.New("rpc error: code = Internal desc = can't delete article: some error")
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mocks := &mocks{storage: internal.NewMockArticleStorage(ctrl)}
			want, wantErr := tt.prepare(tt.args, mocks)
			c := ServerTestCase{
				storage: mocks.storage,
			}
			got, err := c.Delete(tt.args.in0, tt.args.r)
			if wantErr == nil {
				assert.Equal(t, want, got)
				assert.Nil(t, err)
			} else {
				assert.Nil(t, got)
				assert.EqualError(t, err, wantErr.Error())
			}
		})
	}
}

func TestServerTestCase_Get(t *testing.T) {
	type mocks struct {
		storage  *internal.MockArticleStorage
		pbMapper *mock.MockArticleToPbMapper
	}
	type args struct {
		in0 context.Context
		r   *articlev1.GetRequest
	}
	tests := []struct {
		name    string
		args    *args
		prepare func(a *args, m *mocks) (*articlev1.GetResponse, error)
	}{
		{
			name: "correct",
			args: &args{
				in0: context.Background(),
				r:   &articlev1.GetRequest{Id: 1},
			},
			prepare: func(a *args, m *mocks) (*articlev1.GetResponse, error) {
				article := &internal.Article{
					Id:      1,
					Title:   "title",
					Content: "content",
				}
				pbArticle := &articlev1.Article{
					Id:      1,
					Title:   "title",
					Content: "content",
				}
				m.storage.EXPECT().Get(internal.ArticleId(a.r.GetId())).Return(article, nil)
				m.pbMapper.EXPECT().MapArticle(article).Return(pbArticle)

				return &articlev1.GetResponse{Article: pbArticle}, nil
			},
		},
		{
			name: "error",
			args: &args{
				in0: context.Background(),
				r:   &articlev1.GetRequest{Id: 1},
			},
			prepare: func(a *args, m *mocks) (*articlev1.GetResponse, error) {
				m.storage.EXPECT().Get(internal.ArticleId(a.r.GetId())).Return(nil, errors.New("some error"))

				return nil, errors.New("rpc error: code = Internal desc = can't get article: some error")
			},
		},
		{
			name: "not found",
			args: &args{
				in0: context.Background(),
				r:   &articlev1.GetRequest{Id: 1},
			},
			prepare: func(a *args, m *mocks) (*articlev1.GetResponse, error) {
				m.storage.EXPECT().Get(internal.ArticleId(a.r.GetId())).Return(nil, nil)

				return nil, errors.New("rpc error: code = NotFound desc = article doesn't exist")
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mocks := &mocks{
				storage:  internal.NewMockArticleStorage(ctrl),
				pbMapper: mock.NewMockArticleToPbMapper(ctrl),
			}
			want, wantErr := tt.prepare(tt.args, mocks)
			c := ServerTestCase{
				storage:  mocks.storage,
				pbMapper: mocks.pbMapper,
			}
			got, err := c.Get(tt.args.in0, tt.args.r)
			if wantErr == nil {
				assert.Equal(t, want, got)
				assert.Nil(t, err)
			} else {
				assert.Nil(t, got)
				assert.EqualError(t, err, wantErr.Error())
			}
		})
	}
}

func TestServerTestCase_Update(t *testing.T) {
	type mocks struct {
		storage  *internal.MockArticleStorage
		pbMapper *mock.MockArticleToPbMapper
	}
	type args struct {
		in0 context.Context
		r   *articlev1.UpdateRequest
	}
	tests := []struct {
		name    string
		args    *args
		prepare func(a *args, m *mocks) (*articlev1.UpdateResponse, error)
	}{
		{
			name: "correct",
			args: &args{
				in0: context.Background(),
				r: &articlev1.UpdateRequest{
					Id: 1,
					Fields: &articlev1.UpdateRequest_Fields{
						Title:   &wrapperspb.StringValue{Value: "new title"},
						Content: &wrapperspb.StringValue{Value: "new content"},
					},
				},
			},
			prepare: func(a *args, m *mocks) (*articlev1.UpdateResponse, error) {
				m.storage.EXPECT().Get(internal.ArticleId(1)).Return(&internal.Article{
					Id:      1,
					Title:   "title",
					Content: "content",
				}, nil)
				m.storage.EXPECT().Set(
					internal.ArticleId(1),
					a.r.GetFields().GetTitle().GetValue(),
					a.r.GetFields().GetContent().GetValue(),
				).Return(&internal.Article{
					Id:      1,
					Title:   "new title",
					Content: "new content",
				}, nil)
				pbArticle := &articlev1.Article{
					Id:      1,
					Title:   "new title",
					Content: "new content",
				}
				m.pbMapper.EXPECT().MapArticle(&internal.Article{
					Id:      1,
					Title:   "new title",
					Content: "new content",
				}).Return(pbArticle)

				return &articlev1.UpdateResponse{Article: pbArticle}, nil
			},
		},
		{
			name: "invalid id",
			args: &args{
				in0: context.Background(),
				r: &articlev1.UpdateRequest{
					Id: 0,
					Fields: &articlev1.UpdateRequest_Fields{
						Title:   &wrapperspb.StringValue{Value: "new title"},
						Content: &wrapperspb.StringValue{Value: "new content"},
					},
				},
			},
			prepare: func(a *args, m *mocks) (*articlev1.UpdateResponse, error) {
				return nil, errors.New("rpc error: code = InvalidArgument desc = id is required")
			},
		},
		{
			name: "error getting article",
			args: &args{
				in0: context.Background(),
				r: &articlev1.UpdateRequest{
					Id: 1,
					Fields: &articlev1.UpdateRequest_Fields{
						Title:   &wrapperspb.StringValue{Value: "new title"},
						Content: &wrapperspb.StringValue{Value: "new content"},
					},
				},
			},
			prepare: func(a *args, m *mocks) (*articlev1.UpdateResponse, error) {
				m.storage.EXPECT().Get(internal.ArticleId(1)).Return(nil, errors.New("some error"))

				return nil, errors.New("rpc error: code = InvalidArgument desc = can't get article: some error")
			},
		},
		{
			name: "article not found",
			args: &args{
				in0: context.Background(),
				r: &articlev1.UpdateRequest{
					Id: 1,
					Fields: &articlev1.UpdateRequest_Fields{
						Title:   &wrapperspb.StringValue{Value: "new title"},
						Content: &wrapperspb.StringValue{Value: "new content"},
					},
				},
			},
			prepare: func(a *args, m *mocks) (*articlev1.UpdateResponse, error) {
				m.storage.EXPECT().Get(internal.ArticleId(1)).Return(nil, nil)

				return nil, errors.New("rpc error: code = NotFound desc = article doesn't exist")
			},
		},
		{
			name: "nothing to update",
			args: &args{
				in0: context.Background(),
				r: &articlev1.UpdateRequest{
					Id: 1,
				},
			},
			prepare: func(a *args, m *mocks) (*articlev1.UpdateResponse, error) {
				article := &internal.Article{
					Id:      1,
					Title:   "title",
					Content: "content",
				}
				pbArticle := &articlev1.Article{
					Id:      1,
					Title:   "title",
					Content: "content",
				}
				m.storage.EXPECT().Get(internal.ArticleId(1)).Return(article, nil)
				m.pbMapper.EXPECT().MapArticle(article).Return(pbArticle)

				return &articlev1.UpdateResponse{Article: pbArticle}, nil
			},
		},
		{
			name: "setting error",
			args: &args{
				in0: context.Background(),
				r: &articlev1.UpdateRequest{
					Id: 1,
					Fields: &articlev1.UpdateRequest_Fields{
						Title:   &wrapperspb.StringValue{Value: "new title"},
						Content: &wrapperspb.StringValue{Value: "new content"},
					},
				},
			},
			prepare: func(a *args, m *mocks) (*articlev1.UpdateResponse, error) {
				m.storage.EXPECT().Get(internal.ArticleId(1)).Return(&internal.Article{
					Id:      1,
					Title:   "title",
					Content: "content",
				}, nil)
				m.storage.EXPECT().Set(
					internal.ArticleId(1),
					a.r.GetFields().GetTitle().GetValue(),
					a.r.GetFields().GetContent().GetValue(),
				).Return(nil, errors.New("some error"))

				return nil, errors.New("rpc error: code = Internal desc = can't update article: some error")
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mocks := &mocks{
				storage:  internal.NewMockArticleStorage(ctrl),
				pbMapper: mock.NewMockArticleToPbMapper(ctrl),
			}
			want, wantErr := tt.prepare(tt.args, mocks)
			c := ServerTestCase{
				storage:  mocks.storage,
				pbMapper: mocks.pbMapper,
			}
			got, err := c.Update(tt.args.in0, tt.args.r)
			if wantErr == nil {
				assert.Equal(t, want, got)
				assert.Nil(t, err)
			} else {
				assert.Nil(t, got)
				assert.EqualError(t, err, wantErr.Error())
			}
		})
	}
}
