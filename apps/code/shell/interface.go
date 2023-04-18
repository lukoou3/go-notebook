package shell

import (
	"context"
	"go-notebook/apps/code"
)

const (
	ModelName = "code"
)

func registerName() string {
	return code.AppName + "-" + ModelName
}

type ShellCodeService interface {
	Query(context.Context, *QueryShellCodeRequest) (*ShellCodeSet, error)
}
