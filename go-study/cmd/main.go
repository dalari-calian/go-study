package main

import (
	"go-study/internal/controllers"
	"log"
	"net/http"
	"sync"
)

var mutex sync.Mutex

func main() {

	log.Println("Sistema iniciado!")

	http.HandleFunc("/clients", handleRequest(controllers.GetAllClients))
	http.HandleFunc("/cars", handleRequest(controllers.GetAllCars))
	http.HandleFunc("/clients/create", handleRequest(controllers.CreateClient))
	http.HandleFunc("/cars/create", handleRequest(controllers.CreateCar))
	http.HandleFunc("/clients/delete", handleRequest(controllers.DeleteClient))
	http.HandleFunc("/cars/delete", handleRequest(controllers.DeleteCar))
	http.HandleFunc("/clients/update", handleRequest(controllers.UpdateClient))
	http.HandleFunc("/cars/update", handleRequest(controllers.UpdateCar))

	http.ListenAndServe(":8080", nil)
}

func handleRequest(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		defer mutex.Unlock()

		handlerFunc(w, r)
	}
}
