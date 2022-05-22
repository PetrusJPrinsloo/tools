package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

//Todo: replace these contants with passed in argument or config?
const blocked_site = "blockedwebsite.com"
const blocked_site_port = 80

const unblocked_site_port = 80

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", fmt.Sprintf("%s:%d", blocked_site, blocked_site_port))
	if err != nil {
		log.Fatalln("Unable to connect to blockedwebsite.com")
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%d", unblocked_site_port))
	if err != nil {
		log.Fatalln("Unable to bind port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}
