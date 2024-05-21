package reflect

import (
	"reflect"
	"testing"
)

// @Description
// @Author 代码小学生王木木

type user struct {
	Name string
}

func (s *user) GetName() {

}
func (s *user) GetName1() {

}
func (s *user) GetName2() {

}

func TestGetMethod(t *testing.T) {
	var u = &user{}
	sT := reflect.TypeOf(u)
	for s := 0; s < sT.NumMethod(); s++ {
		t.Log(sT.Method(s))
	}
}
