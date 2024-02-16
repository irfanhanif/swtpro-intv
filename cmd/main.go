package main

import (
	"fmt"
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
	patchV1Users "github.com/irfanhanif/swtpro-intv/handler/patch/v1/users"
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
	jwtPrivateKey, err := ioutil.ReadFile("./cert/jwtRS256.key")
	if err != nil {
		fmt.Println("jwt private not found in local")
	}

	jwtPublicKey, err := ioutil.ReadFile("./cert/jwtRS256.key.pub")
	if err != nil {
		fmt.Println("jwt public key not found in local")
	}

	jwtPrivateKey, err = ioutil.ReadFile("/run/secrets/jwtPrivateKey")
	if err != nil {
		panic(err)
	}

	jwtPublicKey, err = ioutil.ReadFile("/run/secrets/jwtPublicKey")
	if err != nil {
		panic(err)
	}

	if len(jwtPrivateKey) < 1 || len(jwtPublicKey) < 1 {
		panic("failed to load jwt keys")
	}

	dbDsn := os.Getenv("DATABASE_URL")

	repo := repository.NewRepository(repository.NewRepositoryOptions{Dsn: dbDsn})

	uuid := utils.NewUUID()
	jwt := utils.NewJWT(jwtPrivateKey, jwtPublicKey)

	userEntityFactory := entity.NewUserFactory(uuid)

	userRegistrationService := service.NewUserRegistration(userEntityFactory, repo)
	tokenGeneratorService := service.NewTokenGenerator(repo, repo, jwt)
	userGetterService := service.NewUserGetter(repo)
	userUpdaterService := service.NewUserUpdater(repo)

	postV1UsersHandler := postV1Users.NewHandler(userRegistrationService)
	postV1TokenHandler := postV1Token.NewHandler(tokenGeneratorService)
	getV1Users := handler.NewAuthenticatorMiddleware(jwt, getV1Users.NewHandler(userGetterService))
	patchV1Users := handler.NewAuthenticatorMiddleware(jwt, patchV1Users.NewHandler(userUpdaterService))

	return &handler.Server{
		PostV1UsersHandler:  postV1UsersHandler,
		PostV1TokenHandler:  postV1TokenHandler,
		GetV1UsersHandler:   getV1Users,
		PatchV1UsersHandler: patchV1Users,
	}
}
