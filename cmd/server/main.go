package main

import (
	"fmt"
	_ "system-management-pg/cmd/swag/docs"
	"system-management-pg/internal/initialize"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Documentation System Management
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  https://github.com/ducvubo

// @contact.name   Vũ Đức Bo
// @contact.url    https://github.com/ducvubo
// @contact.email  vminhduc8@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:13000
// @BasePath  /api/v1
// @schema https

// @securityDefinitions.apikey  ClientIDAuth
// @in                          header
// @name                        id_user_guest

// @securityDefinitions.apikey  AccessTokenAuth
// @in                          header
// @name                        x-at-tk

// @securityDefinitions.apikey  RefreshTokenAuth
// @in                          header
// @name                        x-rf-tk

// @Security ClientIDAuth
// @Security AccessTokenAuth
// @Security RefreshTokenAuth
func main() {
	fmt.Println("Đây là một log thông thường heheh đây là hot reload")
	fmt.Printf("Log có format: %s\n", "Hello, Go!")
	r := initialize.Run()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":13000")

}
