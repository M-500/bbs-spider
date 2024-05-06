package map_v1

import "testing"

// @Description
// @Author 代码小学生王木木

func TestMap(t *testing.T) {
	var a = make(map[string]int, 2)
	a["wulinlin"] = 123
	a["wulinlin2"] = 456
	t.Log(a)
}
