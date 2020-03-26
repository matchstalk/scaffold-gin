package apply

import (
	"github.com/matchstalk/scaffold-gin/app/dao"
	"github.com/matchstalk/scaffold-gin/cmd/migrate/migration"
	"github.com/matchstalk/scaffold-gin/tools"
	"github.com/spf13/cobra"
	"reflect"
)

const migrationPath = "cmd/migrate/migration"

var (
	config string
	name string
	StartCmd = &cobra.Command{
		Use:     "apply",
		Short:   "Migrate apply about gdb operator",
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/app.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Migrate operator name")
}

func setup() {
	//1. 读取配置
	tools.ConfigSetup(config)
	//2. 设置日志
	tools.InitLogger()
	//3. 初始化数据库链接
	dao.Setup()
}

func run() {
	m := &migration.Migration{}
	val := reflect.ValueOf(*m)

	if len(name) == 0 {
		for i := 0; i < val.NumMethod(); i++ {
			val.Method(i).Call(nil)
		}
	}
}