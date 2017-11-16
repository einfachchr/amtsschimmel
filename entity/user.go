package entity

/*
  Das User-Struct hält alle Informationen zu einem User.

	Der Einfachkeit halber verzichten wir auf separate Structs für die Persistenz und JSON-Erzeugung
	(Entity-Object, DTO-Object). Wir fassen beide Sichten auf das Objekt an dieser Stelle zusammen und
	sparen damit die Dopplung des Codes.
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
