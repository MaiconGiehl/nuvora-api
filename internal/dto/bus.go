package dto

type BusInputDTO struct {
	Number    			 	int  `json:"number"`
	MaxPassengers    	int  `json:"maxPassengers"`
}

type BusOutputDTO struct {
	Number    			 	int  	`json:"number"`
	MaxPassengers    	int  	`json:"maxPassengers"`
}