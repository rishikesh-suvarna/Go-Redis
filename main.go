package main

import (
	"fmt"
	"net"
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

	// * AOF READER
	aof, err := NewAof("database.aof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer aof.Close()

	// * READING FROM THE FILE & WRITING INTO MEMORY FOR SUPER FAST ACCESS
	aof.Read(func(value Value) {
		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		fmt.Println(command)
		fmt.Println(args)

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			return
		}

		handler(args)
	})

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
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		if value.typ != "array" {
			fmt.Println("Invalid, Expected an array")
			continue
		}

		if len(value.array) == 0 {
			fmt.Println("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		writer := NewWriter(conn)

		handler, ok := Handlers[command]

		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(Value{typ: "string", str: ""})
			continue
		}

		if command == "SET" || command == "HSET" {
			aof.Write(value)
		}

		result := handler(args)
		writer.Write(result)
	}

}
