package main

import (
	"github.com/gerardcl/go-playing/master-server/routers"
	"github.com/gerardcl/go-playing/master-server/settings"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":3001", n)
}
