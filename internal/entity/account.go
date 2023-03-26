package entity

import "time"

type Account  struct{
	ID 						int
  Username     	string
  Email   			string
  Password 			string
  PersonID 			int
  LastAccess 		time.Time
  TicketsLeft 	int
  DailyTickets 	int
  Created_at 		time.Time
  Updated_at 		time.Time
}