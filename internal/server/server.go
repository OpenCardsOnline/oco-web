package server

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/opencardsonline/oco-web/config"
	"github.com/opencardsonline/oco-web/internal/database"
	"github.com/opencardsonline/oco-web/internal/routers"
)

type Server struct {
	appConfig *config.AppConfig
	db        *pgx.Conn
}

func (_s *Server) Start() {

	// Initialize App Configuration
	_s.appConfig = &config.AppConfig{}
	_s.appConfig.LoadEnvVars()

	// Initialize DB Connection
	_s.db = database.InitializeDBConnection(_s.appConfig.DBConnectionString)

	// Initialize Chi Routers
	r := routers.LoadRouters(_s.appConfig, _s.db)

	// Start listening for requests
	http.ListenAndServe(":3000", r)

}
