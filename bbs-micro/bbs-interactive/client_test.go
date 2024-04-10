package main

import (
	intrv1 "bbs-micro/api/proto/gen/proto/intr/v1"
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 19:03

func TestGRPCConnect(t *testing.T) {
	cc, err := grpc.Dial("127.0.0.1:8090",
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	require.NoError(t, err)
	client := intrv1.NewInteractiveServiceClient(cc)
	get, err := client.IncrReadCnt(context.Background(), &intrv1.IncrReadCntRequest{
		Biz: "test", BizId: 2,
	})
	require.NoError(t, err)
	t.Log(get)
}
