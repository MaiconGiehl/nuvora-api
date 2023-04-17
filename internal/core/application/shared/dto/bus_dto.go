package dto

type BusOutputDTO struct {
	ID            int
	Number        int
	MaxPassengers int
	AccountID     int
}

func NewBusOutputDTO(
	id int,
	number int,
	maxPassengers int,
	accountId int,
) *BusOutputDTO {
	return &BusOutputDTO{
		ID:            id,
		Number:        number,
		MaxPassengers: maxPassengers,
		AccountID:     accountId,
	}
}
