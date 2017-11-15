package entity

/*
  Das User-Struct hält alle Informationen zu einem User
*/

type User struct {

	// Infrastruktur
	ID       int
	Username string
	Password string

	// Nutzdaten
	Firstname string
	Lastname  string
	Email     string
	LKZ       string
}
