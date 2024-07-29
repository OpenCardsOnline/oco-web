package routers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/opencardsonline/oco-web/config"
	"github.com/opencardsonline/oco-web/internal/database"
	"github.com/opencardsonline/oco-web/internal/routers/handlers"
)

func LoadRouters(config *config.AppConfig, db *database.AppDBConn) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Setup File Server
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "public"))
	FileServer(r, "/public", filesDir)

	apiHandlers := &handlers.APIHandlers{}
	apiHandlers.InitializeAPIHandlers(config, db)

	// Healthcheck
	r.Get("/health", apiHandlers.GetHealthCheck)

	// API Handlers
	// V1
	r.Post("/api/v1/auth/register", apiHandlers.AuthRegisterNewUser)
	r.Get("/api/v1/auth/verify", apiHandlers.AuthVerifyNewUser)

	// UI Handlers
	handlers.ParseTemplates()
	r.Get("/", handlers.ComingSoonPageHandler)
	r.Get("/home", handlers.HomePageHandler)
	r.Get("/*", handlers.PageNotFoundHandler)

	return r
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
