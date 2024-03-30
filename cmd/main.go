package main

import (
	"os"
	"strconv"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/helpers"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	privKeyAccessTokenStr := os.Getenv("ACCESS_TOKEN_KEY")
	privKeyRefreshTokenStr := os.Getenv("REFRESH_TOKEN_KEY")
	accessTokenExpiration := os.Getenv("ACCESS_TOKEN_EXPIRATION")
	refreshTokenDuration := os.Getenv("REFRESH_TOKEN_DURATION")

	saltSize, _ := strconv.Atoi(os.Getenv("SALT_SIZE"))
	if saltSize < 1 {
		saltSize = 12 // set default value
	}

	secretKey := os.Getenv("HOLY_KEY") // jk, it is secret key for AES/GCM

	helpers.InitializeJWT(privKeyAccessTokenStr, privKeyRefreshTokenStr, accessTokenExpiration, refreshTokenDuration)

	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn:       dbDsn,
		SaltSize:  saltSize,
		SecretKey: secretKey,
	})

	opts := handler.NewServerOptions{
		Repository: repo,
	}
	return handler.NewServer(opts)
}
