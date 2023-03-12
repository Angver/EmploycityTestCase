package inmemory

import (
	"errors"
	"sync"

	"github.com/angver/employcitytestcase/internal"
)

type ArticleStorage struct {
	mu       sync.RWMutex
	articles map[internal.ArticleId]*internal.Article
	lastId   internal.ArticleId
}

func NewArticleStorage() *ArticleStorage {
	return &ArticleStorage{
		articles: make(map[internal.ArticleId]*internal.Article),
	}
}

func (s ArticleStorage) Set(id internal.ArticleId, title string, content string) (*internal.Article, error) {
	if id == 0 {
		return s.create(title, content)
	}

	return s.update(id, title, content)
}

func (s ArticleStorage) Get(id internal.ArticleId) (*internal.Article, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.articles[id], nil
}

func (s ArticleStorage) Delete(id internal.ArticleId) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.articles, id)

	return nil
}

func (s ArticleStorage) create(title string, content string) (*internal.Article, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	article := &internal.Article{
		Id:      s.lastId + 1,
		Title:   title,
		Content: content,
	}

	s.articles[article.Id] = article
	s.lastId = article.Id

	return article, nil
}

func (s ArticleStorage) update(id internal.ArticleId, title string, content string) (*internal.Article, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	article, ok := s.articles[id]
	if !ok {
		return nil, errors.New("article doesn't exist")
	}

	article.Title = title
	article.Content = content

	return article, nil
}
