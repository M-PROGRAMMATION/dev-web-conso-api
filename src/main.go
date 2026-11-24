package main

import (
	"dev-web-conso-apii/src/routes"
	"dev-web-conso-apii/src/utils"
	"log"
	"net/http"
	"time"
)

func main() {
	addr := ":8080"
	var startTime = time.Now()

	routes.InitRoutes(startTime)

	utils.PrintBanner(addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
