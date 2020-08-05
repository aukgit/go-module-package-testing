package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logger, _ := zap.NewProduction()
	logger.Info("Success")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	fmt.Println("Hello World")

	log.Fatal(http.ListenAndServe(":8080", router))
}
