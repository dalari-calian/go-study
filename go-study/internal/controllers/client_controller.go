package controllers

import (
	"encoding/json"
	"go-study/internal/db"
	"go-study/internal/models"
	"net/http"
)

func GetAllClients(w http.ResponseWriter, r *http.Request) {
	db.InitDB()

	clients, err := db.GetClients()
	if err != nil {
		http.Error(w, "Erro ao obter clientes", http.StatusInternalServerError)
		return
	}

	// Converter clientes para JSON
	jsonClients, err := json.Marshal(clients)
	if err != nil {
		http.Error(w, "Erro ao converter clientes para JSON", http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type para application/json
	w.Header().Set("Content-Type", "application/json")

	// Escrever resposta com os clientes em JSON
	w.Write(jsonClients)
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	db.InitDB()

	// Decodificar o corpo da solicitação JSON em um objeto Client
	var client models.Client
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da solicitação JSON", http.StatusBadRequest)
		return
	}

	// Inserir o novo cliente no banco de dados
	err = db.InsertClient(client)
	if err != nil {
		http.Error(w, "Erro ao inserir o cliente no banco de dados", http.StatusInternalServerError)
		return
	}

	// Retornar uma resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Cliente inserido com sucesso"))
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	db.InitDB()

	// Decodificar o corpo da solicitação JSON em um objeto Client
	var client models.Client
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da solicitação JSON", http.StatusBadRequest)
		return
	}

	hasCars, errMsg := db.HasAssociatedCars(client.ID)
	if hasCars {
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	// Inserir o novo cliente no banco de dados
	err = db.DeleteClient(client)
	if err != nil {
		http.Error(w, "Erro ao deletar o cliente no banco de dados", http.StatusInternalServerError)
		return
	}

	// Retornar uma resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Cliente deletado com sucesso"))
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	db.InitDB()

	var client models.Client
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da solicitação JSON", http.StatusBadRequest)
		return
	}

	// Inserir o novo cliente no banco de dados
	err = db.UpdateClient(client)
	if err != nil {
		http.Error(w, "Erro ao atualizar o cliente no banco de dados", http.StatusInternalServerError)
		return
	}

	// Retornar uma resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Cliente atualizado com sucesso"))
}
