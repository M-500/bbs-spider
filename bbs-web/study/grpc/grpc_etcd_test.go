//@Author: wulinlin
//@Description:
//@File:  grpc_etcd_testr
//@Version: 1.0.0
//@Date: 2024/04/21 16:14

package grpc

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type EtcdTestSuide struct {
	suite.Suite
}

func (s *EtcdTestSuide) TestServer() {

}

func TestEtcd(t *testing.T) {
	suite.Run(t, new(EtcdTestSuide))
}
