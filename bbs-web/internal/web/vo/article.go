package vo

import "bbs-web/internal/domain"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-27 11:55

type ArticleReq struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Summary     string `json:"summary"`
	ContentType string `json:"content_type"`
	Cover       string `json:"cover"`
}

func (req ArticleReq) ToDomain(uid int64) domain.Article {
	return domain.Article{
		Id:          req.Id,
		Title:       req.Title,
		Content:     req.Content,
		Summary:     req.Summary,
		ContentType: req.ContentType,
		Cover:       req.Cover,
		Author: domain.Author{
			Id: uid,
		},
	}
}
