package controllers

import (
	"encoding/json"
	"go-study/internal/db"
	"go-study/internal/models"
	"net/http"
)

func GetAllCars(w http.ResponseWriter, r *http.Request) {
	db.InitDB()

	cars, err := db.GetCars()
	if err != nil {
		http.Error(w, "Erro ao obter veículos", http.StatusInternalServerError)
		return
	}

	// Converter veículos para JSON
	jsonCars, err := json.Marshal(cars)
	if err != nil {
		http.Error(w, "Erro ao converter veículos para JSON", http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type para application/json
	w.Header().Set("Content-Type", "application/json")

	// Escrever resposta com os veículos em JSON
	w.Write(jsonCars)
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	db.InitDB()

	// Decodificar o corpo da solicitação JSON em um objeto Client
	var car models.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da solicitação JSON", http.StatusBadRequest)
		return
	}

	// Inserir o novo cliente no banco de dados
	err = db.InsertCar(car)
	if err != nil {
		http.Error(w, "Erro ao inserir o carro no banco de dados", http.StatusInternalServerError)
		return
	}

	// Retornar uma resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Carro inserido com sucesso"))
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	db.InitDB()

	var car models.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da solicitação JSON", http.StatusBadRequest)
		return
	}

	// Inserir o novo cliente no banco de dados
	err = db.DeleteCar(car)
	if err != nil {
		http.Error(w, "Erro ao deletar o carro no banco de dados", http.StatusInternalServerError)
		return
	}

	// Retornar uma resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Carro deletado com sucesso"))
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	db.InitDB()

	var car models.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da solicitação JSON", http.StatusBadRequest)
		return
	}

	// Inserir o novo cliente no banco de dados
	err = db.UpdateCar(car)
	if err != nil {
		http.Error(w, "Erro ao atualizar o carro no banco de dados", http.StatusInternalServerError)
		return
	}

	// Retornar uma resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Carro atualizado com sucesso"))
}
