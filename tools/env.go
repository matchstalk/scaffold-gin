/**
 * 环境相关枚举类型
 */
package tools

type (
	RunMode string
)

const (
	RunModeDev  RunMode = "dev"      //开发模式
	RunModeTest RunMode = "test"     //测试模式
	RunModeProd RunMode = "prod"     //生产模式
	Mysql               = "mysql"    //mysql数据库标识
	Postgres            = "postgres" //postgres数据库标识
)
