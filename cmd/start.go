package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-notebook/conf"
)

var (
	confFile string
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 demo 后端API",
	Long:  "启动 demo 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(confFile)
		// 加载程序配置
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			return err
		}

		fmt.Println("start and stop")
		return nil
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "conf", "f", "etc/demo.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
