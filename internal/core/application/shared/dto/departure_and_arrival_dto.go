package dto

import "time"

type DepartureOutputDTO struct {
	Time time.Time
	City string
}

func NewDepartureOutputDTO(
	time time.Time,
	city string,
) *DepartureOutputDTO {
	return &DepartureOutputDTO {
		Time: time,
		City: city,
	}
}

type ArrivalOutputDTO struct {
	Time time.Time
	City string
}

func NewArrivalOutputDTO(
	time time.Time,
	city string,
) *ArrivalOutputDTO {
	return &ArrivalOutputDTO {
		Time: time,
		City: city,
	}
}