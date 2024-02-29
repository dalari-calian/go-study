package tests

import (
	"bytes"
	"encoding/json"
	"go-study/internal/controllers"
	"go-study/internal/db"
	"go-study/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllClients(t *testing.T) {
	request, _ := http.NewRequest("GET", "/clients", nil)
	response := httptest.NewRecorder()

	handler := http.HandlerFunc(controllers.GetAllClients)
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Esperava o status %v, recebeu %v", http.StatusOK, response.Code)
	}
}

func TestNonExistingRoute(t *testing.T) {
	request, _ := http.NewRequest("GET", "/non-existing-route", nil)
	response := httptest.NewRecorder()

	// Neste caso, não precisamos de um handler específico, pois queremos apenas testar a rota não existente
	http.DefaultServeMux.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Esperava o status %v, recebeu %v", http.StatusNotFound, response.Code)
	}

	if response.Code == 404 {
		t.Errorf("Retornou status %v", response.Code)
	}
}

func TestCreateCar(t *testing.T) {
	// Criar um objeto Car para enviar no corpo da requisição
	newCar := models.Car{
		ModelName:    "Corsa",
		CarBrand:     "Chevrolet",
		LicensePlate: "ZZZ-9999",
		CarColor:     "Branco",
		CarOwner:     "Munir",
		IdCarOwner:   8,
	}
	// Converter o objeto Car para JSON
	jsonCar, err := json.Marshal(newCar)
	if err != nil {
		t.Fatal(err)
	}

	// Criar uma solicitação HTTP POST com o JSON do novo carro
	request, err := http.NewRequest("POST", "/cars/create", bytes.NewBuffer(jsonCar))
	if err != nil {
		t.Fatal(err)
	}

	// Registrar o manipulador para teste
	handler := http.HandlerFunc(controllers.CreateCar)

	// Criar um gravador de resposta falso para registrar a resposta da solicitação
	responseRecorder := httptest.NewRecorder()

	// Executar o manipulador HTTP com a solicitação falsa e o gravador de resposta falso
	handler.ServeHTTP(responseRecorder, request)

	// Verificar o código de status da resposta
	if responseRecorder.Code != http.StatusCreated {
		t.Errorf("Esperava o status %v, recebeu %v", http.StatusCreated, responseRecorder.Code)
	}

	// Verificar se o corpo da resposta contém a mensagem de sucesso
	expectedResponse := "Carro inserido com sucesso"
	if body := responseRecorder.Body.String(); body != expectedResponse {
		t.Errorf("Esperava a resposta '%s', mas recebeu '%s'", expectedResponse, body)
	}

	// Limpar o banco de dados após o teste
	err = db.DeleteCar(newCar)
	if err != nil {
		t.Errorf("Erro ao limpar o banco de dados após o teste: %v", err)
	}
}
