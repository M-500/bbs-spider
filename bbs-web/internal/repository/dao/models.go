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
	Avatar   string     `gorm:"type:varchar(200);comment:用户头像"`
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

type InteractiveModel struct {
	gorm.Model
	// 同一个资源只能有一行，也就是要有联合唯一索引
	BizId      int64  `gorm:"comment:业务ID;not null;uniqueIndex:id_type_biz"` // 因为建立联合索引时候，索引的顺序只和结构体字段的顺序有关，所以要注意 bizID和biz的顺序不能乱
	Biz        string `gorm:"comment:业务标识符;type:varchar(128);not null;uniqueIndex:id_type_biz"`
	ReadCnt    int64  `gorm:"comment:阅读数;default:0"` // 阅读计数
	LikeCnt    int64  `gorm:"comment:点赞数;default:0"` // 点赞数
	CollectCnt int64  `gorm:"comment:收藏数;default:0"` // 收藏数
	CommentCnt int64  `gorm:"comment:评论数;default:0"` // 评论数
}

func (InteractiveModel) TableName() string {
	return "interactive"
}

// UserLikeBizModel 用户点赞资源关联表
type UserLikeBizModel struct {
	gorm.Model
	// 思考一下为什么联合索引的顺序是 bizid biz uid （提示: 根据业务场景来）
	BizId int64  `gorm:"comment:业务ID;not null;uniqueIndex:uid_type_biz"` // 因为建立联合索引时候，索引的顺序只和结构体字段的顺序有关，所以要注意 bizID和biz的顺序不能乱
	Biz   string `gorm:"comment:业务标识符;type:varchar(128);not null;uniqueIndex:uid_type_biz"`
	Uid   int64  `gorm:"comment:用户ID;not null;uniqueIndex:uid_type_biz"`
}

func (UserLikeBizModel) TableName() string {
	return "user_to_biz_like"
}

// Collection
// @Description: 收藏夹
type Collection struct {
}
