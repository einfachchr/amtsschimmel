package main

/*
  Der Einstiegspunkt für den Server.
*/

import (
	"fmt"
	"local/amtsschimmel/restserver"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	server   *restserver.Server
	stopChan chan (struct{})
)

/*
   Die Start-Routine.
*/
func main() {
	log.Println("---- Start ")

	initialize()
	run()

	// Warten auf SIGINT oder SIGKILL, um das Programm zu beenden
	waitForStop()

	log.Printf("---- Ende")
}

/*
  initialize sorgt für die notwendige Verkabelung, bzw. Initialisierung der Komponenten
*/
func initialize() {
	server = restserver.NewServer(8080)
}

/*
  run führt das eigendliche Programm aus. Alle notwendigen Komponenten müssen initialisiert sein.
*/
func run() {
	server.Serve()
}

/*
	Die Routine hängt sich an SIGINT und SIGKILL, um den Prozess dann geordnet zu beenden. Bis eines
	der Signale kommt, blockiert sie
*/
func waitForStop() {
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- struct{}{}
	}()

	log.Printf("Warte auf SIGINT")
	<-done
	server.Stop()
}
