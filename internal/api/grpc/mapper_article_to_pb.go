package grpc

//go:generate mockgen -source=mapper_article_to_pb.go -destination=./mock/mapper_article_to_pb.go -package=mock

import (
	"github.com/angver/employcitytestcase/internal"
	articlev1 "github.com/angver/employcitytestcase/internal/api/grpc/gen/employcity/microservice/article/v1"
)

type ArticleToPbMapper interface {
	MapArticle(article *internal.Article) *articlev1.Article
}

func NewArticleToPbMapper() *articleToPbMapper {
	return &articleToPbMapper{}
}

type articleToPbMapper struct{}

func (m *articleToPbMapper) MapArticle(article *internal.Article) *articlev1.Article {
	return &articlev1.Article{
		Id:      int32(article.Id),
		Title:   article.Title,
		Content: article.Content,
	}
}
