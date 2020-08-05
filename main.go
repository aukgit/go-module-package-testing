package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logger, _ = zap.NewProduction()
	logger.Info("Success")
}

func main() {
	router  : = httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Hello World")
}
