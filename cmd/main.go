package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/pappz/ha-homework/service"
	"github.com/pappz/ha-homework/web/server"
)

var (
	osSigs    = make(chan os.Signal, 1)
	wgExit    sync.WaitGroup
	webServer server.Server
)

func handleExitSignal() {
	signal.Notify(osSigs, syscall.SIGINT, syscall.SIGTERM)
	_ = <-osSigs
	log.Println("interrupt...")
	tearDownWebServer()
	wgExit.Done()
}

func tearDownWebServer() {
	_ = webServer.TearDown()
	log.Println("teardown complete")
}

func main() {
	cfg := newConfig()

	if cfg.sectorId == nil {
		log.Fatalf("must to set valid sector id")
		return
	}

	sectorService := service.NewSector(*cfg.sectorId)

	log.Printf("start webserver on: %s", cfg.listenAddress)
	webServer = server.NewServer(cfg.listenAddress, sectorService)
	webServer.Listen()

	wgExit.Add(1)
	go handleExitSignal()
	wgExit.Wait()
}
