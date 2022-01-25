package main

import "fmt"

type User struct {
	ID     int64
	Name   string
	Avatar string
}

// GetUserInfo1 示例1
func GetUserInfo1() *User {
	return &User{ID: 13746731, Name: "test", Avatar: "https://image.baidu.com/search/wisemidresult?tn=wisemidresult&word=清纯美女&pn=0&rn=6&size=mid&sp=5&iswiseala=1&ie=utf8&fmpage=index&pos=aimeinv2"}
}

// GetUserInfo2 示例2
func GetUserInfo2(u *User) *User {
	return u
}

// GetUserInfo3 示例3
func GetUserInfo3(u User) *User {
	return &u
}

func main() {
	_ = GetUserInfo1()

	str := new(string)
	*str = "test" // 有没有被作用域之外所引用,这里的作用域仍然保留在 main 中，因此它没有发生逃逸。

	fmt.Println("fmt") // 造成了从栈到堆的分配,得知当形参为 interface 类型时，在编译阶段编译器因为无法确定其具体的类型，因此会造成逃逸，最终将该变量分配到堆上。
	// 从源码来讲的话，实际上是该方法内部的 reflect.TypeOf(arg).Kind() 语句造成逃逸，因此表象就是 interface 类型会导致该对象分配到堆上。

	// leaking param: u to result ~r1 level=0   &User literal does not escape
	_ = GetUserInfo2(&User{ID: 13746731, Name: "test", Avatar: "https://image.baidu.com/search/wisemidresult?tn=wisemidresult&word=清纯美女&pn=0&rn=6&size=mid&sp=5&iswiseala=1&ie=utf8&fmpage=index&pos=aimeinv2"})

	// moved to heap: u
	_ = GetUserInfo3(User{ID: 13746731, Name: "test", Avatar: "https://image.baidu.com/search/wisemidresult?tn=wisemidresult&word=清纯美女&pn=0&rn=6&size=mid&sp=5&iswiseala=1&ie=utf8&fmpage=index&pos=aimeinv2"})

	// 命令
	// go build -gcflags '-m -l' main.go
	// go tool compile -S main.go

	// 结果
	//	.\main.go:13:9: &User literal escapes to heap
	//	.\main.go:17:19: leaking param: u to result ~r1 level=0
	//	.\main.go:22:19: moved to heap: u
	//	.\main.go:29:12: new(string) does not escape
	//	.\main.go:32:13: ... argument does not escape
	//	.\main.go:32:14: "fmt" escapes to heap
	//	.\main.go:36:19: &User literal does not escape
}
