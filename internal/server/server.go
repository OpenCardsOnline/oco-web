package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/opencardsonline/oco-web/config"
	"github.com/opencardsonline/oco-web/internal/database"
	"github.com/opencardsonline/oco-web/internal/routers"
	logger "github.com/opencardsonline/oco-web/logging"
)

type Server struct {
	appConfig *config.AppConfig
	db        *database.AppDBConn
}

func (_s *Server) Start() {

	// For debugging start up time
	var startTime = time.Now()

	// Initialize App Configuration
	_s.appConfig = &config.AppConfig{}
	_s.appConfig.LoadEnvVars()

	// Initialize DB Connection
	db := &database.AppDBConn{}
	db.New(_s.appConfig.DBConnectionString)
	_s.db = db

	// Initialize Chi Routers
	r := routers.LoadRouters(_s.appConfig, _s.db)

	// DEBUG Start Up Time
	elapsed := time.Since(startTime)
	logger.Log.Info(fmt.Sprintf("Startup Time: [%s]", elapsed))

	// Start listening for requests
	http.ListenAndServe(":3000", r)

}
