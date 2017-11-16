package service

import (
	"local/amtsschimmel/entity"
	"local/amtsschimmel/repo"
)

/*
  Der Authentifizierungs-Service kapselt die Fachlichkeit der Authentifizierung. Da in dem Thema
  nicht viel Fachlogik zu finden ist, wird hier ein Großteil der Infrastruktur mit erledigt.

*/

// das zentrale "Objekt"
type AuthentificationService struct {
	userRepo repo.UserRepo
}

func NewAuthentificationService(userRepo repo.UserRepo) *AuthentificationService {
	a := &AuthentificationService{
		userRepo: userRepo,
	}
	return a
}

/*
  Die wesentliche Aufgabe des Dienstes ist es, User zu autentifizieren. Dazu nutzt er Username und
  Passwort. Bei Erfolg werden Informationen zum User zurückgegeben. Konnte mit der Kombination
  kein User authorisiert werden (falscher User oder falsches Kennwort), kommt nil zurück.

	In der aktuellen Version findet ein simpler Abgleich des Passworts statt.
*/
func (a *AuthentificationService) Authorize(u, p string) *entity.User {
	var user = a.userRepo.FindByUsername(u)

	// keine weiteren Prüfungen nötig
	if user == nil {
		return nil
	}

	if user.Password != p {
		return nil
	}

	// Erfolg!
	return user
}
