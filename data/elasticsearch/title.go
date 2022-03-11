// 根据title进行文章标题的搜索
package elasticsearch

import (
	"context"

	"github.com/olivere/elastic"
)

type Article struct {
	ArticleId string `json:"article_id"`
}

type TitleClient interface {
	QueryTitle(ctx context.Context, title string) ([]Article, error)
}

type titleClient struct {
	esclient *elastic.Client
}

func NewTitleClient(client *elastic.Client) TitleClient {
	return &titleClient{
		esclient: client,
	}
}

func (t *titleClient) QueryTitle(ctx context.Context, title string) ([]Article, error) {
	return nil, nil
}
