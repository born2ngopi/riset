// Copyright (c) 2020-present BRIdge.
// See LICENSE.txt for license information.

package api1

import (
	"github.com/gorilla/mux"
)

type Routes struct {
	Root    *mux.Router
	ApiRoot *mux.Router // 	api/v1

	SuperUser       *mux.Router // api/v1/su/
	ManageAdmin     *mux.Router // api/v1/su/admin
	ManageAuthority *mux.Router // api/v1/su/authority

	Admin  *mux.Router //	api/v1/admin
	Viewer *mux.Router // 	role viewer

	Log          *mux.Router // api/v1/log
	ReportLog    *mux.Router // api/v1/log/report
	DashboardLog *mux.Router // api/v1/log/dashboard

	User *mux.Router // 	api/v1/user

	// merchant or platform
	Platform *mux.Router // api/v1/platform

	Auth *mux.Router // api/v1/auth

	Charge *mux.Router // api/v1/charge

}

type API struct {
	BaseRoutes *Routes
}

func Init(root *mux.Router) *API {
	api := &API{
		BaseRoutes: &Routes{},
	}

	api.BaseRoutes.Root = root
	api.BaseRoutes.ApiRoot = root.PathPrefix("/api/v2").Subrouter()

	api.BaseRoutes.SuperUser = api.BaseRoutes.ApiRoot.PathPrefix("/su").Subrouter()
	api.BaseRoutes.ManageAdmin = api.BaseRoutes.SuperUser.PathPrefix("/admin").Subrouter()
	api.BaseRoutes.ManageAuthority = api.BaseRoutes.SuperUser.PathPrefix("/authority").Subrouter()

	api.BaseRoutes.Admin = api.BaseRoutes.ApiRoot.PathPrefix("/admin").Subrouter()
	api.BaseRoutes.Viewer = api.BaseRoutes.ApiRoot.PathPrefix("/viewer").Subrouter()

	api.BaseRoutes.Log = api.BaseRoutes.ApiRoot.PathPrefix("/logs").Subrouter()
	api.BaseRoutes.ReportLog = api.BaseRoutes.Log.PathPrefix("/report").Subrouter()
	api.BaseRoutes.DashboardLog = api.BaseRoutes.Log.PathPrefix("/dashboard").Subrouter()

	api.BaseRoutes.User = api.BaseRoutes.ApiRoot.PathPrefix("/user").Subrouter()

	api.BaseRoutes.Platform = api.BaseRoutes.ApiRoot.PathPrefix("/platform").Subrouter()

	api.BaseRoutes.Auth = api.BaseRoutes.ApiRoot.PathPrefix("/auth").Subrouter()

	api.BaseRoutes.Charge = api.BaseRoutes.ApiRoot.PathPrefix("/charge").Subrouter()

	api.InitAdmin()
	api.InitSuperuser()

	return api
}
