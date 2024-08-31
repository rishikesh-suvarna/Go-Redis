package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Listening on port :6379")

	// * LISTEN ON PORT 6379
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	// * ACCEPT CONNECTIONS
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	// * DEFER CONNECTION CLOSE TO BE RUN WHEN THE FUNCTION IS ABOUT TO TERMINATE
	defer conn.Close()

	// * RUNNING INFINITE LOOP TO READ DATA FROM CLIENT CONNECTIONS
	for {

		// * READING CLIENT BUFFER
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from client", err.Error())
			os.Exit(1)
		}

		/**
		* * TESTING RESP INPUTS
		* First character is $ and the follwing number represents the length of the message, i.e 5 in our case for 'Rishi'.
		* separated by CRLF '\r\n'.
		 */
		input := "$5\r\nRishi\r\n"
		reader := bufio.NewReader(strings.NewReader(input))

		// * READING FIRST BYTE
		b, err := reader.ReadByte()
		if err != nil {
			fmt.Println(err)
			return
		}
		if b != '$' {
			fmt.Println("Invalid input character")
			os.Exit(1)
		}

		size, _ := reader.ReadByte()
		strSize, _ := strconv.ParseInt(string(size), 10, 64)

		reader.ReadByte()
		reader.ReadByte()

		name := make([]byte, strSize)
		reader.Read(name)
		fmt.Println(string(name))

		conn.Write([]byte("+OK\r\n"))
	}

}
