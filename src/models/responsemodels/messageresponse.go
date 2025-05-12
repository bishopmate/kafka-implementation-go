package responsemodels

import (
	"bytes"
	"encoding/binary"
)

type MessageResponse struct {
	responseSize  int32
	correlationId int32
	errorCode int16
	arraySize int8
	apiKey int16
	minVersion int16
	maxVersion int16
}

func (mr *MessageResponse) ResponseSize() int32 {
	return mr.responseSize
}

func (mr *MessageResponse) CorrelationId() int32 {
	return mr.correlationId
}

func (mr *MessageResponse) GetBytes() []byte {
	responseBuf := new(bytes.Buffer)
	_ = binary.Write(responseBuf, binary.BigEndian, mr.responseSize)
	_ = binary.Write(responseBuf, binary.BigEndian, mr.correlationId)
	_ = binary.Write(responseBuf, binary.BigEndian, mr.errorCode)
	_ = binary.Write(responseBuf, binary.BigEndian, mr.arraySize)
	_ = binary.Write(responseBuf, binary.BigEndian, mr.apiKey)
	_ = binary.Write(responseBuf, binary.BigEndian, mr.minVersion)
	_ = binary.Write(responseBuf, binary.BigEndian, mr.maxVersion)
	
	return responseBuf.Bytes()
}

func NewMessageResponse(responseSize int32, correlationId int32) *MessageResponse {
	mr := &MessageResponse{
		correlationId: correlationId,
		errorCode: 0,
		arraySize: 1,
		apiKey: 18,
		minVersion: 0,
		maxVersion: 4,
	}
	mr.responseSize = int32(binary.Size(mr.correlationId) + binary.Size(mr.errorCode) + binary.Size(mr.arraySize) +binary.Size(mr.apiKey) + binary.Size(mr.minVersion) + binary.Size(mr.maxVersion))
	return mr
}
