package main

import (
	"github.com/emicklei/go-restful"
	"io"
	"log"
	"net/http"
)

// This example shows how to have a program with 2 WebServices containers
// each having a http server listening on its own port.
//
// The first "hello" is added to the restful.DefaultContainer (and uses DefaultServeMux)
// For the second "hello", a new container and ServeMux is created
// and requires a new http.Server with the container being the Handler.
// This first server is spawn in its own go-routine such that the program proceeds to create the second.
//
// GET http://localhost:8080/hello
// GET http://localhost:8081/hello

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/{film-id}").To(hello))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func hello(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("film-id")
	magnet := Id2Magnet(id)
	io.WriteString(resp, magnet)
}
