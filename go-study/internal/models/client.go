package models

import "time"

type Client struct {
    ID        int    `json:"id"` 	
    FirstName string `json:"firstname"`
    LastName  string `json:"lastname"`
    CPF       string `json:"cpf"`
    Email     string `json:"email"`
    Phone     string `json:"phone"`
	CreatedAt time.Time `json:"createdat"`	
}