package datasource

import (
	"context"
	"go-notebook/apps/sql"
)

const (
	ModelName = "datasources"
)

func registerName() string {
	return sql.AppName + "-" + ModelName
}

type DatasourceService interface {
	Query(context.Context, *QueryDatasourceRequest) (*DatasourceSet, error)
}
