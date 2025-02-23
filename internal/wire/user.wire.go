//go:build wireinject

package wire

import (
	"system-management-pg/internal/controller"
	"system-management-pg/internal/repo"
	"system-management-pg/internal/service"
	"github.com/google/wire"
)

func InitUserRouterHanlder() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
