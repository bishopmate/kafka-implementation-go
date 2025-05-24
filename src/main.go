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
	
	for{
		conn, err := l.Accept()
		if err != nil{
			fmt.Println("Failed to accept incoming connection on port 9092")
			os.Exit(1)
		}
		go handleConnection(conn)
	}
	
}

func handleConnection(conn net.Conn){
	defer conn.Close()
	for{
		buffer := make([]byte, 1024)
		conn.Read(buffer)
		messageRequest := requestmodels.NewMessageRequestV2(buffer)
		
		// fmt.Println(buffer)
		// fmt.Println(messageRequest.MessageSize(), messageRequest.RequestHeader().RequestApiKey(), messageRequest.RequestHeader().RequestApiVersion(), messageRequest.RequestHeader().CorrelationId())

		correlationId := messageRequest.RequestHeader().CorrelationId()
		apiVersion := messageRequest.RequestHeader().RequestApiVersion()
		mr := responsemodels.NewMessageResponse(correlationId, apiVersion)
		conn.Write(mr.GetBytes())
	}	
}