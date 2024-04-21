//@Author: wulinlin
//@Description:
//@File:  grpc_etcd_testr
//@Version: 1.0.0
//@Date: 2024/04/21 16:14

package grpc

import (
	"context"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type EtcdTestSuide struct {
	suite.Suite
	client etcdv3.Client
}

func (s *EtcdTestSuide) SetupSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, err := etcdv3.NewClient(ctx, []string{"localhost:12379"}, etcdv3.ClientOptions{})
	require.NoError(s.T(), err)
	s.client = client

}
func (s *EtcdTestSuide) TestServer() {
	em, err := NewManager
}

func TestEtcd(t *testing.T) {
	suite.Run(t, new(EtcdTestSuide))
}
