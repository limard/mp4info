package mp4info

import (
	"io"
)

type MovBox struct {
	FileType  *MovFtypBox // ftyp
	MediaData [][]byte    // mdat
	Moive     *MovMoovBox // moov
}

// 解析box的整体调用
func ParseBox(r io.ReadSeeker) (box *MovBox, e error) {
	box = &MovBox{}
	for {
		var base *MovBaseBox
		base, e = ParseMovBaseBox(r)
		if e != nil {
			if e == io.EOF {
				e = nil
			}
			break
		}

		if base.BoxSize <= 0 {
			break
		}

		contentSize := base.BoxSize - 8
		switch base.BoxType {
		default:
			_, e = r.Seek(int64(contentSize), io.SeekCurrent)
		case "ftyp":
			box.FileType, e = NewFtypBox(base, r)
		case "moov":
			box.Moive, e = NewMoovBox(base, r)
		}
		if e != nil {
			if e == io.EOF {
				e = nil
			}
			break
		}
	}
	return
}
