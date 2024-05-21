package base_study

import (
	"fmt"
	"reflect"
	"testing"
)

// @Description 值传递还是引用传递
// @Author 代码小学生王木木

func TestParamDeep(t *testing.T) {
	num := 5
	changeValue(num)
	fmt.Println(num) // 输出 5
}

// 基本类型值传递
func changeValue(val int) {
	val = 10
}

func changeList(nums []int) {
	nums[0] = 10
}
func TestListDeep(t *testing.T) {
	array := []int{1, 2, 3}
	t.Log("类型为", reflect.ValueOf(array).Kind())
	changeList(array)
	fmt.Println(array) // 输出 [1 2 3]
}

func TestUintOverflow(t *testing.T) {
	var a uint8 = 1
	var b uint8 = 2
	t.Log(a - b)
}
func showCat() {
	fmt.Printf("cat")
}
func showDog() {
	fmt.Printf("dog")
}
func showFish() {
	fmt.Printf("fish")
}

type User struct {
	Name string
}

type User1 struct {
	Name string

	age User
}

func TestChannel(t *testing.T) {
	var user = User1{}
	var user2 = User1{}
	t.Log(user == user2)
}
