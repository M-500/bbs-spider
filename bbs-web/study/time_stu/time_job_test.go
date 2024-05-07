//@Author: wulinlin
//@Description:
//@File:  time_job_test
//@Version: 1.0.0
//@Date: 2024/05/04 17:35

package time_stu

import (
	"fmt"
	"testing"
	"time"
)

func TestAfterFunc(t *testing.T) {
	timer := time.AfterFunc(time.Second*3, func() {
		fmt.Println("执行任务啦")
	})

	time.Sleep(time.Second * 5) // 防止提前退出
	timer.Stop()                // 停止定时器
}
