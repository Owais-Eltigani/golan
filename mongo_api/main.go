package main

import (
	"fmt"
	"log"
	"mongo/modules/controllers"
	"mongo/modules/router"
	"net/http"
)

func main() {
	fmt.Printf("the mongo api up and running...\n")

	controllers.DB_Connect()
	defer controllers.DB_Disconnect()

	PORT := ":3000"
	router := router.Route()

	log.Fatal(http.ListenAndServe(PORT, router))

}
