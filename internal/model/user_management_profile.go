package model

type UserManagementProfileOutput struct {
	UsEmail    string `json:"us_email"`
	UsID       string `json:"us_id"`
	UsName     string `json:"us_name"`
	UsAvatar   string `json:"us_avatar"`
	UsPhone    string `json:"us_phone"`
	UsGender   string `json:"us_gender"`
	UsAddress  string `json:"us_address"`
	UsBirthday string `json:"us_birthday"`
	Isdeleted  int32  `json:"isDeleted"`
}

type CreateUserManagementProfileDto struct {
	UsName     string `json:"us_name" binding:"required"`
	UsAvatar   string `json:"us_avatar" binding:"required"`
	UsPhone    string `json:"us_phone" binding:"required"`
	UsGender   string `json:"us_gender" binding:"required"`
	UsAddress  string `json:"us_address" binding:"required"`
	UsBirthday string `json:"us_birthday" binding:"required"`
}

type UpdateUserManagementProfileDto struct {
	UsID       string `json:"us_id" binding:"required"`
	UsName     string `json:"us_name" binding:"required"`
	UsAvatar   string `json:"us_avatar" binding:"required"`
	UsPhone    string `json:"us_phone" binding:"required"`
	UsGender   string `json:"us_gender" binding:"required"`
	UsAddress  string `json:"us_address" binding:"required"`
	UsBirthday string `json:"us_birthday" binding:"required"`
}
