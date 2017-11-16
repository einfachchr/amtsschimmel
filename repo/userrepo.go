package repo

import "local/amtsschimmel/entity"

/*
  Das UserRepo kapselt alle Funktionalität für das Speichern und Suchen der User. Es ist als
  Interface abstrahiert, um die Tests zu erleichern. Die Implementation ist im UserRepository
*/

type UserRepo interface {

	/*
		FindByUsername sucht das User-Objekt anhand des Usernamens (User.Username). Wird keiner gefunden,
		kommt nil zurück.
	*/
	FindByUsername(username string) *entity.User
}
