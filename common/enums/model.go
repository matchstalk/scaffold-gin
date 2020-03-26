/**
 * 模型相关枚举类型
 */
package enums

type (
	//性别
	Sex uint8
	//状态
	Status uint8
)

const (
	_       Sex    = iota
	Male           //男
	Female         //女
	_       Status = iota
	Enable         //启用
	Disable        //禁用
)
