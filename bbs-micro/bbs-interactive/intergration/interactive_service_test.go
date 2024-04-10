package intergration

import (
	intrv1 "bbs-micro/api/proto/gen/proto/intr/v1"
	"bbs-micro/bbs-interactive/grpc"
	"bbs-micro/bbs-interactive/intergration/dep_setup"
	"bbs-micro/bbs-interactive/repository/dao"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
	"time"
)

// 集成测试

type InteractiveTestSuite struct {
	suite.Suite
	rdb    redis.Cmdable
	db     *gorm.DB
	server *grpc.InteractiveServiceServer
}

// SetupSuite
//
//	@Description: 开始测试套件之前
//	@receiver s
func (s *InteractiveTestSuite) SetupSuite() {
	s.db = dep_setup.InitTestDB()
	s.rdb = dep_setup.InitRedis()
	s.server = dep_setup.InitInteractiveGRPCServer()
}

// TearDownTest
//
//	@Description: 结束的时候 通常要清空数据库的所有数据
//	@receiver s
func (s *InteractiveTestSuite) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	err := s.db.Exec("TRUNCATE TABLE `interactive`").Error
	assert.NoError(s.T(), err)
	err = s.db.Exec("TRUNCATE TABLE `user_to_biz_like`").Error
	assert.NoError(s.T(), err)
	err = s.db.Exec("TRUNCATE TABLE `user_to_biz_collect`").Error
	assert.NoError(s.T(), err)
	// 清空 Redis
	err = s.rdb.FlushDB(ctx).Err()
	assert.NoError(s.T(), err)
}

func (s *InteractiveTestSuite) TestIncrReadCnt() {
	testCases := []struct {
		name   string
		before func(t *testing.T)
		after  func(t *testing.T)

		biz   string
		bizId int64

		wantErr  error
		wantResp *intrv1.IncrReadCntResponse
	}{
		{
			name: "增加成功,db和redis", // DB 和缓存都有数据
			before: func(t *testing.T) { // 先要在redis中set一个key
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				defer cancel()
				err := s.db.Create(&dao.InteractiveModel{
					Model: gorm.Model{
						ID: 1,
					},
					BizId:   2,
					Biz:     "test",
					ReadCnt: 3,
					LikeCnt: 4,
				}).Error
				assert.NoError(t, err)
				err = s.rdb.HSet(ctx, "interactive:test:2", "read_cnt", 3).Err()
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				defer cancel()
				var data dao.InteractiveModel
				err := s.db.Where("id = ?", 2).First(&data).Error
				assert.NoError(t, err)
				assert.True(t, data.CreatedAt.Before(data.UpdatedAt)) // 因为无法确定这个更新时间一定大于当前时间 不好测试，所以用这个方式
				data.CreatedAt = time.Time{}
				data.UpdatedAt = time.Time{}
				assert.Equal(t, dao.InteractiveModel{
					Model:   gorm.Model{ID: 1},
					BizId:   2,
					Biz:     "test",
					ReadCnt: 4,
					LikeCnt: 4,
				}, data)
				// 校验redis的值是否相同
				cnt, err := s.rdb.HGet(ctx, "interactive:test:2", "read_cnt").Int()
				assert.NoError(t, err)
				assert.Equal(t, 4, cnt) // 如果没有问题，那么redis读取出来的数据显然应该自增过 = 4
				err = s.rdb.Del(ctx, "interactive:test:2").Err()
				assert.NoError(t, err)
			},
			biz:      "test",
			bizId:    2,
			wantErr:  nil,
			wantResp: &intrv1.IncrReadCntResponse{},
		},
	}
	// 不同于 AsyncSms 服务，我们不需要 mock，所以创建一个就可以
	// 不需要每个测试都创建
	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			tc.before(t)
			resp, err := s.server.IncrReadCnt(context.Background(), &intrv1.IncrReadCntRequest{
				Biz: tc.biz, BizId: tc.bizId,
			})
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantResp, resp)
			tc.after(t)
		})
	}
}
