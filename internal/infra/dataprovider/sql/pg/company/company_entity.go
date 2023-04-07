package entity

type Company struct {
	ID int 
	FantasyName string
	// Email sql.NullString
	// LastAccess sql.NullTime
	// TicketsLeft sql.NullInt16
	PermissionLevel int
}
