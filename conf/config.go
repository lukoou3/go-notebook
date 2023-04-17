package conf

import "fmt"

// 全局config实例对象,
// 也就是我们程序，在内存中的配置对象
// 程序内部获取配置, 都通过读取该对象
// 该Config对象 什么时候被初始化喃?
//
//	配置加载时:
//	   LoadConfigFromToml
//	   LoadConfigFromEnv
//
// 为了不被程序在运行时恶意修改, 设置成私有变量
var config *Config

// 要想获取配置, 单独提供函数
// 全局Config对象获取函数
func C() *Config {
	return config
}

// 初始化一个有默认值的Config对象
func NewDefaultConfig() *Config {
	return &Config{
		App:    NewDefaultApp(),
		Log:    NewDefaultLog(),
		Sqlite: NewDefaultSqlite(),
	}
}

// Config 应用配置
// 通过封装为一个对象, 来与外部配置进行对接
type Config struct {
	App    *App    `toml:"app"`
	Log    *Log    `toml:"log"`
	Sqlite *Sqlite `toml:"sqlite"`
}

func NewDefaultApp() *App {
	return &App{
		Name: "demo",
		Host: "127.0.0.1",
		Port: "8050",
	}
}

type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
}

func (a *App) HttpAddr() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

func NewDefaultLog() *Log {
	return &Log{
		// debug, info, error, warn
		Level:  "info",
		Format: TextFormat,
		To:     ToStdout,
	}
}

// Log todo
// 用于配置全局Logger对象
type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
}
