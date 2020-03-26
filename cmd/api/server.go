package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/matchstalk/scaffold-gin/app/dao"
	"github.com/matchstalk/scaffold-gin/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config string
	port   string
	//auto tls
	//StartCmd : set up restful api server
	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start API server",
		Example: "scaffold-gin server config/app.yaml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/app.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "9098", "Tcp port server listening on")
}

func usage() {
	usageStr := `starting api server`
	fmt.Printf("%s\n", usageStr)
}

func setup() {
	//1. 读取配置
	tools.ConfigSetup(config)
	//2. 设置日志
	tools.InitLogger()
	//3. 初始化数据库链接
	dao.Setup()
	//4. 设置gin mode
	if viper.GetString("app.runmode") == string(tools.RunModeProd) {
		gin.SetMode(gin.ReleaseMode)
	}
}

func run() error {
	return nil
}