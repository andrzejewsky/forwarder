package main

import (
	"flag"
	"log"
	"net"

	"github.com/andrzejewsky/forwarder/connection"
)

var listeningAddr string
var forwardTo string

func init() {
	flag.StringVar(&listeningAddr, "listen", "127.0.0.1:3030", "listening addr")
	flag.StringVar(&forwardTo, "forward-to", "127.0.0.1:3306", "destination addr")
}

func main() {
	flag.Parse()

	log.Printf("Start listening on: %s", listeningAddr)
	log.Printf("Forward to: %s", forwardTo)

	listener, err := net.Listen("tcp", listeningAddr)

	if err != nil {
		log.Fatalf("Filed to setup listening: %v", err)
	}

	forwarder := connection.NewForwarder(listener, forwardTo)

	sem := make(chan bool)
	go forwarder.StartForwarding()
	<-sem
}
