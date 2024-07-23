package server

import (
	"net/http"

	"github.com/opencardsonline/oco-web/config"
	"github.com/opencardsonline/oco-web/internal/database"
	"github.com/opencardsonline/oco-web/internal/routers"
)

func RunServer() {

	config := config.LoadEnvVars()

	database.InitializeDBConnection(config.DBConnectionString)

	r := routers.LoadRouters()

	http.ListenAndServe(":3000", r)

}
