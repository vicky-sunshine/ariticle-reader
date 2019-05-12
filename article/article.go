package article

type Article interface {
	GetID() string
	GetTitle() string
	GetAuthor() string
	GetTimestamp() int64
}

type ArticleReader interface {
	TopArticles(number int) ([]Article, error)
	GetArticle(id string) (Article, error)
}
