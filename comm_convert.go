package mp4info

import (
	"fmt"
	"strconv"
)

// binary.BigEndian
func bytesToInt(buf []byte) int {
	switch len(buf) {
	case 0:
		return 0
	case 1:
		return int(buf[0])
	case 2:
		return int(buf[0])<<8 | int(buf[1])
	case 3:
		return int(buf[0])<<16 | int(buf[1])<<8 | int(buf[2])
	default:
		return int(buf[0])<<24 | int(buf[1])<<16 | int(buf[2])<<8 | int(buf[3])
	}
}

func bytesToFloat32(bufa, bufb []byte) (float32, error) {
	front := bytesToInt(bufa)
	end := bytesToInt(bufb)
	s := fmt.Sprintf("%d.%d", front, end)
	tmp, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0.0, err
	}
	return float32(tmp), nil
}
