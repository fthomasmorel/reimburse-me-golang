package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route type is used to define a route of the API
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes type is an array of Route
type Routes []Route

// NewRouter is the constructeur of the Router
// It will create every routes from the routes variable just above
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	//Debt
	Route{"GetDebt", "GET", "/debt/{id}", GetDebtController},
	Route{"DeleteDebt", "DELETE", "/debt/{id}", DeleteDebtController},
	Route{"AddImageDebt", "POST", "/debt/{id}/image", AddImageDebtController},
	Route{"ReimburseDebt", "PUT", "/debt/{id}", ReimburseDebtController},

	//User + Debt
	Route{"GetMyDebts", "GET", "/user/{userID}/mydebt", GetMyDebtsController},
	Route{"GetTheirDebts", "GET", "/user/{userID}/debt", GetTheirDebtsController},
	Route{"CreateDebt", "POST", "/user/{userID}/debt", CreateDebtController},

	//User
	Route{"LogUser", "GET", "/user/{id}/login", LogUserController},
	Route{"GetUser", "GET", "/user/{id}", GetUserController},
	Route{"CreateUser", "POST", "/user", CreateUserController},
	Route{"DeleteUser", "DELETE", "/user/{id}", DeleteUserController},
	Route{"AddPayee", "POST", "/user/{id}/payee/{payeeID}", AddPayeeController},
	Route{"RemovePayee", "DELETE", "/user/{id}/payee/{payeeID}", RemovePayeeController},
}
