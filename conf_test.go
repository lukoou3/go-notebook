package main_test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/infraboard/mcube/sqlbuilder"
	"github.com/stretchr/testify/assert"
	"go-notebook/conf"
	"strings"
	"testing"
	"time"
)

func TestConvertTables(t *testing.T) {
	should := assert.New(t)
	var bytes []byte
	err := conf.LoadConfigFromToml("etc/demo.toml")
	if should.NoError(err) {
		db := conf.C().Sqlite.GetDB()
		rst, err := db.Query("SELECT name FROM sqlite_master WHERE type='table' and name like 'sql_%' and name != 'sqlite_sequence' ORDER BY name")
		if err != nil {
			panic(err)
		}
		var names []string
		for rst.Next() {
			var name string
			rst.Scan(&name)
			names = append(names, name)
		}
		rst.Close()

		for _, name := range names {
			rst, err = db.Query(fmt.Sprintf("select * from %s limit 1", name))
			if err != nil {
				panic(err)
			}

			columns, _ := rst.Columns()
			columnTypes, _ := rst.ColumnTypes()

			/*for i := range columns {
				fmt.Println(columns[i])
				fmt.Println(columnTypes[i].Name())
				fmt.Println(columnTypes[i].Nullable())
				fmt.Println(columnTypes[i].DatabaseTypeName())
				fmt.Println(columnTypes[i].ScanType())
			}*/

			fmt.Println(name)
			bytes, _ = json.Marshal(columns)
			typeNames := make([]string, len(columnTypes))
			for i := range columnTypes {
				typeNames[i] = columnTypes[i].DatabaseTypeName()
			}
			fmt.Println(string(bytes))
			bytes, _ = json.Marshal(typeNames)
			fmt.Println(string(bytes))
			fuzhis := make([]string, len(columns))
			for i := range columns {
				col := columns[i]
				var tp string
				typeName := strings.ToLower(columnTypes[i].DatabaseTypeName())
				if strings.HasPrefix(typeName, "int") {
					tp = "int"
				} else if strings.HasPrefix(typeName, "varchar") || typeName == "text" {
					tp = "string"
				} else if typeName == "datetime" {
					tp = "string"
				} else {
					panic("不支持的类型" + typeName)
				}
				fmt.Println(fmt.Sprintf("    %s %s  `json:\"%s\"`", strcase.ToCamel(col), tp, strcase.ToSnake(col)))
				fuzhis[i] = fmt.Sprintf("&%s.%s", "item", strcase.ToCamel(col))
			}
			fmt.Println(strings.Join(columns, ","))
			fmt.Println(strings.Join(fuzhis, ","))
			rst.Close()
			fmt.Println("------------------")
		}

	}
}

