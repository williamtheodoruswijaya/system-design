package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// buat object router dari http router
	router := httprouter.New()

	// inisialisasi server dimana router yang tadi dibuat itu handler-nya dan address server-nya
	server := http.Server{
		Handler: router,
		Addr: "localhost:3000",
	}

	// httprouter.Handle
	// Nah, ini ibaratnya kita bikin route di web server kita '/api/nama_route' dan kita kasih handler-nya berupa fungsi
	// Ada 3 parameter yang kita butuhkan, yaitu method HTTP-nya (GET, POST, PUT, DELETE), route-nya, dan handler-nya
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "Hello HttpRouter")
	})

	// Jalanin server-nya
	server.ListenAndServe()
}