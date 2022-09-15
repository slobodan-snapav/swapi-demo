/*
 * SWAPI-DEV
 *
 * defaultDescription
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

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
		"/",
		Index,
	},

	Route{
		"ApiFilmsGet",
		strings.ToUpper("Get"),
		"/api/films",
		ApiFilmsGet,
	},

	Route{
		"ApiGet",
		strings.ToUpper("Get"),
		"/api/",
		ApiGet,
	},

	Route{
		"ApiPeopleGet",
		strings.ToUpper("Get"),
		"/api/people",
		ApiPeopleGet,
	},

	Route{
		"ApiPlanetsGet",
		strings.ToUpper("Get"),
		"/api/planets",
		ApiPlanetsGet,
	},

	Route{
		"ApiSpeciesGet",
		strings.ToUpper("Get"),
		"/api/species",
		ApiSpeciesGet,
	},

	Route{
		"ApiStarshipsGet",
		strings.ToUpper("Get"),
		"/api/starships",
		ApiStarshipsGet,
	},

	Route{
		"ApiVehiclesGet",
		strings.ToUpper("Get"),
		"/api/vehicles",
		ApiVehiclesGet,
	},
}
