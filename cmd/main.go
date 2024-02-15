package main

import (
	"github.com/irfanhanif/swtpro-intv/entity"
	"github.com/irfanhanif/swtpro-intv/service"
	"github.com/irfanhanif/swtpro-intv/utils"
	"os"

	"github.com/irfanhanif/swtpro-intv/generated"
	"github.com/irfanhanif/swtpro-intv/handler"
	"github.com/irfanhanif/swtpro-intv/repository"

	"github.com/labstack/echo/v4"

	postV1Users "github.com/irfanhanif/swtpro-intv/handler/post/v1/users"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	repo := repository.NewRepository(repository.NewRepositoryOptions{Dsn: dbDsn})

	uuid := utils.NewUUID()

	userEntityFactory := entity.NewUserFactory(uuid)

	userRegistrationService := service.NewUserRegistration(userEntityFactory, repo)

	postV1UsersHandler := postV1Users.NewHandler(userRegistrationService)

	return &handler.Server{
		PostV1UsersHandler: postV1UsersHandler,
	}
}
