package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/bridge/api1"
)

func main() {
	fmt.Println("server running")

	r := mux.NewRouter()
	api := api1.Init(r)

	fmt.Println("server runnig")
	http.ListenAndServe(":8000", api.BaseRoutes.Root)

}
