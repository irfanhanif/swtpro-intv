package main

import (
	"github.com/irfanhanif/swtpro-intv/entity"
	"github.com/irfanhanif/swtpro-intv/service"
	"github.com/irfanhanif/swtpro-intv/utils"
	"io/ioutil"
	"os"

	"github.com/irfanhanif/swtpro-intv/generated"
	"github.com/irfanhanif/swtpro-intv/handler"
	"github.com/irfanhanif/swtpro-intv/repository"

	"github.com/labstack/echo/v4"

	getV1Users "github.com/irfanhanif/swtpro-intv/handler/get/v1/users"
	postV1Token "github.com/irfanhanif/swtpro-intv/handler/post/v1/token"
	postV1Users "github.com/irfanhanif/swtpro-intv/handler/post/v1/users"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	jwtPrivateKey, err := ioutil.ReadFile("cert/jwtRS256.key")
	if err != nil {
		panic(err)
	}

	jwtPublicKey, err := ioutil.ReadFile("cert/jwtRS256.key.pub")
	if err != nil {
		panic(err)
	}

	dbDsn := os.Getenv("DATABASE_URL")

	repo := repository.NewRepository(repository.NewRepositoryOptions{Dsn: dbDsn})

	uuid := utils.NewUUID()
	jwt := utils.NewJWT(jwtPrivateKey, jwtPublicKey)

	userEntityFactory := entity.NewUserFactory(uuid)

	userRegistrationService := service.NewUserRegistration(userEntityFactory, repo)
	tokenGeneratorService := service.NewTokenGenerator(repo, repo, jwt)
	userGetterService := service.NewUserGetter(repo)

	postV1UsersHandler := postV1Users.NewHandler(userRegistrationService)
	postV1TokenHandler := postV1Token.NewHandler(tokenGeneratorService)
	getV1Users := handler.NewAuthenticatorMiddleware(jwt, getV1Users.NewHandler(userGetterService))

	return &handler.Server{
		PostV1UsersHandler: postV1UsersHandler,
		PostV1TokenHandler: postV1TokenHandler,
		GetV1UsersHandler:  getV1Users,
	}
}
