package shell

import (
	"context"
	"database/sql"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/sqlbuilder"
	"go-notebook/apps"
	"go-notebook/apps/code"
	"go-notebook/conf"
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
	l  logger.Logger
	db *sql.DB
}

func (i *ShellCodeServiceImpl) Query(ctx context.Context, req *QueryShellCodeRequest) (*ShellCodeSet, error) {
	b := sqlbuilder.NewBuilder(SelectSQL)

	if req.Keywords != "" {
		// (r.`name`='%' OR r.description='%' OR r.private_ip='%' OR r.public_ip='%')
		//  10.10.1, 接口测试
		b.Where("name like ? or code like ? or code desc ?",
			"%"+req.Keywords+"%",
			"%"+req.Keywords+"%",
			"%"+req.Keywords+"%",
		)
	}
	b.Limit(req.OffSet(), req.GetPageSize())
	querySQL, args := b.Build()

	i.l.Infof("query sql: %s, args: %v", querySQL, args)

	stmt, err := i.db.PrepareContext(ctx, querySQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	set := NewShellCodeSet()
	for rows.Next() {
		item := NewShellCode()
		if err := rows.Scan(
			&item.Id, &item.Name, &item.Code, &item.Desc, &item.CreateTime, &item.UpdateTime,
		); err != nil {
			return nil, err
		}
		//item.Code = ""
		set.Add(item)
	}

	// total统计
	countSQL, args := b.BuildCount()
	i.l.Infof("count sql: %s, args: %v", countSQL, args)
	countStmt, err := i.db.PrepareContext(ctx, countSQL)
	if err != nil {
		return nil, err
	}
	defer countStmt.Close()
	if err := countStmt.QueryRowContext(ctx, args...).Scan(&set.Total); err != nil {
		return nil, err
	}

	return set, nil
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
	i.db = conf.C().Sqlite.GetDB()
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
