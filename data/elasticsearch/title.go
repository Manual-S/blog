// 根据title进行文章标题的搜索
package elasticsearch

import (
	"context"

	"github.com/olivere/elastic"
)

type TitleClient interface {
	QueryTitle(ctx context.Context, title string) error
}

type titleClient struct {
	esclient *elastic.Client
}

func NewTitleClient(client *elastic.Client) TitleClient {
	return &titleClient{
		esclient: client,
	}
}

func (t *titleClient) QueryTitle(ctx context.Context, title string) error {
	return nil
}
