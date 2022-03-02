package mysql

import "gorm.io/gorm"

type ArticleClient interface {
	QueryArticleList() error
}

type Article struct {
}

func (a *Article) TableName() string {
	return "blog_article"
}

type articleClient struct {
	db *gorm.DB
}

func NewArticleClient(db *gorm.DB) ArticleClient {
	return &articleClient{
		db: db,
	}
}

func (a *articleClient) QueryArticleList() error {
	return nil
}
