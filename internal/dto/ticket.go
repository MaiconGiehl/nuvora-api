package dto

type TicketInputDTO struct {
  AccountID						int
  Status 							int
  TravelID 						int
}

type TicketOutputDTO struct {
  ID 									int
	AccountID						int
  Status 							int
  TravelID 						int
}