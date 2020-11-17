package api1

import "net/http"

func (api *API) ApiHandler(f func(http.ResponseWriter, *http.Request)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next := http.HandlerFunc(f)
		next.ServeHTTP(w, r)
	})

}

func (api *API) ApiAuthorizeRequeired(f func(http.ResponseWriter, *http.Request)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next := http.HandlerFunc(f)
		next.ServeHTTP(w, r)
	})

}
