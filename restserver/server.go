package restserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/*
  Definiert den REST-Server
*/

/*
  Das Server-"Objekt"
*/
type Server struct {
	Port int

	// Privates
	router   *mux.Router     // definiert die Endpunkte
	stopChan chan (struct{}) // Abbruchmechanismus
}

/*
  Factory-Methode. Initialisiert den Server und gibt einen Pointer darauf zurück. Hier werden auch
	die Endpunkte konfiguriert.

  Parameter
    Port:   der Port, auf dem der Server lauschen soll

*/
func NewServer(port int) *Server {
	s := &Server{
		Port:     port,
		router:   mux.NewRouter(),
		stopChan: make(chan struct{}),
	}

	// Registrieren der Routen
	s.router.HandleFunc("/", s.Home)
	s.router.HandleFunc("/authorize", s.PostAuthorize).Methods("POST")

	return s
}

/*
  Server.MachDeineArbeit

	Die Methode startet nebenläufig einen HTTP-Server und wartet dann auf die Stop-Nachricht
*/
func (s *Server) Serve() {
	go func() {

		srv := &http.Server{
			Handler: s.router,
			Addr:    fmt.Sprintf(":%d", s.Port),
			// Good practice: enforce timeouts for servers you create!
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		log.Printf("Starte Server auf Port %d\n", s.Port)
		log.Fatal(srv.ListenAndServe())
	}()

	log.Printf("Warte auf stopChan")
	<-s.stopChan
}

/*
  Beendet den Server über den eingebauten Stop-Mechanismus
*/
func (s *Server) Stop() {
	log.Print("Stoppe Server")
	s.stopChan <- struct{}{}
}

//
// ----------------------------------- REST-Endpunkte
//

/*
	GET /
		liefert eine statische Seite
*/
func (s *Server) Home(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Home Sweet Home"))
}

/*
	POST /authorize

	Die Methode erwartet Daten im Format application/x-www-form-urlencoded.

	curl -i -d "username=chef&password=geheim" -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:8080/authorize

*/
func (s *Server) PostAuthorize(rw http.ResponseWriter, r *http.Request) {

	// Auslesen der POST-Daten (Username und Passwort), falls das nicht klappt, beenden wir den Request
	var u, p, err = s.readRequestData(r)
	if err != nil {
		log.Printf("Fehler beim Auslesen der Userdaten: %s", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("User: %s, Password: %s", u, p)

	rw.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Authorize"))
}

//
// ------------------------------- Private ----------------------------------------------

/*
	Liest Username und Password aus dem Request.
*/
func (s *Server) readRequestData(req *http.Request) (string, string, error) {
	var username, password string

	req.ParseForm()

	for k, v := range req.Form {
		if k == "username" && len(v) > 0 {
			username = v[0]
		}

		if k == "password" && len(v) > 0 {
			password = v[0]
		}
	}

	if len(username) > 0 && len(password) > 0 {
		return username, password, nil
	}

	return "", "", fmt.Errorf("Konnte Usernamen und Passwort nicht auslesen")
}
