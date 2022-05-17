package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf // rest api 配置
	Auth          struct {
		AccessSecret string // jwt密钥
		AccessExpire int64  // 有效期，单位：秒
	}
	MySql struct {
		DataSource string // mysql 连接地址
	}
	CacheRedis cache.CacheConf    // redis缓存
	UserRpc    zrpc.RpcClientConf // rpc client 配置
}
