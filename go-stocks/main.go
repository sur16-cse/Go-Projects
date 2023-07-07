package main

import (
	"fmt"
	router "go-stocks/router"
	"log"
	"net/http"
)

func main()  {
	r:=router.Router()
	fmt.Println("Starting server on the port 4000...")
	log.Fatal(http.ListenAndServe(":4000",r))
}