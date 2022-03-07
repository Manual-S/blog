package mysql

import "gorm.io/gorm"

type ArticleClient interface {
	QueryArticleList() ([]Article, error)
	CreateArticle(article Article) error
	QueryArticleInfo(articleID string) (Article, error)
}

type Article struct {
	*gorm.Model
	ArticleId string `json:"article_id"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
	Status    int    `json:"status"`
	Votes     int    `json:"votes"`
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

func (a *articleClient) QueryArticleInfo(articleID string) (Article, error) {
	var aInfo Article
	err := a.db.Where("article_id = ?", articleID).Find(&aInfo).Error
	return aInfo, err
}
