package models

import "time"

type Car struct {
    ID        		int			`json:"id"` 	
	CreatedAt 		time.Time	`json:"createdat"`	
    ModelName 		string		`json:"modelname"`
    CarBrand  		string		`json:"carbrand"`
    LicensePlate 	string		`json:"licenseplate"`
    CarColor     	string		`json:"carcolor"`
    CarOwner     	string		`json:"carowner"`
	IdCarOwner		int			`json:"idcarowner"`
}
