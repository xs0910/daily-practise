package test

import (
	"fmt"
	"testing"
)

type T1 struct {
	id int
}

func TestIteration(t *testing.T) {
	// iteration1 是错误结果
	iteration1()
	// iteration2/iteration3 是正确结果
	iteration2()
	iteration3()
}

func iteration1() {
	t1 := T1{id: 1}
	t2 := T1{id: 2}
	ts1 := []T1{t1, t2}
	var ts2 []*T1
	for _, t := range ts1 {
		ts2 = append(ts2, &t)
	}
	for _, t := range ts2 {
		fmt.Println((*t).id)
	}
	// 输出结果是 2 2
	// 迭代变量t使用短变量声明的方式声明, 它的声明周期就是for代码块.
	// 这个变量在第一次循环时是第一个元素的值, 在第二次循环时是第二个元素的值, 但是在内存的某个地方保存着slice被遍历结束时的值.
	// t没有指向slice底层数组指向的值--这是一个临时的桶, 下一个元素会覆盖当前值.
	// t是一个辅助变量来保存当前迭代的元素, 所以&t每次循环都是相同的值
}

func iteration2() {
	t1 := T1{id: 1}
	t2 := T1{id: 2}
	ts1 := []T1{t1, t2}
	var ts2 []*T1
	for _, t := range ts1 {
		tmp := t
		ts2 = append(ts2, &tmp)
	}
	for _, t := range ts2 {
		fmt.Println((*t).id)
	}
	// 输出结果是 1 2
}

func iteration3() {
	t1 := T1{id: 1}
	t2 := T1{id: 2}
	ts1 := []T1{t1, t2}
	var ts2 []*T1
	for i, _ := range ts1 {
		ts2 = append(ts2, &ts1[i])
	}
	for _, t := range ts2 {
		fmt.Println((*t).id)
	}
	// 输出结果是 1 2
}
