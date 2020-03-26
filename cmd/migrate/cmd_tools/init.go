package cmdTools

import (
	"github.com/jinzhu/gorm"
	"github.com/matchstalk/scaffold-gin/common/gdb"
	"github.com/spf13/cobra"
)

type Tool struct {
	db *gorm.DB
}

func (e *Tool) getDb() *gorm.DB {
	return gdb.GetDB()
}

var (
	config string
	StartCmd = &cobra.Command{
		Use:     "tool",
		Short:   "Migrate about gdb operator tool",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/app.yml", "Start server with provided configuration file")
}

func run() {

}