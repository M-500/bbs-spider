package dao

import (
	"gorm.io/gorm"
	"time"
)

type UserMode struct {
	gorm.Model
	Username    string     `gorm:"uniqueIndex;not null;type:varchar(32);column:username;comment:用户名，唯一索引，不能为空"`
	Phone       string     `gorm:"index;type:varchar(32);comment:手机号;column:phone"`
	Email       string     `gorm:"index;type:varchar(128);comment:邮箱;column:email"`
	HomePage    string     `gorm:"type:varchar(500);comment:个人主页;column:home_page"`
	Description string     `gorm:"type:text;comment:个性签名(描述);column:description"`
	Password    string     `gorm:"type:varchar(128);comment:密码;column:password"`
	Nickname    string     `gorm:"type:varchar(128);comment:昵称;column:nickname"`
	Gender      string     `gorm:"type:varchar(128);comment:性别;column:gender"`
	Birthday    *time.Time `gorm:"comment:生日;column:birthday"`
	IsAdmin     int64      `gorm:"int;default:0;comment:是否是admin，默认0,非管理员;column:gender"`
	Avatar      string     `gorm:"type:varchar(200);comment:用户头像"`
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

// CollectionModle
// @Description: 收藏夹
type CollectionModle struct {
	gorm.Model
	// 用户id和收藏夹的联合索引
	UserId      int64  `gorm:"用户ID;not null;column:user_id;uniqueIndex:uni_uid_cname"`
	CName       string `gorm:"comment:收藏夹名称;type:varchar(32);not null;column:c_name;uniqueIndex:uni_uid_cname"`
	Description string `gorm:"comment:收藏夹描述(可选);type:varchar(1024);column:description"`
	Sort        int64  `gorm:"comment:收藏夹排序;default:0;type:int;column:sort"`
	ResourceNum int64  `gorm:"comment:收藏夹内资源的数量;not null;default:0;column:resource_num"`
	IsPub       bool   `gorm:"comment:是否公开;column:is_pub"`
	CommentNum  int64  `gorm:"comment:收藏夹评论数量;not null;default:0;column:comment_num"`
}

func (CollectionModle) TableName() string {
	return "collect"
}

// UserCollectBizModel
// @Description: 业务收藏夹关联表
type UserCollectBizModel struct {
	gorm.Model
	BizId int64  `gorm:"comment:业务ID;not null;uniqueIndex:uid_type_biz"` // 因为建立联合索引时候，索引的顺序只和结构体字段的顺序有关，所以要注意 bizID和biz的顺序不能乱
	Biz   string `gorm:"comment:业务标识符;type:varchar(128);not null;uniqueIndex:uid_type_biz"`
	Uid   int64  `gorm:"comment:用户ID;not null;uniqueIndex:uid_type_biz"`
	Cid   int64  `gorm:"comment:收藏夹的ID;index"`
}

func (UserCollectBizModel) TableName() string {
	return "user_to_biz_collect"
}

type JobModel struct {
	gorm.Model

	Cfg    string `gorm:"comment:;"`
	Status int    // 用来表示状态

	NextExecTime int64 // 定时任务的下一次执行的时间

	Version int //MySQL乐观锁 实现并发安全

}

const (
	jobStatusWaiting = iota
	jobStatusRunning // 已经被抢占了
	jobStatusPaused  // 暂停了的 不会被调度的

)
