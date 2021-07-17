package rest

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// SetRoutes sets up routes
func (app *AppContext) SetRoutes() {
	commonHandlers := alice.New(context.ClearHandler, identifierHandler, loggingHandler, recoverHandler)

	r := mux.NewRouter()
	r.Handle("/", commonHandlers.ThenFunc(app.indexHandler))

	app.router = r
}
