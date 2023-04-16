package shell

import (
	"context"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go-notebook/apps"
	"go-notebook/apps/code"
)

// 接口实现的静态检查
// 这样写, 会造成 conf.C()并准备好, 造成conf.C().MySQL.GetDB()该方法pannic
// var impl = NewShellCodeServiceImpl()

// 把对象的注册和对象的注册这2个逻辑独立出来
var impl = &ShellCodeServiceImpl{}

// NewShellCodeServiceImpl 保证调用该函数之前, 全局conf对象已经初始化
func NewShellCodeServiceImpl() *ShellCodeServiceImpl {
	return &ShellCodeServiceImpl{
		l: zap.L().Named("ShellCode"),
	}
}

type ShellCodeServiceImpl struct {
	l logger.Logger
}

func (i *ShellCodeServiceImpl) Query(ctx context.Context, req *QueryShellCodeRequest) (*ShellCodeSet, error) {
	var items []*ShellCode
	return &ShellCodeSet{Total: 20, Items: items}, nil
}

// 只需要保证 全局对象Config和全局Logger已经加载完成
func (i *ShellCodeServiceImpl) Config() {
	// Host service 服务的子Loggger
	// 封装的Zap让其满足 Logger接口
	// 为什么要封装:
	// 		1. Logger全局实例
	// 		2. Logger Level的动态调整, Logrus不支持Level共同调整
	// 		3. 加入日志轮转功能的集合
	i.l = zap.L().Named("Host")
}

// 服务服务的名称
func (i *ShellCodeServiceImpl) Name() string {
	return code.AppName + "-" + ModelName
}

// _ import app 自动执行注册逻辑
func init() {
	//  对象注册到ioc层
	apps.RegistryImpl(impl)
}
