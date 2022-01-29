package stdlib

import (
	"reflect"
	"testing"
)

func TestCommon(t *testing.T) {

}

type Plan struct {
	PlanTitle string `json:"plan_title" wxp:"planTitle"`
	PlanCode  string `json:"plan_code" wxp:"planCode"`
}

// 读取结构体标签
func TestStructTag(t *testing.T) {
	// 这是一个接口类型值，可以打开注释run一下，效果是一样的。
	// 原因：reflect.TypeOf的返回值代表的是 接口类型的动态类型。
	//var plan interface{} = Plan{}

	// 这是一个非接口类型值，可以打开注释run一下，效果是一样的。
	var plan = Plan{}

	planType := reflect.TypeOf(plan)
	for i := 0; i < planType.NumField(); i++ {
		t.Log(planType.Field(i).Tag.Get("wxp"))
	}
}
