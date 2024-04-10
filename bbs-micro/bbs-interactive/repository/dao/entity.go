package dao

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 12:51

import "gorm.io/gorm"

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

type Collection struct {
	gorm.Model
	Name string `gorm:"comment:收藏夹名称;type:varchar(128)"`
	Uid  int64  `gorm:"comment:用户ID;not null;"`
}

func (Collection) TableName() string {
	return "collect"
}
