package main

import (
	"fmt"
	"net"
	"os"

	"github.com/bishopmate/kafka-implementation-go/responsemodels"
)

var _ = net.Listen
var _ = os.Exit

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	defer l.Close()
	// fmt.Printf("Type: %T\n", l)
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}

	conn, err := l.Accept()
	defer conn.Close()
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println(string(buffer))

	messageSize := int32(1)
	correlationId := int32(7)

	mr := responsemodels.NewMessageResponse(messageSize, correlationId)
	conn.Write(mr.GetBytes())
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
}
