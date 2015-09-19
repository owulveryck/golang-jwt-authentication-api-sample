package main

import (
	"github.com/owulveryck/golang-jwt-authentication-api-sample/routers"
	"github.com/owulveryck/golang-jwt-authentication-api-sample/settings"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
