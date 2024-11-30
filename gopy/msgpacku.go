package gopy

import (
	"encoding/binary"
	"fmt"
)

// stripMsgPackHeader extracts the payload data from a MessagePack extension format
func stripMsgPackHeader(data []byte) ([]byte, error) {
	// Need at least 1 byte for format
	if len(data) < 1 {
		return nil, fmt.Errorf("invalid data: empty buffer")
	}
	switch data[0] {
	case 0xC7: // ext 8 format (1-byte length)
		if len(data) < 3 {
			return nil, fmt.Errorf("invalid ext 8 format: insufficient data")
		}
		length := int(data[1])
		return data[3 : 3+length], nil

	case 0xC8: // ext 16 format (2-byte length)
		if len(data) < 4 {
			return nil, fmt.Errorf("invalid ext 16 format: insufficient data")
		}
		length := int(binary.BigEndian.Uint16(data[1:3]))
		return data[4 : 4+length], nil

	case 0xC9: // ext 32 format (4-byte length)
		if len(data) < 6 {
			return nil, fmt.Errorf("invalid ext 32 format: insufficient data")
		}
		length := int(binary.BigEndian.Uint32(data[1:5]))
		return data[6 : 6+length], nil

	case 0xD4: // fixext 1 (1 byte of data)
		if len(data) < 3 {
			return nil, fmt.Errorf("invalid fixext 1 format: insufficient data")
		}
		return data[2:3], nil

	case 0xD5: // fixext 2 (2 bytes of data)
		if len(data) < 4 {
			return nil, fmt.Errorf("invalid fixext 2 format: insufficient data")
		}
		return data[2:4], nil

	case 0xD6: // fixext 4 (4 bytes of data)
		if len(data) < 6 {
			return nil, fmt.Errorf("invalid fixext 4 format: insufficient data")
		}
		return data[2:6], nil

	case 0xD7: // fixext 8 (8 bytes of data)
		if len(data) < 10 {
			return nil, fmt.Errorf("invalid fixext 8 format: insufficient data")
		}
		return data[2:10], nil

	case 0xD8: // fixext 16 (16 bytes of data)
		if len(data) < 18 {
			return nil, fmt.Errorf("invalid fixext 16 format: insufficient data")
		}
		return data[2:18], nil

	default:
		return nil, fmt.Errorf("unsupported extension format marker: 0x%X", data[0])
	}
}
