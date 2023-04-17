package conf

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"sync"
	"time"
)

// 全局MySQL 客户端实例
var sqliteDb *sql.DB

func NewDefaultSqlite() *Sqlite {
	return &Sqlite{
		File:        "datas/base_udf_data",
		MaxOpenConn: 20,
		MaxIdleConn: 10,
		MaxLifeTime: 1000 * 60 * 60 * 6,
		MaxIdleTime: 1000 * 60 * 60,
	}
}

type Sqlite struct {
	File string `toml:"file" env:"SQLITE_FILE"`
	// 控制当前程序的MySQL打开的连接数
	MaxOpenConn int `toml:"max_open_conn" env:"SQLITE_MAX_OPEN_CONN"`
	// 控制MySQL复用, 比如5, 最多运行5个来复用
	MaxIdleConn int `toml:"max_idle_conn" env:"SQLITE_MAX_IDLE_CONN"`
	// 一个连接的生命周期, 这个和MySQL Server配置有关系, 必须小于Server配置
	// 一个连接用12h 换一个conn, 保证一定的可用性
	MaxLifeTime int `toml:"max_life_time" env:"SQLITE_MAX_LIFE_TIME"`
	// Idle 连接 最多允许存活多久
	MaxIdleTime int `toml:"max_idle_time" env:"SQLITE_MAX_idle_TIME"`

	// 作为私有变量, 用户与控制GetDB
	lock sync.Mutex
}

// 1. 第一种方式, 使用LoadGlobal 在加载时 初始化全局db实例
// 2. 第二种方式, 惰性加载, 获取DB是，动态判断再初始化
func (m *Sqlite) GetDB() *sql.DB {
	// 直接加锁, 锁住临界区
	m.lock.Lock()
	defer m.lock.Unlock()

	// 如果实例不存在, 就初始化一个新的实例
	if sqliteDb == nil {
		conn, err := m.getDBConn()
		if err != nil {
			panic(err)
		}
		sqliteDb = conn
	}

	// 全局变量db就一定存在了
	return sqliteDb
}

// 连接池, driverConn具体的连接对象, 他维护着一个Socket
// pool []*driverConn, 维护pool里面的连接都是可用的, 定期检查我们的conn健康情况
// 某一个driverConn已经失效, driverConn.Reset(), 清空该结构体的数据, Reconn获取一个连接, 让该conn借壳存活
// 避免driverConn结构体的内存申请和释放的一个成本
func (m *Sqlite) getDBConn() (*sql.DB, error) {
	var err error
	db, err := sql.Open("sqlite3", m.File)
	if err != nil {
		return nil, fmt.Errorf("connect to sqlite3(%s) error, %s", m.File, err.Error())
	}

	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping sqlite3<%s> error, %s", m.File, err.Error())
	}
	return db, nil
}
