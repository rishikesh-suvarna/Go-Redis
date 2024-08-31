package main

import (
	"fmt"
	"net"
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
		resp := NewResp(conn)
		_, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		// fmt.Println(value)

		conn.Write([]byte("+OK\r\n"))
	}

}
