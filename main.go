package main

import (
	"log"
	"net"
)

func main() {

	s := newServer()
	go s.run()

	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Unable to start the server: %v", err)
	}

	defer lis.Close()
	log.Println("Server listening into port :8888")

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("Unable to accept the connection: %v", err)
		}

		go s.newClient(conn)
	}

}
