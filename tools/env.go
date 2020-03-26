/**
 * 环境相关枚举类型
 */
package tools

type (
	RunMode string
)

const (
	RunModeDev  RunMode = "dev"  //开发模式
	RunModeTest RunMode = "test" //测试模式
	RunModeProd RunMode = "prod" //生产模式
)
