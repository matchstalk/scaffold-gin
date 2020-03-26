package cmd

import (
	"fmt"
	"github.com/matchstalk/scaffold-gin/cmd/api"
	"github.com/matchstalk/scaffold-gin/cmd/migrate"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:               "scaffold-gin",
	Short:             "-v",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `scaffold-gin`,
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		usageStr := `scaffold-gin 0.0.1`
		fmt.Printf("%s\n", usageStr)
	},
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
