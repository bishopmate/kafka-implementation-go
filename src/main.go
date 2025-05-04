package main

import (
	"fmt"
	"net"
	"os"

	"github.com/bishopmate/kafka-implementation-go/src/models/responsemodels"
	"github.com/bishopmate/kafka-implementation-go/src/models/requestmodels"
)

var _ = net.Listen
var _ = os.Exit

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}
	defer l.Close()
	// fmt.Printf("Type: %T\n", l)
	

	conn, err := l.Accept()
	if err != nil{
		fmt.Println("Failed to accept incoming connection on port 9092")
		os.Exit(1)
	}
	defer conn.Close()
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	messageRequest := requestmodels.NewMessageRequestV2(buffer)
	
	// fmt.Println(buffer)
	// fmt.Println(messageRequest.MessageSize(), messageRequest.RequestHeader().RequestApiKey(), messageRequest.RequestHeader().RequestApiVersion(), messageRequest.RequestHeader().CorrelationId())

	messageSize := int32(0)
	correlationId := messageRequest.RequestHeader().CorrelationId()

	mr := responsemodels.NewMessageResponse(messageSize, correlationId)
	conn.Write(mr.GetBytes())
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
}
