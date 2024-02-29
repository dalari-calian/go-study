package db

import (
	"database/sql"
	"go-study/internal/models"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "q1w2e3r4@"
	dbname   = "db_teste"
)

var DB *sql.DB

func InitDB() {
	connStr := "postgres://" + user + ":" + password + "@" + host + ":" + strconv.Itoa(port) + "/" + dbname + "?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	DB = db

}

func GetClients() ([]models.Client, error) {
	rows, err := DB.Query("SELECT * FROM Client")
	if err != nil {
		log.Println("Erro ao consultar clientes:", err)
		return nil, err
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		err := rows.Scan(&client.ID, &client.FirstName, &client.LastName, &client.CPF, &client.Email, &client.Phone, &client.CreatedAt)
		if err != nil {
			log.Println("Erro ao ler resultado da consulta:", err)
			return nil, err
		}
		clients = append(clients, client)
	}
	if err := rows.Err(); err != nil {
		log.Println("Erro ao iterar sobre os resultados da consulta:", err)
		return nil, err
	}

	return clients, nil
}

func InsertClient(client models.Client) error {
	// Preparar a declaração SQL para inserção
	stmt, err := DB.Prepare("INSERT INTO Client (firstname, lastname, cpf, email, phone, createdat) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Println("Erro ao preparar a declaração SQL para inserção de cliente:", err)
		return err
	}
	defer stmt.Close()

	// Executar a declaração SQL com os valores do cliente
	_, err = stmt.Exec(client.FirstName, client.LastName, client.CPF, client.Email, client.Phone, time.Now())
	if err != nil {
		log.Println("Erro ao executar a declaração SQL para inserção de cliente:", err)
		return err
	}

	log.Println("Cliente inserido com sucesso!")
	return nil
}

func DeleteClient(client models.Client) error {
	// Preparar a declaração SQL para inserção
	stmt, err := DB.Prepare("DELETE FROM Client WHERE id = $1")
	if err != nil {
		log.Println("Erro ao preparar a declaração SQL para deletar o cliente:", err)
		return err
	}
	defer stmt.Close()

	// Executar a declaração SQL com os valores do cliente
	_, err = stmt.Exec(client.ID)
	if err != nil {
		log.Println("Erro ao executar a declaração SQL para deletar o cliente:", err)
		return err
	}

	log.Println("Cliente deletado com sucesso!")
	return nil
}

func UpdateClient(client models.Client) error {
	stmt, err := DB.Prepare("UPDATE Client " +
		"SET firstname        = $1," +
		"    lastname         = $2," +
		"    cpf              = $3," +
		"    email            = $4," +
		"    phone            = $5 " +
		"WHERE id = $6")
	if err != nil {
		log.Println("Erro ao preparar a declaração SQL para atualizar o cliente:", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(client.FirstName, client.LastName, client.CPF, client.Email, client.Phone, client.ID)
	if err != nil {
		log.Println("Erro ao executar a declaração SQL para atualizar o carro:", err)
		return err
	}

	log.Println("Client atualizado com sucesso!")
	return nil
}

func HasAssociatedCars(clientID int) (bool, string) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM Car WHERE idCarOwner = $1", clientID).Scan(&count)
	if err != nil {
		log.Println("Erro ao verificar veículos associados:", err)
		return false, "Erro ao verificar se o cliente possui veículos associados"
	}
	if count > 0 {
		return true, "Este cliente possui veículos associados e não pode ser excluído!"
	}
	return false, ""
}

func GetCars() ([]models.Car, error) {
	rows, err := DB.Query("SELECT * FROM Car")
	if err != nil {
		log.Println("Erro ao consultar veículos:", err)
		return nil, err
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		err := rows.Scan(&car.ID, &car.CreatedAt, &car.ModelName, &car.CarBrand, &car.LicensePlate, &car.CarColor, &car.CarOwner, &car.IdCarOwner)
		if err != nil {
			log.Println("Erro ao ler resultado da consulta:", err)
			return nil, err
		}
		cars = append(cars, car)
	}
	if err := rows.Err(); err != nil {
		log.Println("Erro ao iterar sobre os resultados da consulta:", err)
		return nil, err
	}

	return cars, nil
}

func InsertCar(car models.Car) error {
	// Preparar a declaração SQL para inserção
	stmt, err := DB.Prepare("INSERT INTO Car (createdat, modelname, carbrand, licenseplate, carcolor, carowner, idcarowner) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		log.Println("Erro ao preparar a declaração SQL para inserção de carro:", err)
		return err
	}
	defer stmt.Close()

	// Executar a declaração SQL com os valores do carro
	_, err = stmt.Exec(time.Now(), car.ModelName, car.CarBrand, car.LicensePlate, car.CarColor, car.CarOwner, car.IdCarOwner)
	if err != nil {
		log.Println("Erro ao executar a declaração SQL para inserção de carro:", err)
		return err
	}

	log.Println("Carro inserido com sucesso!")
	return nil
}

func DeleteCar(car models.Car) error {
	stmt, err := DB.Prepare("DELETE FROM Car WHERE id = $1")
	if err != nil {
		log.Println("Erro ao preparar a declaração SQL para deletar carro:", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(car.ID)
	if err != nil {
		log.Println("Erro ao executar a declaração SQL para deletar o carro:", err)
		return err
	}

	log.Println("Carro inserido com sucesso!")
	return nil
}

func UpdateCar(car models.Car) error {
	stmt, err := DB.Prepare("UPDATE Car " +
		"SET modelname        = $1," +
		"    carbrand         = $2," +
		"    licenseplate     = $3," +
		"    carcolor         = $4," +
		"    carowner         = $5," +
		"    idcarowner       = $6," +
		"WHERE id = $7")
	if err != nil {
		log.Println("Erro ao preparar a declaração SQL para atualizar o carro:", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(car.ModelName, car.CarBrand, car.LicensePlate, car.CarColor, car.IdCarOwner, car.ID)
	if err != nil {
		log.Println("Erro ao executar a declaração SQL para atualizar o carro:", err)
		return err
	}

	log.Println("Carro atualizado com sucesso!")
	return nil
}
