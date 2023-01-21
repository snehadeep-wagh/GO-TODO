package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/snehadeep-wagh/go-todo/router"
)

func main() {
	r := mux.NewRouter()
	router.TaskRoutes(r)
	fmt.Print("Starting server on port : 9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}
