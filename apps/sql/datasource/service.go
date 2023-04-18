package datasource

import (
	"context"
	"database/sql"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/sqlbuilder"
	"go-notebook/apps"
	"go-notebook/conf"
)

var impl = &DatasourceServiceImpl{}

// NewDatasourceServiceImpl 保证调用该函数之前, 全局conf对象已经初始化
func NewDatasourceServiceImpl() *DatasourceServiceImpl {
	return &DatasourceServiceImpl{
		l: zap.L().Named("Datasource"),
	}
}

type DatasourceServiceImpl struct {
	l  logger.Logger
	db *sql.DB
}

func (i *DatasourceServiceImpl) Query(ctx context.Context, req *QueryDatasourceRequest) (*DatasourceSet, error) {
	b := sqlbuilder.NewBuilder(SelectSQL)

	if req.Keywords != "" {
		b.Where("(name like ? or alias like ?)",
			"%"+req.Keywords+"%",
			"%"+req.Keywords+"%",
		)
	}
	if req.Cate > 0 {
		b.Where("cate = ?", req.Cate)
	}
	//b.Order("update_time").Desc()
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

	set := NewDatasourceSet()
	for rows.Next() {
		item := NewDatasource()
		if err := rows.Scan(
			&item.Id, &item.Name, &item.Cate, &item.Alias, &item.Introduction, &item.Sql, &item.CreateTime, &item.UpdateTime,
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

func (i *DatasourceServiceImpl) Config() {
	i.l = zap.L().Named("Datasource")
	i.db = conf.C().Sqlite.GetDB()
}

func (i *DatasourceServiceImpl) Name() string {
	return registerName()
}

func init() {
	apps.RegistryImpl(impl)
}
