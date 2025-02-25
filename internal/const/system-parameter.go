package consts

type SystemParameter string

const (
	SystemEmail  SystemParameter = "cd0c100f-a0af-406f-a94a-79bc80ea7f98"
	SystemMobile SystemParameter = "56ca0dd0-557c-4c43-a1a1-dc1f93defcea"
	BannerHeader SystemParameter = "79170c98-15b1-4934-8ced-13f64d32de58"
)

func (s SystemParameter) String() string {
	return string(s)
}

var SystemParameterMap = map[string]SystemParameter{
	string(SystemEmail):  SystemEmail,
	string(SystemMobile): SystemMobile,
}

// Hàm kiểm tra ID có hợp lệ không và trả về loại SystemParameter
func GetSystemParameter(value string) (SystemParameter, bool) {
	param, exists := SystemParameterMap[value]
	return param, exists
}
