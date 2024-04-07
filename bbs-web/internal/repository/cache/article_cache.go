package cache

import (
	"bbs-web/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-07 14:41

type ArticleCache interface {
	GetFirstPage(ctx context.Context, uid int64) ([]domain.Article, error)
	SetFirstPage(ctx context.Context, uid int64, data []domain.Article) error
	DelFirstPage(ctx context.Context, uid int64) error

	Set(ctx context.Context, data domain.Article) error
}

type articleCache struct {
	client redis.Cmdable
}

func NewArticleCache(cmd redis.Cmdable) ArticleCache {
	return &articleCache{client: cmd}
}

func (c *articleCache) GetFirstPage(ctx context.Context, uid int64) ([]domain.Article, error) {
	res, err := c.client.Get(ctx, c.firstPageKey(uid)).Bytes()
	if err != nil {
		return nil, err
	}
	var ans []domain.Article
	err = json.Unmarshal(res, &ans)
	return ans, err
}

func (c *articleCache) SetFirstPage(ctx context.Context, uid int64, data []domain.Article) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, c.firstPageKey(uid), res, time.Minute*10).Err() // 设置10分钟的过期时间
}

func (c *articleCache) DelFirstPage(ctx context.Context, uid int64) error {
	return c.client.Del(ctx, c.firstPageKey(uid)).Err()
}

func (c *articleCache) Set(ctx context.Context, data domain.Article) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, c.authorArtKey(data.Id), res, time.Second*30).Err() // 设置30秒过期
}
func (c *articleCache) firstPageKey(uid int64) string {
	return fmt.Sprintf("article:first_page:%d", uid)
}

func (c *articleCache) authorArtKey(uid int64) string {
	return fmt.Sprintf("article:author:%d", uid)
}
