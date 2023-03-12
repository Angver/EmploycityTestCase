package memcached

import (
	"errors"
	"reflect"
	"testing"

	"github.com/angver/employcitytestcase/internal"
	"github.com/angver/employcitytestcase/internal/memcached/client"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestArticleStorage_Delete(t *testing.T) {
	type mocks struct {
		mc *client.MockMemcachedClient
	}
	type args struct {
		id internal.ArticleId
	}
	tests := []struct {
		name    string
		args    *args
		prepare func(a *args, m *mocks) error
	}{
		{
			name: "correct",
			args: &args{id: 1},
			prepare: func(a *args, m *mocks) error {
				m.mc.EXPECT().Delete(string(a.id)).Return(nil)

				return nil
			},
		},
		{
			name: "not found",
			args: &args{id: 1},
			prepare: func(a *args, m *mocks) error {
				m.mc.EXPECT().Delete(string(a.id)).Return(memcache.ErrCacheMiss)

				return nil
			},
		},
		{
			name: "error",
			args: &args{id: 1},
			prepare: func(a *args, m *mocks) error {
				m.mc.EXPECT().Delete(string(a.id)).Return(errors.New("some error"))

				return errors.New("can't delete article: some error")
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mocks := &mocks{
				mc: client.NewMockMemcachedClient(ctrl),
			}
			wantErr := tt.prepare(tt.args, mocks)
			s := &ArticleStorage{
				mc: mocks.mc,
			}
			err := s.Delete(tt.args.id)
			if wantErr != nil {
				assert.EqualError(t, err, wantErr.Error())
			}
		})
	}
}

func TestArticleStorage_Get(t *testing.T) {
	type mocks struct {
		mc *client.MockMemcachedClient
	}
	type args struct {
		id internal.ArticleId
	}
	tests := []struct {
		name    string
		args    *args
		prepare func(a *args, m *mocks) (*internal.Article, error)
	}{
		{
			name: "correct",
			args: &args{id: 1},
			prepare: func(a *args, m *mocks) (*internal.Article, error) {
				m.mc.EXPECT().Get(string(a.id)).Return(&memcache.Item{Value: []byte(`{"Id":1,"Title":"q","Content":"q"}`)}, nil)

				return &internal.Article{
					Id:      1,
					Title:   "q",
					Content: "q",
				}, nil
			},
		},
		{
			name: "not found",
			args: &args{id: 1},
			prepare: func(a *args, m *mocks) (*internal.Article, error) {
				m.mc.EXPECT().Get(string(a.id)).Return(nil, memcache.ErrCacheMiss)

				return nil, nil
			},
		},
		{
			name: "error",
			args: &args{id: 1},
			prepare: func(a *args, m *mocks) (*internal.Article, error) {
				m.mc.EXPECT().Get(string(a.id)).Return(nil, errors.New("some error"))

				return nil, errors.New("can't get article: some error")
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mocks := &mocks{
				mc: client.NewMockMemcachedClient(ctrl),
			}
			s := &ArticleStorage{
				mc: mocks.mc,
			}
			want, wantErr := tt.prepare(tt.args, mocks)
			got, err := s.Get(tt.args.id)
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

func TestArticleStorage_Set(t *testing.T) {
	type mocks struct {
		mc *client.MockMemcachedClient
	}
	type args struct {
		id      internal.ArticleId
		title   string
		content string
	}
	tests := []struct {
		name    string
		args    *args
		prepare func(a *args, m *mocks) (*internal.Article, error)
	}{
		{
			name: "correct",
			args: &args{
				id:      1,
				title:   "q",
				content: "q",
			},
			prepare: func(a *args, m *mocks) (*internal.Article, error) {
				m.mc.EXPECT().Set(&memcache.Item{
					Key:   string(a.id),
					Value: []byte(`{"Id":1,"Title":"q","Content":"q"}`),
				}).Return(nil)

				return &internal.Article{
					Id:      1,
					Title:   "q",
					Content: "q",
				}, nil
			},
		},
		{
			name: "error",
			args: &args{
				id:      1,
				title:   "q",
				content: "q",
			},
			prepare: func(a *args, m *mocks) (*internal.Article, error) {
				m.mc.EXPECT().Set(&memcache.Item{
					Key:   string(a.id),
					Value: []byte(`{"Id":1,"Title":"q","Content":"q"}`),
				}).Return(errors.New("some error"))

				return nil, errors.New("can't set article: some error")
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mocks := &mocks{
				mc: client.NewMockMemcachedClient(ctrl),
			}
			want, wantErr := tt.prepare(tt.args, mocks)
			s := &ArticleStorage{
				mc: mocks.mc,
			}
			got, err := s.Set(tt.args.id, tt.args.title, tt.args.content)
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

func TestNewArticleStorage(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
		want *ArticleStorage
	}{
		{
			name: "correct",
			args: args{addr: "127.0.0.1:11211"},
			want: &ArticleStorage{mc: memcache.New("127.0.0.1:11211")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArticleStorage(tt.args.addr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArticleStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}
