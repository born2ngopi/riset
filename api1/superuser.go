package api1

import (
	"encoding/json"
	"net/http"
)

func (api *API) InitSuperuser() {

	api.BaseRoutes.SuperUser.Handle("/", api.ApiAuthorizeRequeired(GetProfileSuperuser)).Methods(http.MethodGet)
	api.BaseRoutes.SuperUser.Handle("/", api.ApiAuthorizeRequeired(UpdateProfileSuperuser)).Methods(http.MethodPut)

	api.BaseRoutes.ManageAdmin.Handle("/", api.ApiAuthorizeRequeired(FetchAllAdmins)).Methods(http.MethodGet)
	api.BaseRoutes.ManageAdmin.Handle("/{admin_id:[0-9]+}/detail", api.ApiAuthorizeRequeired(FetchAdmin)).Methods(http.MethodGet)
	api.BaseRoutes.ManageAdmin.Handle("/", api.ApiAuthorizeRequeired(UpdateAdmin)).Methods(http.MethodPost)

	api.BaseRoutes.ManageAuthority.Handle("/{admin_id:[0-9]}", api.ApiAuthorizeRequeired(GetAdminAuthority)).Methods(http.MethodGet)
	api.BaseRoutes.ManageAuthority.Handle("/{admin_id:[0-9]}", api.ApiAuthorizeRequeired(UpdateAdminAuthority)).Methods(http.MethodPut)
}

func GetProfileSuperuser(w http.ResponseWriter, r *http.Request) {

	resp := struct {
		Status string `json:"status"`
		Data   string `json:"data"`
	}{
		"success",
		"mantab",
	}
	json.NewEncoder(w).Encode(resp)
}

func UpdateProfileSuperuser(w http.ResponseWriter, r *http.Request) {

}

func FetchAllAdmins(w http.ResponseWriter, r *http.Request) {

}

func FetchAdmin(w http.ResponseWriter, r *http.Request) {

}

func UpdateAdmin(w http.ResponseWriter, r *http.Request) {

}

func GetAdminAuthority(w http.ResponseWriter, r *http.Request) {

}

func UpdateAdminAuthority(w http.ResponseWriter, r *http.Request) {

}
