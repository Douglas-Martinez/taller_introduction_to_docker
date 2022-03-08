package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

var PORT = "8080"

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc (
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			
			next.ServeHTTP(w, req)
		})
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)
	
	router.Use(middlewareCors)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Participacion desde el contenedor de docker, saludos")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	enableCORS(router)
	
	router.HandleFunc("/", rootHandler).Methods("GET")

	fmt.Println("Servidor corriendo en puerto", PORT)
	if err := http.ListenAndServe(":"+PORT, router); err != nil {
		log.Fatal(err)
		return
	}
}