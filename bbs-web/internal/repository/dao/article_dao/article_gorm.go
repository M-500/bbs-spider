package article_dao

import (
	"bbs-web/internal/repository/dao"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 16:06

type ArticleDAO interface {
	Insert(ctx context.Context, art dao.ArticleModel) (int64, error)
	UpdateById(ctx context.Context, art dao.ArticleModel) error
	GetByAuthor(ctx context.Context, author int64, offset, limit int) ([]dao.ArticleModel, error)
	GetById(ctx context.Context, id int64) (dao.ArticleModel, error)
	GetPubById(ctx context.Context, id int64) (dao.PublishArticleModels, error)
	Transaction(ctx context.Context, bizFunc func(txDao ArticleDAO) error) error
	Sync(ctx context.Context, art dao.ArticleModel) (int64, error)
	Upsert(ctx context.Context, art dao.PublishArticleModels) error
	SyncStatus(ctx context.Context, author, id int64, status uint8) error
	ListPub(ctx context.Context, start time.Time, offset int, limit int) ([]dao.ArticleModel, error)
}

type gormArticleDao struct {
	db *gorm.DB
}

func NewGormArticleDao(db *gorm.DB) ArticleDAO {
	return &gormArticleDao{
		db: db,
	}
}

func (a *gormArticleDao) Insert(ctx context.Context, art dao.ArticleModel) (int64, error) {
	err := a.db.WithContext(ctx).Create(&art).Error
	return int64(art.ID), err
}

func (a *gormArticleDao) UpdateById(ctx context.Context, art dao.ArticleModel) error {
	now := time.Now()
	res := a.db.WithContext(ctx).Model(&dao.ArticleModel{}).
		Where("id = ? AND author_id = ?", art.ID, art.AuthorId).
		Updates(map[string]any{
			"title": art.Title,
			//"author_id":    art.AuthorId, // 创作者不能修改
			"summary":      art.Summary,
			"content":      art.Content,
			"content_type": art.ContentType,
			"cover":        art.Cover,
			"status":       art.Status,
			"updated_at":   now,
		})

	err := res.Error
	if err != nil {
		return err
	}
	if res.RowsAffected == 0 {
		return errors.New("更新数据失败")
	}
	return err
}

func (a *gormArticleDao) GetByAuthor(ctx context.Context, author int64, offset, limit int) ([]dao.ArticleModel, error) {
	var arts []dao.ArticleModel
	// 涉及order by的时候 一定要让order by 条件命中索引，因为索引天然有序
	// SQL优化案例 早起的order by 没有命中索引，所以内存排序很慢，优化这个查询
	err := a.db.WithContext(ctx).Model(&dao.ArticleModel{}).
		Where("author_id = ?", author).
		Offset(offset).
		Limit(limit).
		// 方式一 :
		Order("updated_at DESC, created_at ASC").
		// 方式二:
		//Order(clause.OrderBy{
		//	Columns: []clause.OrderByColumn{
		//		{Column: clause.Column{Name: "update_at"}, Desc: true},
		//		{Column: clause.Column{Name: "create_at"}, Desc: false},
		//	},
		//	Expression: nil,
		//}).
		Find(&arts).Error
	return arts, err
}

// GetByAuthorV1
//
//	@Description: 需要文章表和用户表进行关联，获取所有的文章和用户的信息
func (a *gormArticleDao) GetByAuthorV1(ctx context.Context, author int64, offset, limit int) ([]dao.ArticleModel, error) {
	var arts []dao.ArticleModel
	// 涉及order by的时候 一定要让order by 条件命中索引，因为索引天然有序
	// SQL优化案例 早起的order by 没有命中索引，所以内存排序很慢，优化这个查询
	err := a.db.WithContext(ctx).Model(&dao.ArticleModel{}).
		Where("author_id = ?", author).
		Offset(offset).
		Limit(limit).
		// 方式一 :
		Order("updated_at DESC, created_at ASC").
		// 方式二:
		//Order(clause.OrderBy{
		//	Columns: []clause.OrderByColumn{
		//		{Column: clause.Column{Name: "update_at"}, Desc: true},
		//		{Column: clause.Column{Name: "create_at"}, Desc: false},
		//	},
		//	Expression: nil,
		//}).
		Find(&arts).Error
	return arts, err
}

func (a *gormArticleDao) GetById(ctx context.Context, id int64) (dao.ArticleModel, error) {
	var art dao.ArticleModel
	err := a.db.WithContext(ctx).Model(&dao.ArticleModel{}).Where("id = ?", id).First(&art).Error
	return art, err
}

func (a *gormArticleDao) GetPubById(ctx context.Context, id int64) (dao.PublishArticleModels, error) {
	var art dao.PublishArticleModels
	err := a.db.WithContext(ctx).Model(&dao.ArticleModel{}).Where("id = ?", id).First(&art).Error
	return art, err
}

// Transaction
//
//	@Description:
func (a *gormArticleDao) Transaction(ctx context.Context, bizFunc func(txDao ArticleDAO) error) error {
	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txDao := NewGormArticleDao(tx)
		return bizFunc(txDao)
	})
}

