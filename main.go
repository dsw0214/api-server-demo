package main

import (
	// "errors"
	// "fmt"
	"net/http"
	// "time"
	_ "net/http/pprof"
	"api-server-demo/config"
	"api-server-demo/src/model"
	"api-server-demo/src/router"
	"api-server-demo/src/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main()  {
	pflag.Parse()
	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	// Routes.
	router.Load(
		// Cores.
		g,
		// Middlwares.
		middleware.Logging(),
		middleware.RequestId(),
	)

	// ListenAndServe
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}