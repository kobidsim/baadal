package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kobidsim/baadal/handler"
)

func main() {
	port := "8080"
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/upload", handler.UploadHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
