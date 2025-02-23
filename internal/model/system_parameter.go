package model

type SaveSystemParameterDto struct {
	SysParaID          string `json:"sys_para_id" binding:"required,uuid4"`
	SysParaDescription string `json:"sys_para_description" binding:"required"`
	SysParaValue       string `json:"sys_para_value" binding:"required"`
}
