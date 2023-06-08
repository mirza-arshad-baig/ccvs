package main

import (
	"ccvs/common/libs"

	app "ccvs/cmd/init"
	"ccvs/common/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	// ===========================================================================
	// Load Environment Config
	// ===========================================================================
	libs.ConfigFile = "app"
	// ConfigPaths is config file paths
	libs.InitConfig("./config")

	libs.ConfigFile = "banned_countries"
	libs.InitConfig("./config")
	viper.GetStringSlice("banned_countries")
}

func main() {
	router := gin.New()

	// ===========================================================================
	// Load Default Middlewares
	// ===========================================================================
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(middleware.CORSMiddleware())

	// ===========================================================================
	// Load Dependencies
	// ===========================================================================
	app.InitializeDependencies()

	// ===========================================================================
	// Load Routes
	// ===========================================================================
	app.InitializeRoutes(router)

	hostname := viper.GetString("app.address")
	log.Infoln("Serving at", hostname)
	srv := &http.Server{
		ReadHeaderTimeout: 1 * time.Minute,
		Addr:              hostname,
		Handler:           router,
	}
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
