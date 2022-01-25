package test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type T struct {
	f1     string `one:"1" two:"2"`
	f2     string
	f3     string `f three`
	f4, f5 int64  `f four and five`
}

func TestReflectTag(t *testing.T) {
	reflectTest()
}

func reflectTest() {
	t := reflect.TypeOf(T{})
	f1, _ := t.FieldByName("f1")
	fmt.Println(f1.Tag)

	v, ok := f1.Tag.Lookup("one")
	fmt.Printf("%s, %t\n", v, ok)

	f4, _ := t.FieldByName("f4")
	fmt.Println(f4)

	f5, _ := t.FieldByName("f5")
	fmt.Println(f5.Tag)
}

func TestReflectJson(t *testing.T) {
	type T struct {
		F1 int `json:"f_1"`
		F2 int `json:"f_2,omitempty"`
		F3 int `json:"f_3,omitempty"`
		F4 int `json:"-"`
	}
	tt := T{1, 2, 0, 4}
	b, err := json.Marshal(tt)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Printf("%s\n", b)
	// 结果是：{"f_1":1,"f_2":2}，F3=0:表示空值，不会在JSON中出现
}
