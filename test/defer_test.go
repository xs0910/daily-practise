package test

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	var result int
	a()
	b()
	result = c()
	fmt.Println(result)

	result = d()
	fmt.Println(result)

	result = e()
	fmt.Println(result)

	result = f()
	fmt.Println(result)
}

func a() {
	i := 0
	defer fmt.Println(i) // 输出的是0
	i++
	defer fmt.Println(i) // 输出的是1
	return
	// 输出结果 1 0
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func c() (result int) {
	defer func() {
		result++
	}()
	return 0
	// 返回值是1，在defer中被更改了
}

func d() (r int) {
	t := 5
	r = t
	defer func() {
		t = t + 5
	}()
	return t
	// 返回值是5，在defer中并没有修改r的值
}

func e() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
	// 返回值是1，defer的传入参数是值类型，并不会改变返回结果r的值
}

func f() (r int) {
	defer func(r *int) {
		*r = *r + 5
	}(&r)
	return 1
	// 返回值是6，defer的传入参数是引用类型，取地址操作会改变最终r的值
}

func TestPanic(t *testing.T) {
	panicTest()
}

func panicTest() {
	defer func() { fmt.Println(1) }()
	defer func() { fmt.Println(2) }()
	//panic("手动触发异常")
	fmt.Println("触发异常，将无法执行")
}

func TestRecover(t *testing.T) {
	recoverTest()
}

func recoverTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover捕获到panic")
			fmt.Println(err)
		}
	}()

	fmt.Println("recoverTest运行开始")
	panic("运行出现异常")
}
