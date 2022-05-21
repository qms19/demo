package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/qms19/web-demo/userservice"
	"log"
	"net/http"
)

func main()  {
	restful.Add(userservice.New())
	fmt.Println("start web-demo listen 8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}
