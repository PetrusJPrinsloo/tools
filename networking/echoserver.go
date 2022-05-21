package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Connection to a client did not close properly")
			log.Print(err)
		}
	}(conn)

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	log.Println("listening on port 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()

		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go echo(conn)
	}
}
