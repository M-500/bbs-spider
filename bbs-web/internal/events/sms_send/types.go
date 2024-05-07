//@Author: wulinlin
//@Description:
//@File:  types
//@Version: 1.0.0
//@Date: 2024/05/05 12:14

package sms_send

import "context"

type SMSSendEvent struct {
	Uid   int64  // 用户ID
	Phone string // 手机号
}

type SMSProducer interface {
	ProducerSMSSendEvent(ctx context.Context, evt SMSSendEvent) error
}
