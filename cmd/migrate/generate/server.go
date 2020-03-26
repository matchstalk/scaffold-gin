package generate

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"github.com/matchstalk/scaffold-gin/app/dao"
	"github.com/matchstalk/scaffold-gin/tools"
	"github.com/matchstalk/scaffold-gin/tools/process"
	"github.com/spf13/cobra"
	"time"
)

const (
	migrationPath = "cmd/migrate/migration"
	toolsPath = "cmd/migrate/cmd_tools"
)

var (
	config string
	name string
	tool bool
	StartCmd = &cobra.Command{
		Use:     "generate",
		Short:   "Migrate generate about gdb operator",
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().BoolVarP(&tool, "tool", "t", false, "是否为工具")
	StartCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "操作名")
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/app.yml", "需要载入的配置文件")
}

func setup() {
	//1. 读取配置
	tools.ConfigSetup(config)
	//2. 设置日志
	tools.InitLogger()
	//3. 初始化数据库链接
	dao.Setup()
	//4. 验证数据
	if len(name) == 0 {
		panic("请输入操作名称")
	}
}

func run() {
	if tool {
		generateTool()
	} else {
		generateMigration()
	}
}

func generateTool() {
	method := fmt.Sprintf("%s_tool", name)
	f := NewFile("cmdTools")
	f.Var().Params(
		Id(name + "Cmd").Op("=").Op("&").Qual("github.com/spf13/cobra", "Command").Values(Dict{
			Id("Use"): Lit(name + "Cmd"),
			Id("Short"): Lit("Migrate about gdb operator tool " + name),
			Id("Run"): Id(name),
		}),
	)
	f.Func().Id("init").Params().Block(
		Id("StartCmd").Dot("AddCommand").Params(Id(name + "Cmd")),
	)
	f.Func().Id(name).Params(
		Id("cmd").Op("*").Qual("github.com/spf13/cobra", "Command"),
		Id("args").Op("[]").String(),
	).Block()
	fmt.Println(method)
	fmt.Printf("%#v", f)
	err := f.Save(toolsPath + "/" + method + ".go")
	if err != nil {
		panic(err)
	}
}

func generateMigration()  {
	method := fmt.Sprintf("m_%d_%s_table", time.Now().Unix(), name)
	f := NewFile("migration")
	f.Func().
		Params(Id("e").Id("Migration")).
		Id(process.Marshal(method)).
		Params().Block(
		Id("version").Op(":=").Lit(method),
		Var().Id("err").Error(),
		Id("e").Dot("db").Op("=").Id("e").Dot("getDb").Call().Dot("Begin").Call(),
		Id("tx").Op(":=").Id("e").Dot("db"),
		Id("defer").Func().Call().Block(
			If(Id("err").Op("!=").Nil().Block(
				Id("tx").Dot("Rollback").Call(),
				Qual("fmt", "Printf").Params(Lit("%s run failed!"), Id("version")),
			),
			).Else().Block(
				Id("tx").Dot("Commit").Call(),
				Qual("fmt", "Printf").Params(Lit("%s run finished!"), Id("version")),
			),
		).Call(),
		Id("err").Op("=").Id("e").Dot("AddMigrateVersion").Params(Id("version")),
		If(
			Id("err").Op("!=").Nil().Block(Return()),
		),
		Comment("todo migration block"),
		Comment(""),
		Comment("todo migration block"),
	)

	err := f.Save(migrationPath + "/" + method + ".go")
	if err != nil {
		panic(err)
	}
}