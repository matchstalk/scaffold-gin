package migrate

import (
	"fmt"
	"github.com/matchstalk/scaffold-gin/app/dao"
	"github.com/matchstalk/scaffold-gin/app/models"
	"github.com/matchstalk/scaffold-gin/cmd/migrate/apply"
	cmdTools "github.com/matchstalk/scaffold-gin/cmd/migrate/cmd_tools"
	"github.com/matchstalk/scaffold-gin/cmd/migrate/generate"
	"github.com/matchstalk/scaffold-gin/common/gdb"
	"github.com/matchstalk/scaffold-gin/common/migrate"
	"github.com/matchstalk/scaffold-gin/tools"
	"github.com/spf13/cobra"
)

var (
	config   string
	StartCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate about gdb operator",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.AddCommand(apply.StartCmd)
	StartCmd.AddCommand(generate.StartCmd)
	StartCmd.AddCommand(cmdTools.StartCmd)
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/app.yml", "Start server with provided configuration file")
}

func run() {
	usage := `start migrate`
	fmt.Println(usage)
	//1. 读取配置
	tools.ConfigSetup(config)
	//2. 设置日志
	tools.InitLogger()
	//3. 初始化数据库链接
	dao.Setup()
	//4. 数据库迁移
	db := gdb.GetDB()
	db.AutoMigrate(migrateModel()...)
	usage = `finish`
	fmt.Println(usage)
}

func migrateModel() []interface{} {
	return []interface{}{
		&migrate.Migrate{},
		&models.User{},
		&models.Application{},
	}
}