func TestGetTables(t *testing.T) {
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

func TestGetTableInfo(t *testing.T) {
	should := assert.New(t)
	var rst *sql.Rows
	var columns []string
	var columnTypes []*sql.ColumnType
	err := conf.LoadConfigFromToml("etc/demo.toml")
	if should.NoError(err) {
		db := conf.C().Sqlite.GetDB()
		rst, err = db.Query("SELECT type, name, tbl_name, sql FROM sqlite_master")
		fmt.Println(err)
		fmt.Println(rst)
		columns, err = rst.Columns()
		columnTypes, err = rst.ColumnTypes()
		for _, columnType := range columnTypes {
			fmt.Println(*columnType)
		}
		for rst.Next() {
			fmt.Println(columns)
			var tp string
			var name string
			var tbl_name string
			var sql string
			rst.Scan(&tp, &name, &tbl_name, &sql)
			fmt.Println(tp)
			fmt.Println(name)
			fmt.Println(tbl_name)
			fmt.Println(sql)
			fmt.Println("------------------")
		}
	}
}

type codeShellcode struct {
	Id         int32     `json:"id"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	Desc       string    `json:"desc"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func TestGetTableRows(t *testing.T) {
	should := assert.New(t)
	var rst *sql.Rows
	var columns []string
	var columnTypes []*sql.ColumnType
	err := conf.LoadConfigFromToml("etc/demo.toml")
	if should.NoError(err) {
		db := conf.C().Sqlite.GetDB()
		rst, err = db.Query("SELECT * FROM code_shellcode limit 2")
		fmt.Println(err)
		fmt.Println(rst)
		columns, err = rst.Columns()
		columnTypes, err = rst.ColumnTypes()
		for _, columnType := range columnTypes {
			fmt.Println(*columnType)
		}
		fmt.Println(columns)
		for rst.Next() {
			var id int32
			var name string
			var code string
			var desc string
			var create_time time.Time
			var update_time time.Time
			rst.Scan(&id, &name, &code, &desc, &create_time, &update_time)
			fmt.Println(id)
			fmt.Println(name)
			//fmt.Println(code)
			fmt.Println(desc)
			fmt.Println(create_time)
			fmt.Println(update_time)
			fmt.Println(create_time.Format("2006-01-02 15:04:05"))
			var data codeShellcode
			data.Id = id
			data.Name = name
			//data.Code = code
			data.Desc = desc
			data.CreateTime = create_time
			data.UpdateTime = update_time
			bytes, _ := json.Marshal(data)
			fmt.Println(create_time.Unix())
			fmt.Println(string(bytes))
			fmt.Println("------------------")
		}
	}
}

func TestGetTableRowsPage(t *testing.T) {
	should := assert.New(t)
	var rst *sql.Rows
	var columns []string
	var columnTypes []*sql.ColumnType
	err := conf.LoadConfigFromToml("etc/demo.toml")
	if should.NoError(err) {
		db := conf.C().Sqlite.GetDB()
		b := sqlbuilder.NewBuilder("select * from code_shellcode")
		b.Limit(10, 10)
		querySQL, args := b.Build()
		fmt.Println(querySQL)
		rst, err = db.Query(querySQL, args...)
		fmt.Println(err)
		fmt.Println(rst)
		columns, err = rst.Columns()
		columnTypes, err = rst.ColumnTypes()
		for _, columnType := range columnTypes {
			fmt.Println(*columnType)
		}
		fmt.Println(columns)
		for rst.Next() {
			var id int32
			var name string
			var code string
			var desc string
			var create_time time.Time
			var update_time time.Time
			rst.Scan(&id, &name, &code, &desc, &create_time, &update_time)
			fmt.Println(id)
			fmt.Println(name)
			//fmt.Println(code)
			fmt.Println(desc)
			fmt.Println(create_time)
			fmt.Println(update_time)
			fmt.Println(create_time.Format("2006-01-02 15:04:05"))
			var data codeShellcode
			data.Id = id
			data.Name = name
			//data.Code = code
			data.Desc = desc
			data.CreateTime = create_time
			data.UpdateTime = update_time
			bytes, _ := json.Marshal(data)
			fmt.Println(create_time.Unix())
			fmt.Println(string(bytes))
			fmt.Println("------------------")
		}
	}
}

func TestGetTableRows1(t *testing.T) {
	should := assert.New(t)
	var rst *sql.Rows
	var columns []string
	var columnTypes []*sql.ColumnType
	err := conf.LoadConfigFromToml("etc/demo.toml")
	if should.NoError(err) {
		db := conf.C().Sqlite.GetDB()
		rst, err = db.Query("SELECT * FROM code_shellcode limit 2")
		fmt.Println(err)
		fmt.Println(rst)
		columns, err = rst.Columns()
		columnTypes, err = rst.ColumnTypes()
		for _, columnType := range columnTypes {
			fmt.Println(*columnType)
		}
		fmt.Println(columns)
		for rst.Next() {
			var data codeShellcode
			//不行
			rst.Scan(&data)
			fmt.Println(data)
			fmt.Println("------------------")
		}
	}
}

func TestGetTableRows2(t *testing.T) {
	should := assert.New(t)
	var rst *sql.Rows
	var columns []string
	var columnTypes []*sql.ColumnType
	err := conf.LoadConfigFromToml("etc/demo.toml")
	if should.NoError(err) {
		db := conf.C().Sqlite.GetDB()
		rst, err = db.Query("SELECT * FROM code_shellcode limit 2")
		fmt.Println(err)
		fmt.Println(rst)
		columns, err = rst.Columns()
		columnTypes, err = rst.ColumnTypes()
		for _, columnType := range columnTypes {
			fmt.Println(*columnType)
		}
		fmt.Println(columns)
		for rst.Next() {
			var id int32
			var name string
			var code string
			var desc string
			// 2021-12-13T09:53:55.885222Z
			var create_time string
			var update_time string
			rst.Scan(&id, &name, &code, &desc, &create_time, &update_time)
			fmt.Println(id)
			fmt.Println(name)
			//fmt.Println(code)
			fmt.Println(desc)
			fmt.Println(create_time)
			fmt.Println(update_time)
			var t time.Time
			t, _ = time.Parse(time.RFC3339, create_time)
			fmt.Println(t)
			fmt.Println("------------------")
		}
	}
}
