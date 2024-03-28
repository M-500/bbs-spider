package vo

import "bbs-web/internal/domain"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-27 11:55

type ArticleReq struct {
	Id          int64  `json:"id" binding:"-"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Summary     string `json:"summary" binding:"-"`
	ContentType string `json:"content_type" binding:"required"`
	Cover       string `json:"cover" binding:"-"`
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
