package dto

type AccountInputDTO  struct{
  Username     	string
  Email   			string
  Password 			string
  PersonID 			int
  DailyTickets 	int
}

type AccountOutputDTO  struct{
	ID 						int
  Username     	string
  TicketsLeft 	int
}