package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/fannyhasbi/parkir-restful-go/realtime"
)

const PORT = ":8080"

func handleIndex(w http.ResponseWriter, r *http.Request) {
	response := "This is web service for React Parking System application"
	fmt.Fprint(w, response)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handleIndex).Methods("GET", "POST")

	router.HandleFunc("/api/realtime", realtime.ReturnRealtime).Methods("GET")

	http.Handle("/", router)

	fmt.Printf("Connected to port %v", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
