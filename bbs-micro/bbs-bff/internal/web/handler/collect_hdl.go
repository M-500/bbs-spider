package handler

import (
	"bbs-micro/bbs-bff/internal/service"
	"bbs-micro/bbs-bff/internal/web/jwtx"
	"bbs-micro/bbs-bff/internal/web/vo"
	"bbs-micro/pkg/ginplus"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-12 11:23

type CollectHandler struct {
	svc service.ICollectService
	biz string
}

func NewCollectHandler(svc service.ICollectService) *CollectHandler {
	return &CollectHandler{
		svc: svc,
		biz: "article",
	}
}

func (h CollectHandler) CreateCollect(ctx *gin.Context, req vo.CreateCollectReq, user jwtx.UserClaims) (ginplus.Result, error) {
	collect, err := h.svc.CreateCollect(ctx, user.Id, req.CollectName, req.Desc, req.IsPublic)
	if err != nil {
		return ginplus.Result{
			Code: 5003402,
			Msg:  "创建失败",
		}, err
	}
	return ginplus.Result{
		Data: collect,
		Msg:  "OK",
	}, nil
}

func (h CollectHandler) GetCollectById(ctx *gin.Context) (ginplus.Result, error) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ginplus.Result{
			Code: 503401,
			Msg:  "参数错误",
		}, nil
	}
	collects, err := h.svc.GetByUid(ctx, id, 0, 1)
	if err != nil {
		return ginplus.Result{
			Code: 503001,
			Msg:  "系统错误",
		}, nil
	}
	return ginplus.Result{
		Data: collects,
	}, nil
}

func (h CollectHandler) CollectEntityByID(ctx *gin.Context, user jwtx.UserClaims) (ginplus.Result, error) {
	cidStr := ctx.Param("cid")
	bidStr := ctx.Param("bid")
	cid, err := strconv.ParseInt(cidStr, 10, 64)
	if err != nil {
		return ginplus.Result{
			Code: 503401,
			Msg:  "参数错误",
		}, nil
	}
	bid, err := strconv.ParseInt(bidStr, 10, 64)
	if err != nil {
		return ginplus.Result{
			Code: 503401,
			Msg:  "参数错误",
		}, nil
	}
	entity, err := h.svc.CollectEntity(ctx, h.biz, user.Id, cid, bid)
	if err != nil {

	}

	return ginplus.Result{
		Data: entity,
	}, nil
}
