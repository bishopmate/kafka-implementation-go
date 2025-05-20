package responsemodels

import (
	"bytes"
	"encoding/binary"
)

type MessageResponse struct {
	responseSize  int32
	correlationId int32
	errorCode     int16
	arraySize     int8
	apiKey        int16
	minVersion    int16
	maxVersion    int16
}

func (mr *MessageResponse) ResponseSize() int32 {
	return mr.responseSize
}

func (mr *MessageResponse) CorrelationId() int32 {
	return mr.correlationId
}

func (mr *MessageResponse) GetBytes() []byte {
	messageBuf := new(bytes.Buffer)
	_ = binary.Write(messageBuf, binary.BigEndian, mr.correlationId)
	_ = binary.Write(messageBuf, binary.BigEndian, mr.errorCode)
	_ = binary.Write(messageBuf, binary.BigEndian, mr.arraySize)
	_ = binary.Write(messageBuf, binary.BigEndian, mr.apiKey)
	_ = binary.Write(messageBuf, binary.BigEndian, mr.minVersion)
	_ = binary.Write(messageBuf, binary.BigEndian, mr.maxVersion)
	_ = binary.Write(messageBuf, binary.BigEndian, int8(0))  // tagged fileds
	_ = binary.Write(messageBuf, binary.BigEndian, int32(0)) // throttle time
	_ = binary.Write(messageBuf, binary.BigEndian, int8(0))  // tagged fileds

	mr.responseSize = int32(messageBuf.Len())

	// responseBuf = message size buffer + message buffer
	responseBuf := new(bytes.Buffer)
	_ = binary.Write(responseBuf, binary.BigEndian, mr.responseSize)
	_ = binary.Write(responseBuf, binary.BigEndian, messageBuf.Bytes())

	// fmt.Println(responseBuf.Bytes())

	return responseBuf.Bytes()
}

func NewMessageResponse(correlationId int32, apiVersion int16) *MessageResponse {
	errorCode := int16(0)
	if apiVersion < 0 || apiVersion > 4 {
		errorCode = 35
	}
	mr := &MessageResponse{
		correlationId: correlationId,
		errorCode:     errorCode,
		arraySize:     2,
		apiKey:        18,
		minVersion:    3,
		maxVersion:    4,
	}

	return mr
}
