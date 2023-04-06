package mp4info

import (
	"encoding/binary"
	"io"
)

// 基础的 box 头
type MovBaseBox struct {
	BoxSize int // 包括box header和box body整个box的大小
	BoxType string
}

// 解析基础包头
func ParseMovBaseBox(r io.ReadSeeker) (box *MovBaseBox, e error) {
	b := &struct {
		Size int32
		Type [BASEBOX_HEAD_TYPE_LEN]byte
	}{}

	e = binary.Read(r, binary.BigEndian, b)
	if e != nil {
		return
	}

	box = &MovBaseBox{
		BoxSize: int(b.Size),
		BoxType: string(b.Type[:]),
	}
	return
}
