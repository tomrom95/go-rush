package routes

import (
	"github.com/tomrom95/go-rush/api/handlers"
)

type Route struct {
	Name    string
	Pattern string
	Methods map[string]handlers.HandlerWithContext
}

type Routes []Route

func GetRoutes() []*Route {
	return []*Route{
		&Route{
			Name:    "Rushees",
			Pattern: "/rushees",
			Methods: map[string]handlers.HandlerWithContext{
				"GET": handlers.GetRushees,
			},
		},
	}
}
