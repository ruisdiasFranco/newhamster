package newhamster
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
		index,
	},
	Route{
		"productList",
		"GET",
		"/product",
		productList,
	},
	Route{
		"prodcutShow",
		"GET",
		"/product/{id}",
		productShow,
	},
	Route{
		"prodcutAdd",
		"POST",
		"/product",
		productAdd,
	},
	Route{
		"prodcutRemove",
		"DELETE",
		"/product",
		productRemove,
	},
	Route{
		"prodcutUpdate",
		"PUT",
		"/product/{id}",
		productUpdate,
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
		shoppingCartAddProduct,
	},
	Route{
		"shoppingCartRemove",
		"DELETE",
		"/shopping-cart/{id}",
		shoppingCartAddProduct,
	},

}