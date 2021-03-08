package defaultapp

import "net/http"

type App struct{}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// all the incoming requests is handle by ServeHTTP first
	// so may we need to check every method and routes types
}
func New() *App {
	return &App{}
}
