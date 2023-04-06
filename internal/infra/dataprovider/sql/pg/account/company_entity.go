package entity

type Account struct {
	ID int 
	// Email sql.NullString
	// LastAccess sql.NullTime
	PersonID int
	TicketsLeft int
	// PermissionLevel sql.NullInt16
	// CityID int
}

type CompanyAccount struct {
	
}