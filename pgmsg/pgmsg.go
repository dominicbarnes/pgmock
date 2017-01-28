package pgmsg

import (
	"fmt"
)

type Message interface {
	Encode() ([]byte, error)
}

func ParseBackend(typeByte byte, body []byte) (Message, error) {
	switch typeByte {
	case 'E':
		return ParseErrorResponse(body)
	case 'K':
		return ParseBackendKeyData(body)
	case 'R':
		return ParseAuthentication(body)
	case 'S':
		return ParseParameterStatus(body)
	case 'Z':
		return ParseReadyForQuery(body)
	default:
		return ParseUnknownMessage(typeByte, body)
	}
}

func ParseFrontend(typeByte byte, body []byte) (Message, error) {
	switch typeByte {
	case 'p':
		return ParsePasswordMessage(body)
	case 'X':
		return ParseTerminate(body)
	default:
		return ParseUnknownMessage(typeByte, body)
	}
}

type invalidMessageLenErr struct {
	messageType string
	expectedLen int
	actualLen   int
}

func (e *invalidMessageLenErr) Error() string {
	return fmt.Sprintf("%s body must have length of %d, but it is %d", e.messageType, e.expectedLen, e.actualLen)
}
