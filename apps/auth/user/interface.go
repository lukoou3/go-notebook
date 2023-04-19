package user

import (
	"go-notebook/apps/auth"
)

const (
	ModelName = "user"
)

func registerName() string {
	return auth.AppName + "-" + ModelName
}
