package model

// CreateUserManagementAccountDto đại diện cho dữ liệu đầu vào khi tạo tài khoản
type CreateUserManagementAccountDto struct {
	UsId    string `json:"us_id" binding:"required,uuid4"`
	UsEmail string `json:"us_email" binding:"required,email"`
	UsPass  string `json:"us_pass" binding:"required,min=8,max=32"`
}

type UserManagementAccountOutput struct {
	UsId    string `json:"us_id"`
	UsEmail string `json:"us_email"`
}

type LoginUserManagementAccountDto struct {
	UsEmail string `json:"us_email" binding:"required,email"`
	UsPass  string `json:"us_pass" binding:"required,min=8,max=32"`
}

type LoginUserManagementAccountOutput struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ClientId     string `json:"client_id"`
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
