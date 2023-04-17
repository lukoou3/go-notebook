package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-notebook/conf"
	"testing"
)

func TestGetDB(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("etc/demo.toml")
	if should.NoError(err) {
		db := conf.C().Sqlite.GetDB()
		rst, err := db.Query("SELECT name FROM sqlite_master WHERE type='table' ORDER BY name")
		fmt.Println(err)
		fmt.Println(rst)
		for rst.Next() {
			var name string
			rst.Scan(&name)
			fmt.Println(name)
		}
	}
}
