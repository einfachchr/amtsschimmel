package service

import (
	"local/amtsschimmel/entity"
	"testing"
)

func TestAuthorizeSuccessfulLogin(t *testing.T) {

	//
	// Given: - Repo mit dem User "u", Passwort "p"
	//				- AuthentificationService
	//
	var username = "u"
	var password = "p"
	var mockRepo = NewMockUserRepo()
	var as = NewAuthentificationService(mockRepo)

	//
	// When - ein Authorizierungs-Request für den User "u" mit dem Passwort "p" kommt
	//
	var u = as.Authorize(username, password)

	//
	// Then	- wird der User gefunden und zurückgeliefert
	//
	if u == nil {
		t.Errorf("User sollte erfolgreich mit (%s/%s) authorisiert werden, Ergebnis war nil", username, password)
		t.FailNow() // hier abbrechen, sonst kriegen wir einen Dereferenzierungsfehler unten
	}

	if username != u.Username {
		t.Errorf("Falscher User gefunden, Username SOLL: %s, IST: %s", username, u.Username)
	}
}

func TestAuthorizeNoUser(t *testing.T) {
	//
	// Given: - Repo mit dem User "u", Passwort "p"
	//				- AuthentificationService
	//
	var mockRepo = NewMockUserRepo()
	var as = NewAuthentificationService(mockRepo)

	//
	// When - ein Authorizierungs-Request für den User "a" mit dem Passwort "b" kommt
	//
	var u = as.Authorize("a", "b")

	//
	// Then - wird kein User gefunden und nil zurückgeliefert
	//
	if u != nil {
		t.Errorf("User gefunden, obwohl keine hätte gefunden werden sollen")
	}
}

func TestAuthorizeWrongPassword(t *testing.T) {
	//
	// Given: - Repo mit dem User "u", Passwort "p"
	//				- AuthentificationService
	//
	var mockRepo = NewMockUserRepo()
	var as = NewAuthentificationService(mockRepo)

	//
	// When - ein Authorizierungs-Request für den User "u" mit dem Passwort "x" kommt
	//
	var u = as.Authorize("u", "x")

	//
	// Then - wird kein User gefunden und nil zurückgeliefert
	//
	if u != nil {
		t.Error("User gefunden, obwohl keine hätte gefunden werden sollen, falsches Passwort")
	}
}

// ------------------------------------------- Mocks

type MockUserRepo struct {
	User map[string]*entity.User
}

func NewMockUserRepo() *MockUserRepo {
	var m = &MockUserRepo{
		User: make(map[string]*entity.User),
	}

	m.User["u"] = &entity.User{
		ID:       1,
		Username: "u",
		Password: "p",
	}

	return m
}

func (m *MockUserRepo) FindByUsername(username string) *entity.User {
	return m.User[username]
}
