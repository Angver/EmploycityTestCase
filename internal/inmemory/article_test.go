package inmemory

import (
	"reflect"
	"testing"

	"github.com/angver/employcitytestcase/internal"
	"github.com/stretchr/testify/assert"
)

func TestArticleStorage_Delete(t *testing.T) {
	type fields struct {
		articles map[internal.ArticleId]*internal.Article
		lastId   internal.ArticleId
	}
	type args struct {
		id internal.ArticleId
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "correct",
			fields: fields{
				articles: map[internal.ArticleId]*internal.Article{
					1: {
						Id:      1,
						Title:   "q",
						Content: "q",
					},
					2: {
						Id:      2,
						Title:   "w",
						Content: "w",
					},
				},
				lastId: 3,
			},
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ArticleStorage{
				articles: tt.fields.articles,
				lastId:   tt.fields.lastId,
			}
			if err := s.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArticleStorage_Get(t *testing.T) {
	type fields struct {
		articles map[internal.ArticleId]*internal.Article
		lastId   internal.ArticleId
	}
	type args struct {
		id internal.ArticleId
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *internal.Article
		wantErr bool
	}{
		{
			name: "correct",
			fields: fields{
				articles: map[internal.ArticleId]*internal.Article{
					1: {
						Id:      1,
						Title:   "q",
						Content: "q",
					},
					2: {
						Id:      2,
						Title:   "w",
						Content: "w",
					},
				},
				lastId: 3,
			},
			args: args{id: 2},
			want: &internal.Article{
				Id:      2,
				Title:   "w",
				Content: "w",
			},
			wantErr: false,
		},
		{
			name: "empty result",
			fields: fields{
				articles: map[internal.ArticleId]*internal.Article{
					1: {
						Id:      1,
						Title:   "q",
						Content: "q",
					},
					2: {
						Id:      2,
						Title:   "w",
						Content: "w",
					},
				},
				lastId: 3,
			},
			args:    args{id: 3},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ArticleStorage{
				articles: tt.fields.articles,
				lastId:   tt.fields.lastId,
			}
			got, err := s.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleStorage_Set(t *testing.T) {
	type fields struct {
		articles map[internal.ArticleId]*internal.Article
		lastId   internal.ArticleId
	}
	type args struct {
		id      internal.ArticleId
		title   string
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *internal.Article
		wantErr bool
	}{
		{
			name: "correct create",
			fields: fields{
				articles: map[internal.ArticleId]*internal.Article{
					1: {
						Id:      1,
						Title:   "q",
						Content: "q",
					},
					2: {
						Id:      2,
						Title:   "w",
						Content: "w",
					},
				},
				lastId: 3,
			},
			args: args{
				id:      0,
				title:   "e",
				content: "e",
			},
			want: &internal.Article{
				Id:      4,
				Title:   "e",
				Content: "e",
			},
			wantErr: false,
		},
		{
			name: "correct update",
			fields: fields{
				articles: map[internal.ArticleId]*internal.Article{
					1: {
						Id:      1,
						Title:   "q",
						Content: "q",
					},
					2: {
						Id:      2,
						Title:   "w",
						Content: "w",
					},
				},
				lastId: 3,
			},
			args: args{
				id:      2,
				title:   "e",
				content: "e",
			},
			want: &internal.Article{
				Id:      2,
				Title:   "e",
				Content: "e",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ArticleStorage{
				articles: tt.fields.articles,
				lastId:   tt.fields.lastId,
			}
			got, err := s.Set(tt.args.id, tt.args.title, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleStorage_update(t *testing.T) {
	type fields struct {
		articles map[internal.ArticleId]*internal.Article
		lastId   internal.ArticleId
	}
	type args struct {
		id      internal.ArticleId
		title   string
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *internal.Article
		wantErr bool
	}{
		{
			name: "correct",
			fields: fields{
				articles: map[internal.ArticleId]*internal.Article{
					1: {
						Id:      1,
						Title:   "q",
						Content: "q",
					},
					2: {
						Id:      2,
						Title:   "w",
						Content: "w",
					},
				},
				lastId: 3,
			},
			args: args{
				id:      2,
				title:   "e",
				content: "e",
			},
			want: &internal.Article{
				Id:      2,
				Title:   "e",
				Content: "e",
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				articles: map[internal.ArticleId]*internal.Article{
					1: {
						Id:      1,
						Title:   "q",
						Content: "q",
					},
					2: {
						Id:      2,
						Title:   "w",
						Content: "w",
					},
				},
				lastId: 3,
			},
			args: args{
				id:      3,
				title:   "e",
				content: "e",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ArticleStorage{
				articles: tt.fields.articles,
				lastId:   tt.fields.lastId,
			}
			got, err := s.update(tt.args.id, tt.args.title, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewArticleStorage(t *testing.T) {
	storage := NewArticleStorage()
	assert.Len(t, storage.articles, 0)
	assert.Zero(t, storage.lastId)
}
