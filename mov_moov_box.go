package mp4info

import (
	"io"
)

type MovMoovBox struct {
	MovBaseBox
	MovieHeader *MovMvhdBox   // mvhd
	Track       []*MovTrakBox //trak
	// udta
	// iods
}

// 解析moov box
func NewMoovBox(head *MovBaseBox, r io.ReadSeeker) (moov *MovMoovBox, e error) {
	moov = new(MovMoovBox)
	moov.BoxType = head.BoxType
	moov.BoxSize = head.BoxSize
	unreadSize := head.BoxSize

	for {
		if unreadSize <= 0 {
			break
		}

		var base *MovBaseBox
		base, e = ParseMovBaseBox(r)
		if e != nil {
			if e == io.EOF {
				e = nil
			}
			break
		}
		unreadSize -= base.BoxSize

		if base.BoxSize <= 0 {
			break
		}

		contentSize := base.BoxSize - 8
		switch base.BoxType {
		default:
			buf := make([]byte, contentSize)
			_, e = r.Read(buf)
		case "mvhd":
			moov.MovieHeader, e = NewMvhdBox(base, r)
		case "trak":
			_, e = r.Seek(int64(contentSize), io.SeekCurrent)
			// trak, e := NewTrakBox(base)
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
