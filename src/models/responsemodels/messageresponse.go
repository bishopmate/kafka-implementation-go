package responsemodels

import (
	"bytes"
	"encoding/binary"
)

type MessageResponse struct {
	responseSize  int32
	correlationId int32
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
	return responseBuf.Bytes()
}

func NewMessageResponse(responseSize int32, correlationId int32) *MessageResponse {
	mr := &MessageResponse{
		responseSize:  responseSize,
		correlationId: correlationId,
	}
	return mr
}
