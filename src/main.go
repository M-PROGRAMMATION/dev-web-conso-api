package main

import (
	_ "dev-web-conso-apii/docs"
	"dev-web-conso-apii/src/routes"
	"dev-web-conso-apii/src/utils"
	"log"
	"net/http"
	"time"
)

// @title			Dev Web Conso API
// @version		1.0
// @description	API REST en Go développée dans le cadre du cours Développement Web / Consommation d'API (YNOV).
// @host			localhost:8080
// @BasePath		/
func main() {
	addr := ":8080"
	var startTime = time.Now()

	routes.InitRoutes(startTime)

	utils.PrintBanner(addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
