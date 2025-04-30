package model

type Account struct {
	ID           string `json:"_id"`
	Email        string `json:"account_email"`
	Password     string `json:"account_password"`
	Type         string `json:"account_type"`
	Role         string `json:"account_role"`
	RestaurantID string `json:"account_restaurant_id"`
	EmployeeID   string `json:"account_employee_id"`
}