package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		io.WriteString(conn, "\n Welcome to the tcp server!\n")
		fmt.Println("New connection from remote address:", conn.RemoteAddr())
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(60 * time.Second))
	if err != nil {
		exit()
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "exit" || ln == "quit" {
			fmt.Println("User closed connection")
			break
		}
		fmt.Println(ln)
	}

	defer conn.Close()
	fmt.Println("Connection closed")

}

func exit() {
	log.Fatalln("CONN TIMEOUT")
}
