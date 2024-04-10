package dao

import "context"

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 12:50

type InteractiveDao interface {
	IncrReadCnt(ctx context.Context, biz string, bizId int64) error
	BatchIncrReadCnt(ctx context.Context, bizs []string, bizIds []int64) error
	IncrLikeInfo(ctx context.Context, biz string, id int64, uid int64) error
	DelLikeInfo(ctx context.Context, biz string, bizId int64, uid int64) error
	IncrLikeCnt(ctx context.Context, biz string, id int64, uid int64) error
	DecrLikeCnt(ctx context.Context, biz string, id int64, uid int64) error
	Get(ctx context.Context, biz string, id int64) (InteractiveModel, error)
	GetLikeInfo(ctx context.Context, biz string, id int64, uid int64) (UserLikeBizModel, error)
	GetCollectInfo(ctx context.Context, biz string, id int64, uid int64) (UserCollectBizModel, error)
}

type interactiveDao struct {
}

func NewInteractiveDao() InteractiveDao {
	return &interactiveDao{}
}

func (dao *interactiveDao) IncrReadCnt(ctx context.Context, biz string, bizId int64) error {
	//TODO implement me
	panic("implement me")
}

func (dao *interactiveDao) BatchIncrReadCnt(ctx context.Context, bizs []string, bizIds []int64) error {
	//TODO implement me
	panic("implement me")
}

func (dao *interactiveDao) IncrLikeInfo(ctx context.Context, biz string, id int64, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (dao *interactiveDao) DelLikeInfo(ctx context.Context, biz string, bizId int64, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (dao *interactiveDao) IncrLikeCnt(ctx context.Context, biz string, id int64, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (dao *interactiveDao) DecrLikeCnt(ctx context.Context, biz string, id int64, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (dao *interactiveDao) Get(ctx context.Context, biz string, id int64) (InteractiveModel, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *interactiveDao) GetLikeInfo(ctx context.Context, biz string, id int64, uid int64) (UserLikeBizModel, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *interactiveDao) GetCollectInfo(ctx context.Context, biz string, id int64, uid int64) (UserCollectBizModel, error) {
	//TODO implement me
	panic("implement me")
}