// Upsert
//
//	@Description: GORM实现Upsert的语义 Insert Or Update
func (a *gormArticleDao) Upsert(ctx context.Context, art dao.PublishArticleModels) error {
	now := time.Now()
	// OnConflict的意思是数据冲突了  用MySQL就只需要关注一个地方 DoUpdates
	var err = a.db.Clauses(clause.OnConflict{
		// 哪些列冲突了
		//Columns: []clause.Column{clause.Column{Name: "id"}},
		// 如果数据冲突了 就啥也不干
		//DoNothing: true,
		// 如果数据冲突了，并且符合Where条件，就会执行更新
		//Where:
		// MySQL只需要关心这个 其他SQL标准的就不一样啦！
		DoUpdates: clause.Assignments(map[string]interface{}{
			"title": art.Title,
			//"author_id":    art.AuthorId, // 创作者不能修改
			"summary":      art.Summary,
			"content":      art.Content,
			"content_type": art.ContentType,
			"cover":        art.Cover,
			"status":       art.Status,
			"updated_at":   now,
		}),
	}).Create(&art).Error // 最终生成的语句 INSERT xxx ON DUPLICATE KEY UPDATE xxx
	return err
}

func (a *gormArticleDao) Sync(ctx context.Context, art dao.ArticleModel) (int64, error) {
	// 这里采用闭包形态操作事务 GORM帮我们管理了事务的生命周期
	var id = int64(art.ID)
	err := a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var err error
		txDao := NewGormArticleDao(tx)
		if id > 0 {
			err = txDao.UpdateById(ctx, art)
		} else {
			id, err = txDao.Insert(ctx, art)
		}
		if err != nil {
			return err
		}
		// 操作线上表了
		return txDao.Upsert(ctx, dao.PublishArticleModels{art})
	})
	return id, err
}

func (a *gormArticleDao) SyncStatus(ctx context.Context, author, id int64, status uint8) error {
	// 也要开启事务
	now := time.Now()
	err := a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&dao.ArticleModel{}).
			Where("id = ? AND author_id = ?", id, author).
			Updates(map[string]any{
				"status":     status,
				"updated_at": now,
			})
		if res.Error != nil {
			// 数据库有问题
			return res.Error
		}
		if res.RowsAffected != 1 {
			// 要么Id不存在，要么作者不对
			return fmt.Errorf("文章不存在 Id: %d 作者: %d", id, author)
		}
		// 操作线上库
		return tx.Model(&dao.PublishArticleModels{}).
			Where("id = ?", id). // 这里因为上面已经过滤了authorid，所以这里不需要再次查询
			Updates(map[string]any{
				"status":     status,
				"updated_at": now,
			}).Error
	})
	return err
}

func (a *gormArticleDao) ListPub(ctx context.Context, start time.Time, offset int, limit int) ([]dao.ArticleModel, error) {
	//TODO implement me
	panic("implement me")
}
