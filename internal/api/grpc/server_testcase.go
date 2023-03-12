package grpc

import (
	"context"
	"fmt"

	"github.com/angver/employcitytestcase/internal"
	articlev1 "github.com/angver/employcitytestcase/internal/api/grpc/gen/employcity/microservice/article/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewServerTestCase(
	storage internal.ArticleStorage,
	pbMapper ArticleToPbMapper,
) *ServerTestCase {
	return &ServerTestCase{
		storage:  storage,
		pbMapper: pbMapper,
	}
}

type ServerTestCase struct {
	articlev1.UnimplementedArticleAPIServer
	storage  internal.ArticleStorage
	pbMapper ArticleToPbMapper
}

func (c ServerTestCase) Get(_ context.Context, r *articlev1.GetRequest) (*articlev1.GetResponse, error) {
	article, err := c.storage.Get(internal.ArticleId(r.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't get article: %s", err))
	}

	return &articlev1.GetResponse{Article: c.pbMapper.MapArticle(article)}, nil
}

func (c ServerTestCase) Create(_ context.Context, r *articlev1.CreateRequest) (*articlev1.CreateResponse, error) {
	article, err := c.storage.Set(0, r.GetTitle(), r.GetContent())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't create article: %s", err))
	}

	return &articlev1.CreateResponse{Article: c.pbMapper.MapArticle(article)}, nil
}

func (c ServerTestCase) Update(_ context.Context, r *articlev1.UpdateRequest) (*articlev1.UpdateResponse, error) {
	id := r.GetId()
	if id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	article, err := c.storage.Get(internal.ArticleId(id))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("can't get article: %s", err))
	}
	if article == nil {
		return nil, status.Error(codes.NotFound, "article doesn't exist")
	}

	if r.GetFields() == nil {
		return &articlev1.UpdateResponse{Article: c.pbMapper.MapArticle(article)}, nil
	}

	title := article.Title
	if r.GetFields().GetTitle() != nil {
		title = r.GetFields().GetTitle().GetValue()
	}

	content := article.Content
	if r.GetFields().GetContent() != nil {
		content = r.GetFields().GetContent().GetValue()
	}

	article, err = c.storage.Set(article.Id, title, content)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't update article: %s", err))
	}

	return &articlev1.UpdateResponse{Article: c.pbMapper.MapArticle(article)}, nil
}

func (c ServerTestCase) Delete(_ context.Context, r *articlev1.DeleteRequest) (*articlev1.DeleteResponse, error) {
	for _, id := range r.GetIds() {
		err := c.storage.Delete(internal.ArticleId(id))
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("can't delete article: %s", err))
		}
	}

	return &articlev1.DeleteResponse{}, nil
}
