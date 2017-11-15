package service

import "local/amtsschimmel/entity"

/*
  Der Authentifizierungs-Service kapselt die Fachlichkeit der Authentifizierung. Da in dem Thema
  nicht viel Fachlogik zu finden ist, wird hier ein Großteil der Infrastruktur mit erledigt.

*/

// das zentrale "Objekt"
type AuthentificationService struct{}

func NewAutentificationService() *AuthentificationService {
	a := &AuthentificationService{}
	return a
}

/*
  Die wesentliche Aufgabe des Dienstes ist es, User zu autentifizieren. Dazu nutzt er Username und
  Passwort. Bei Erfolg werden Informationen zum User zurückgegeben. Konnte mit der Kombination
  kein User authorisiert werden (falscher User oder falsches Kennwort), kommt nil zurück.
*/
func (a *AuthentificationService) Authorize(u, p string) *entity.User {

	return nil
}
