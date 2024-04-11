//@Author: wulinlin
//@Description:
//@File:  types
//@Version: 1.0.0
//@Date: 2024/04/11 23:40

package job

type Job interface {
	Name() string
	Run() error
}
