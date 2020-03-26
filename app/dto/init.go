package dto

type GeneralDelDto struct {
	Id string `uri:"id" json:"id" binding:"required"`
}
type GeneralGetDto struct {
	Id string `uri:"id" json:"id" binding:"required"`
}
