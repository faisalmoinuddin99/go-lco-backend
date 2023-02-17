package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/faisalmoinuddin99/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started ....")
	log.Fatal(http.ListenAndServe(":8085", r))
	fmt.Println("Listening at port 8085 ...")
}
