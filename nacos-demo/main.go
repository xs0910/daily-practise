package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"time"
)

/*
nacos-go-sdk 简单使用
https://github.com/nacos-group/nacos-sdk-go/blob/master/README_CN.md
Client Create
1. 创建clientConfig
2. 创建ServiceConfig
3. 创建服务发现客户端
4. 创建动态配置客户端
服务发现
5. 注册实例RegisterInstance
6. 注销实例DeRegisterInstance
7. 获取服务信息GetService
8. 获取所有的实例列表SelectAllInstance
9. 获取实例列表SelectInstances
10.获取一个健康的实例（加权随机轮询）SelectOneHealthyInstance
11.监听服务变化Subscribe
12.取消服务监听Unsubscribe
13.获取服务名列表GetAllServicesInfo
动态配置
14.发布配置PublishConfig
15.删除配置DeleteConfig
16.获取配置GetConfig
17.监听配置变化ListenConfig
18.取消配置监听CancelListenConfig
19.搜索配置SearchConfig
*/
func main() {
	// ServerConfig：至少一个ServiceConfig,我们可以配置多个ServerConfig，客户端会对这些服务端做轮询请求
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "192.168.81.110",
			Port:   8848,
		},
	}

	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		// 2cf5ca94-d27f-43cd-82a6-de06c56f4588 是自定义了一个命名空间的Id
		NamespaceId:         "2cf5ca94-d27f-43cd-82a6-de06c56f4588", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建动态配置客户端的另一种方式(推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		panic(err)
	}

	// 创建服务发现客户端的另一种方式(推荐)
	namingClient, err := clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  &clientConfig,
		ServerConfigs: serverConfigs,
	})

	// 注册实例：RegisterInstance
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "192.168.81.110",
		Port:        8848,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ServiceName: "demo.go",
		ClusterName: "DEFAULT",       // 默认值DEFAULT
		GroupName:   "DEFAULT_GROUP", // 默认值DEFAULT_GROUP
		Ephemeral:   true,
	})
	fmt.Printf("注册实例结果：%t\n", success)
	if err != nil {
		return
	}

	// 获取服务信息：GetService
	service, err := namingClient.GetService(vo.GetServiceParam{
		ServiceName: "demo.go",
		Clusters:    []string{"DEFAULT"}, // 默认值DEFAULT
		GroupName:   "DEFAULT_GROUP",     // 默认值DEFAULT_GROUP
	})
	if err != nil {
		return
	}
	marshal, _ := json.Marshal(service)
	fmt.Println("获取服务信息:" + string(marshal))

	time.Sleep(time.Second * 10)

	// 注销实例：DeregisterInstance
	//success, err = namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
	//	Ip:          "192.168.81.110",
	//	Port:        8848,
	//	ServiceName: "demo.go",
	//	Cluster:     "DEFAULT",       // 默认值DEFAULT
	//	GroupName:   "DEFAULT_GROUP", // 默认值DEFAULT_GROUP
	//	Ephemeral:   true,
	//})
	//fmt.Printf("注销实例结果：%t\n", success)
	//if err != nil {
	//	return
	//}

	// 监听配置
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "nacos.cfg.dataId",
		Group:  "test",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ",dataId:" + dataId + ",data:" + data)
		},
	})

	if err != nil {
		return
	}

	time.Sleep(time.Second * 1000)
}
