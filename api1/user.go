package api1

import "net/http"

func (api *API) InitUser() {

	api.BaseRoutes.User.Handle("/", api.ApiAuthorizeRequeired(GetUserProfile)).Methods(http.MethodGet)
	api.BaseRoutes.User.Handle("/", api.ApiAuthorizeRequeired(UpdateUserProfile)).Methods(http.MethodPost)
	api.BaseRoutes.User.Handle("/password/change", api.ApiAuthorizeRequeired(UpdateUserPassword)).Methods(http.MethodPut)
	api.BaseRoutes.User.Handle("/{token}/confirm/password", api.ApiAuthorizeRequeired(ConfirmUserPassword)).Methods(http.MethodGet)

	api.BaseRoutes.User.Handle("/signup", api.ApiHandler(RegisterUser)).Methods(http.MethodPost)
	api.BaseRoutes.User.Handle("/{token}/activated", api.ApiHandler(ActivatedUser)).Methods(http.MethodGet)
	api.BaseRoutes.User.Handle("/login", api.ApiHandler(LoginUser)).Methods(http.MethodPost)

}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {

}

func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {

}

func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {

}

func ConfirmUserPassword(w http.ResponseWriter, r *http.Request) {

}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

}

func ActivatedUser(w http.ResponseWriter, r *http.Request) {

}

func LoginUser(w http.ResponseWriter, r *http.Request) {

}
