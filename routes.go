package main
import(
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct{
	Name string
	Method string
	Pattern string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes{
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"product",
		"GET",
		"/product",
		showProduct,
	},
	Route{
		"productShow",
		"GET",
		"/product/{id}",
		showProduct,
	},
	Route{
		"productAdd",
		"POST",
		"/product",
		addProduct,
	},
	Route{
		"productRemove",
		"DELETE",
		"/product",
		removeProduct,
	},
	Route{
		"productUpdate",
		"PUT",
		"/product/{id}",
		updateProduct,
	},
	Route{
		"shoppingCartList",
		"GET",
		"/shoppging-cart",
		shoppingCartList,
	},
	Route{
		"shoppingCartAddProduct",
		"POST",
		"/shoppging-cart",
		addProductToShoppingCart,
	},
	Route{
		"shoppingCartRemove",
		"DELETE",
		"/shopping-cart/{id}",
		removeProductFromShoppingCart,
	},
	Route{
		"shoppingCartUpdate",
		"PUT",
		"/shopping-cart/{id}",
		shoppingCartUpdate,
	},

}