package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

type User struct {
	Id, Name string
}

func main()  {
	service := new(restful.WebService)
	service.
		Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	service.Route(service.GET("/{user-id}").To(FindUser))
	restful.Add(service)
	fmt.Println("start web-demo listen 9000")
	log.Fatal(http.ListenAndServe(":9000",nil))
}

func FindUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	usr := &User{Id: id, Name: "qms19"}
	response.WriteEntity(usr)
}