package rest

import (
	"net/http"

	"collab/pkg/collab"
)

// AppContext holds core components for running this web application i.e. router and user service.
type AppContext struct {
	router        http.Handler
	collabService collab.Servicer
}

// NewAppContext is a factory func to create a new AppContext.
func NewAppContext(ps collab.Servicer) *AppContext {
	app := &AppContext{collabService: ps}
	app.SetRoutes()
	return app
}

// GetRouter returns the app's router.
func (app *AppContext) GetRouter() http.Handler {
	return app.router
}
