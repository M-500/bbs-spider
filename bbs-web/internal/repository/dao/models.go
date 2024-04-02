package dao

import (
	"gorm.io/gorm"
	"time"
)

type UserMode struct {
	gorm.Model
	Username string     `gorm:"uniqueIndex;not null;type:varchar(32);column:username;comment:用户名，唯一索引，不能为空"`
	Phone    string     `gorm:"index;type:varchar(32);comment:手机号;column:phone"`
	Email    string     `gorm:"index;type:varchar(128);comment:邮箱;column:email"`
	Password string     `gorm:"type:varchar(128);comment:密码;column:password"`
	Nickname string     `gorm:"type:varchar(128);comment:昵称;column:nickname"`
	Gender   string     `gorm:"type:varchar(128);comment:性别;column:gender"`
	Birthday *time.Time `gorm:"comment:生日;column:birthday"`
	IsAdmin  int64      `gorm:"int;default:0;comment:是否是admin，默认0,非管理员;column:gender"`
}

func (UserMode) TableName() string {
	return "users"
}

// 思考，为什么没有把
// 制作库，如何设计索引？从哪些角度考虑设计索引？
// 1. Where 查询条件 以及查询频率
type ArticleModel struct {
	gorm.Model
	AuthorId    int64  `gorm:"index:idx_article_user_id;comment:作者ID" `                                  // 所属用户编号
	Title       string `gorm:"size:128;not null;comment:标题" `                                            // 标题
	Summary     string `gorm:"type:text;comment:摘要" `                                                    // 摘要
	Content     string `gorm:"type:longtext;not null;comment:内容" `                                       // 内容
	ContentType string `gorm:"type:varchar(32);not null;comment:内容类型" `                                  // 内容类型：markdown、html
	Cover       string `gorm:"type:text;comment:封面图" `                                                   // 封面图
	Status      int    `gorm:"type:int(11);index:idx_article_status;comment:状态 0 草稿，1 待审核，2 审核通过，3 发布" ` // 状态
}

func (ArticleModel) TableName() string {
	return "articles_edit"
}

type PublishArticleModels struct {
	ArticleModel
}

func (PublishArticleModels) TableName() string {
	return "articles_pub"
}
