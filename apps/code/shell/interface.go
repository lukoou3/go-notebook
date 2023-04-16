package shell

import (
	"context"
)

const (
	ModelName = "code"
)

type ShellCodeService interface {
	Query(context.Context, *QueryShellCodeRequest) (*ShellCodeSet, error)
}
