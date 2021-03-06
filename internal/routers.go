package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/",
		Index,
	},

	Route{
		"GetGame",
		strings.ToUpper("Get"),
		"/api/games/{id}/",
		GetGame,
	},

	Route{
		"ListGames",
		strings.ToUpper("Get"),
		"/api/games/",
		ListGames,
	},

	Route{
		"NewGame",
		strings.ToUpper("Put"),
		"/api/games/",
		NewGame,
	},

	Route{
		"Play",
		strings.ToUpper("Post"),
		"/api/games/{id}/play",
		Play,
	},

	Route{
		"Authenticate",
		strings.ToUpper("Post"),
		"/api/login/",
		Authenticate,
	},

	Route{
		"RefreshAuthToken",
		strings.ToUpper("Post"),
		"/api/refresh/",
		RefreshAuthToken,
	},

	Route{
		"CreateUser",
		strings.ToUpper("Post"),
		"/api/users/",
		CreateUser,
	},
}
