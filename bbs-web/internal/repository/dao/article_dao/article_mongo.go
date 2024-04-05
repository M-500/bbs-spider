package article_dao

import (
	"context"
	"time"

	"github.com/bwmarrin/snowflake"
	"go.mongodb.org/mongo-driver/mongo"

	"bbs-web/internal/repository/dao"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-05 17:41

type mongoArticleDao struct {
	client *mongo.Client
	// 代表数据库
	database *mongo.Database
	// 代表制作库
	col *mongo.Collection
	// 代表线上库
	livCol *mongo.Collection
	// 用于雪花算法
	node snowflake.Node
}

func NewMongoArticleDao(client *mongo.Client, database *mongo.Database, col *mongo.Collection, livCol *mongo.Collection) *mongoArticleDao {
	return &mongoArticleDao{client: client, database: database, col: col, livCol: livCol}
}

func (m *mongoArticleDao) Insert(ctx context.Context, art dao.ArticleModel) (int64, error) {
	id := m.node.Generate().Int64()
	art.ID = uint(id)
	_, err := m.col.InsertOne(ctx, art)
	if err != nil {
		return 0, err
	}
	// 但是这里搞不到主键啊，MongoDB的Object ID是12位的
	/**
	1. 接口返回的类型 全部换成string类型，同事修改gorm对应的实现
	2. 使用ID生成策略，比如说雪花算法
	3. 定义一个新的接口
	4. 使用GUID  全局唯一ID  global Unify ID
	*/
	return id, nil
}

func (m *mongoArticleDao) UpdateById(ctx context.Context, art dao.ArticleModel) error {
	//TODO implement me
	panic("implement me")
}

func (m *mongoArticleDao) GetByAuthor(ctx context.Context, author int64, offset, limit int) ([]dao.ArticleModel, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mongoArticleDao) GetById(ctx context.Context, id int64) (dao.ArticleModel, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mongoArticleDao) GetPubById(ctx context.Context, id int64) (dao.ArticleModel, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mongoArticleDao) Transaction(ctx context.Context, bizFunc func(txDao ArticleDAO) error) error {
	//TODO implement me
	panic("implement me")
}

func (m *mongoArticleDao) Sync(ctx context.Context, art dao.ArticleModel) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mongoArticleDao) Upsert(ctx context.Context, art dao.PublishArticleModels) error {
	//TODO implement me
	panic("implement me")
}

func (m *mongoArticleDao) SyncStatus(ctx context.Context, author, id int64, status uint8) error {
	//TODO implement me
	panic("implement me")
}

func (m *mongoArticleDao) ListPub(ctx context.Context, start time.Time, offset int, limit int) ([]dao.ArticleModel, error) {
	//TODO implement me
	panic("implement me")
}
