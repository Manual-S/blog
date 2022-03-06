package mysql

import "gorm.io/gorm"

type ArticleClient interface {
	QueryArticleList() ([]Article, error)
	CreateArticle(article Article) error
}

type Article struct {
	*gorm.Model
	ArticleId string
	Title     string
	Summary   string
	Content   string
	Status    int
	Votes     int
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

func (a *articleClient) QueryArticleList() ([]Article, error) {
	var list []Article
	err := a.db.Select("article_id,title,summary,votes").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (a *articleClient) CreateArticle(article Article) error {
	err := a.db.Create(&article).Error
	return err
}
