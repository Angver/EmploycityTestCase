package memcached

import (
	"encoding/json"
	"fmt"

	"github.com/angver/employcitytestcase/internal"
	"github.com/angver/employcitytestcase/internal/memcached/client"
	"github.com/bradfitz/gomemcache/memcache"
)

// NewArticleStorage создает новое хранилище в мемкеше
func NewArticleStorage(addr string) *ArticleStorage {
	return &ArticleStorage{mc: memcache.New(addr)}
}

// ArticleStorage хранилище в мемкеше
type ArticleStorage struct {
	mc client.MemcachedClient
}

func (s *ArticleStorage) Set(id internal.ArticleId, title string, content string) (*internal.Article, error) {
	article := &internal.Article{
		Id:      id,
		Title:   title,
		Content: content,
	}
	cached, err := json.Marshal(article)
	if err != nil {
		return nil, fmt.Errorf("can't marshal article: %w", err)
	}

	err = s.mc.Set(&memcache.Item{Key: string(id), Value: cached})
	if err != nil {
		return nil, fmt.Errorf("can't set article: %w", err)
	}

	return article, nil
}

func (s *ArticleStorage) Get(id internal.ArticleId) (*internal.Article, error) {
	item, err := s.mc.Get(string(id))
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return nil, nil
		}

		return nil, fmt.Errorf("can't get article: %w", err)
	}
	var article *internal.Article
	err = json.Unmarshal(item.Value, &article)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal article: %w", err)
	}

	return article, nil
}

func (s *ArticleStorage) Delete(id internal.ArticleId) error {
	err := s.mc.Delete(string(id))
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return nil
		}

		return fmt.Errorf("can't delete article: %w", err)
	}

	return nil
}
