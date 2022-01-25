package gomock

import (
	spider "daily-practise/gomock/spider/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetGoVersion(t *testing.T) {
	// 首先，需要在单元测试代码里创建一个 Mock 控制器，将*testing.T传递给 GoMock，生成一个Controller对象，该对象控制了整个 Mock 的过程。
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 然后，就可以调用 Mock 的对象了，这里的spider是 mockgen 命令里面传递的包名，后面是 NewMockXxx 格式的对象创建函数，Xxx是接口名
	mockSpider := spider.NewMockSpider(ctrl)

	// 接着，有了 Mock 实例，我们就可以调用其断言方法EXPECT()了。
	// Mock 一个接口的方法，我们需要 Mock 该方法的入参和返回值。我们可以通过参数匹配来 Mock 入参，通过 Mock 实例的 Return 方法来 Mock 返回值
	mockSpider.EXPECT().GetBody().Return("go1.17.6")
	// mockSpider.EXPECT().GetBody(gomock.Any(), gomock.Eq("admin")).Return("go1.8.3")
	// gomock.Any()：     可以用来表示任意的入参。
	// gomock.Eq(value)： 用来表示与 value 等价的值。
	// gomock.Not(value)：用来表示非 value 以外的值。
	// gomock.Nil()：     用来表示 None 值。

	goVer := GetGoVersion(mockSpider)

	if goVer != "go1.17.6" {
		t.Errorf("Get wrong veresion %s", goVer)
	}
}
