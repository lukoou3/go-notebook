package apps

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin"
)

// IOC 容器层: 管理所有的服务的实例

var (
	// 维护当前所有的服务
	implApps    = map[string]ImplService{}
	ginApps     = map[string]GinService{}
	restfulApps = map[string]RestfulService{}
)

type ImplService interface {
	Config()
	Name() string
}

// 注册Gin编写的Handler
// 比如 编写了Http服务A, 只需要实现Registry方法, 就能把Handler注册给Root Router
type GinService interface {
	Registry(r gin.IRouter)
	Config()
	Name() string
}

type RestfulService interface {
	Registry(ws *restful.WebService)
	Config()
	Name() string
}

func RegistryGin(svc GinService) {
	// 服务实例注册到svcs map当中
	if _, ok := ginApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	ginApps[svc.Name()] = svc
}

func RegistryImpl(svc ImplService) {
	// 服务实例注册到svcs map当中
	if _, ok := implApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	implApps[svc.Name()] = svc
}

// 用户初始化 注册到Ioc容器里面的所有服务
func InitImpl() {
	for _, v := range implApps {
		v.Config()
	}
}

// Get 一个Impl服务的实例：implApps
// 返回一个对象, 任何类型都可以, 使用时, 由使用方进行断言
func GetImpl(name string) interface{} {
	for k, v := range implApps {
		if k == name {
			return v
		}
	}

	return nil
}

// 用户初始化 注册到Ioc容器里面的所有服务
func InitGin(r gin.IRouter) {
	// 先初始化好所有对象
	for _, v := range ginApps {
		v.Config()
	}

	// 完成Http Handler的注册
	for _, v := range ginApps {
		v.Registry(r)
	}
}

// 已经加载完成的Gin App由哪些
func LoadedGinApps() (names []string) {
	for k := range ginApps {
		names = append(names, k)
	}

	return
}
