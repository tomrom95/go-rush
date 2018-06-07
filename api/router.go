package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tomrom95/go-rush/api/handlers"
	"github.com/tomrom95/go-rush/api/logging"
	"github.com/tomrom95/go-rush/api/routes"
)

func NewRouter(context *handlers.Context) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes.GetRoutes() {
		for method, handler := range route.Methods {
			fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handler(context, w, r)
			})
			router.
				Methods(method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(logging.WrapRoute(fn, route.Name))
		}
	}
	return router
}
