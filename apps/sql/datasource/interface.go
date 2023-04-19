package datasource

import (
	"context"
	"go-notebook/apps/sql"
)

const (
	ModelName = "datasource"
)

func registerName() string {
	return sql.AppName + "-" + ModelName
}

type DatasourceService interface {
	Query(context.Context, *QueryDatasourceRequest) (*DatasourceSet, error)
}
