package requestmodels

import (
	"encoding/binary"
	"unsafe"
)

type MessageRequestV2 struct {
	messageSize   int32
	requestHeader *MessageRequestHeaderV2
	body          string
}


func (mr *MessageRequestV2) MessageSize() int32 {
	return mr.messageSize
}

func (mr *MessageRequestV2) SetMessageSize(messageSize int32) {
	mr.messageSize = messageSize
}

func (mr *MessageRequestV2) RequestHeader() *MessageRequestHeaderV2 {
	return mr.requestHeader
}

func (mr *MessageRequestV2) SetRequestHeader(requestHeader MessageRequestHeaderV2) {
	mr.requestHeader = &requestHeader;
}

func (mr *MessageRequestV2) Body() string {
	return mr.body
}

func NewMessageRequestV2(buffer []byte) *MessageRequestV2{
	messageRequestV2 := &MessageRequestV2{}
	var cursor int32 = 0
	cursor = parseMessageSize(messageRequestV2, buffer, cursor)
	mrh := NewMessageRequestHeaderV2(buffer[cursor:])
	messageRequestV2.SetRequestHeader(*mrh)

	return messageRequestV2
}

func parseMessageSize(messageRequestV2 *MessageRequestV2, buffer []byte, cursor int32) (cursorNext int32) {
	messageSizeNumberOfBytes := unsafe.Sizeof(messageRequestV2.MessageSize())
	messageSizeBytes := buffer[cursor:messageSizeNumberOfBytes]
	var messageSize int32
	binary.Decode(messageSizeBytes, binary.BigEndian, &messageSize)
	messageRequestV2.SetMessageSize(messageSize)
	return int32(messageSizeNumberOfBytes)
}