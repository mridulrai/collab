package rest

import (
	"fmt"
	"net/http"

	"collab/pkg/logger"
)

// indexHandler is a handler for the home page.
func (app *AppContext) indexHandler(w http.ResponseWriter, r *http.Request) {
	slog := logger.InitSugarLogger()
	defer slog.Sync()
	fmt.Fprintf(w, "Verloop.io - Collab Service")
	slog.Debug("Index Handler")
}
