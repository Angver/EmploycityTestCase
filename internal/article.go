package internal

//go:generate mockgen -source=article.go -destination=./article_mock.go -package=internal

type ArticleId int32

type Article struct {
	Id      ArticleId
	Title   string
	Content string
}

type ArticleStorage interface {
	// Set создаёт или обновляет данные
	Set(id ArticleId, title string, content string) (*Article, error)
	// Get получает данные по идентификатору
	Get(id ArticleId) (*Article, error)
	// Delete удаляет данные
	Delete(id ArticleId) error
}
