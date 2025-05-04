package requestmodels

import (
	"encoding/binary"
	"bytes"
)

type MessageRequestHeaderV2 struct {
	requestApiKey     int16
	requestApiVersion int16
	correlationId     int32
	clientId          string // nullable field
}

func (mrh *MessageRequestHeaderV2) RequestApiKey() int16 {
	return mrh.requestApiKey
}

func (mrh *MessageRequestHeaderV2) RequestApiVersion() int16 {
	return mrh.requestApiVersion
}

func (mrh *MessageRequestHeaderV2) CorrelationId() int32 {
	return mrh.correlationId
}

func (mrh *MessageRequestHeaderV2) ClientId() string {
	return mrh.clientId
}

func NewMessageRequestHeaderV2(buffer []byte) *MessageRequestHeaderV2 {
	var cursor int32 = 0
	messageReqHeader := &MessageRequestHeaderV2{}
	
	cursor = parseBytes(buffer, cursor, &messageReqHeader.requestApiKey)
	cursor = parseBytes(buffer, cursor, &messageReqHeader.requestApiVersion)
	_ = parseBytes(buffer, cursor, &messageReqHeader.correlationId)

	return messageReqHeader
}

func parseBytes(buffer []byte, cursor int32, x interface{}) (cursorNext int32) {
	numberOfBytes := binary.Size(x)
	cursorNext = cursor + int32(numberOfBytes)
	messageSizeBytes := buffer[cursor:cursorNext]
	reader := bytes.NewReader(messageSizeBytes)
	_ = binary.Read(reader, binary.BigEndian, x)
	return cursorNext
}
