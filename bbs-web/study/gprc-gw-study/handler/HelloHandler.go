package handler

import (
	proto "bbs-web/study/gprc-gw-study/pb/hello"
	"context"
)

// @Description
// @Author 代码小学生王木木

type HelloServer struct {
	proto.UnimplementedHelloServer
}

func (h *HelloServer) SayHello(ctx context.Context, req *proto.HelloReq) (*proto.HelloResp, error) {
	res := new(proto.HelloResp)
	res.Value = "你好呀" + req.Key
	return res, nil
}
