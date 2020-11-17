package api1

import (
	"net/http"
)

func (api *API) InitAdmin() {

	api.BaseRoutes.Admin.Handle("/", api.ApiAuthorizeRequeired(GetProfileAdmin)).Methods(http.MethodGet)
	api.BaseRoutes.Admin.Handle("/", api.ApiAuthorizeRequeired(UpdateProfileAdmin)).Methods(http.MethodPut)
	api.BaseRoutes.Admin.Handle("/password", api.ApiAuthorizeRequeired(UpdatePasswordAdmin)).Methods(http.MethodPut)
}

func GetProfileAdmin(w http.ResponseWriter, r *http.Request) {

}

func UpdateProfileAdmin(w http.ResponseWriter, r *http.Request) {

}

func UpdatePasswordAdmin(w http.ResponseWriter, r *http.Request) {

}
