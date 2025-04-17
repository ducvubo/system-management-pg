package model

type RegisterUserPatoAccountDto struct {
	UsPtEmail   string `json:"us_pt_email" binding:"required,email"`
	UsPtPass    string `json:"us_pt_pass" binding:"required,min=8,max=32"`
	UrlRedirect string `json:"url_redirect" binding:"required,url"`
}

type RegisterUserPatoAccountOutput struct {
	UsPtEmail string `json:"us_pt_email"`
}

type ResendOtpDto struct {
	UsPtEmail   string `json:"us_pt_email" binding:"required,email"`
	UrlRedirect string `json:"url_redirect" binding:"required,url"`
}

type ActiveAccountDto struct {
	UsPtEmail string `json:"us_pt_email" binding:"required,email"`
	UsPtOtp   int    `json:"us_pt_otp" binding:"required"`
}

type LoginUserPatoAccountDto struct {
	UsPtEmail string `json:"us_pt_email" binding:"required,email"`
	UsPtPass  string `json:"us_pt_pass" binding:"required,min=8,max=32"`
}

// trả về cặp token
type LoginUserPatoAccountOutput struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
