package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}

	data, err := io.ReadAll(conn)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "exit" || ln == "quit" {
			fmt.Fprintf(conn, "\n%s", ln)
			fmt.Println("User closed connection")
			break
		}
		io.WriteString(conn, ln)
	}
	defer conn.Close()

}
