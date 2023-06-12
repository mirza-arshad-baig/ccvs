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

	// ===========================================================================
	// Load Banned Countries List Config
	// ===========================================================================
	libs.ConfigFile = "banned_countries"
	libs.InitConfig("./config")
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

	// Get the hostname from the app configuration
	hostname := viper.GetString("app.address")
	log.Infoln("Serving at", hostname)

	// Create an HTTP server with custom configurations
	srv := &http.Server{
		ReadHeaderTimeout: 1 * time.Minute,
		Addr:              hostname,
		Handler:           router,
	}

	// Start listening and serving requests
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
