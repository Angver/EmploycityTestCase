package grpc

import (
	"reflect"
	"testing"

	"github.com/angver/employcitytestcase/internal"
	articlev1 "github.com/angver/employcitytestcase/internal/api/grpc/gen/employcity/microservice/article/v1"
)

func TestNewArticleToPbMapper(t *testing.T) {
	tests := []struct {
		name string
		want *articleToPbMapper
	}{
		{
			name: "correct",
			want: &articleToPbMapper{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArticleToPbMapper(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArticleToPbMapper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_articleToPbMapper_MapArticle(t *testing.T) {
	type args struct {
		article *internal.Article
	}
	tests := []struct {
		name string
		args args
		want *articlev1.Article
	}{
		{
			name: "correct",
			args: args{article: &internal.Article{
				Id:      111,
				Title:   "title",
				Content: "content",
			}},
			want: &articlev1.Article{
				Id:      111,
				Title:   "title",
				Content: "content",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := articleToPbMapper{}
			if got := m.MapArticle(tt.args.article); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}
