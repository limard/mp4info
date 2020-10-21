package mov

import (
	"io"
)

type MovTrakBox struct {
	MovBaseBox
	TrackHeader *MovTkhdBox // tkhd
	Media       *MovMdiaBox // mdia
	// tref
	// edts
}

func NewTrakBox(head *MovBaseBox, r io.ReadSeeker) (trak *MovTrakBox, e error) {
	trak = new(MovTrakBox)
	trak.BoxType = head.BoxType
	trak.BoxSize = head.BoxSize

	for {
		var base *MovBaseBox
		base, e = ParseMovBaseBox(r)
		if e != nil {
			// log.Println("ParseMovBaseBox:", e)
			if e == io.EOF {
				e = nil
			}
			break
		}
		// fmt.Println(base)

		if base.BoxSize <= 0 {
			break
		}

		contentSize := base.BoxSize - 8
		switch base.BoxType {
		default:
			_, e = r.Seek(int64(contentSize), io.SeekCurrent)
		case "tkhd":
			trak.TrackHeader, e = NewTkhdBox(base, r)
		case "mdia":
			trak.Media, e = NewMdiaBox(base, r)
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

func (trak *MovTrakBox) Parse(r io.ReadSeeker) error {
	// for{
	// 	boxLen , err := ParseBox(buf)
	// 	if err != nil{
	// 		return err
	// 	}
	// 	buf = buf[boxLen:]
	// 	if len(buf) == 0{
	// 		break
	// 	}
	// }

	// case "tkhd":
	// 	// tkhd := NewTkhdBox(base)
	// 	// e = tkhd.Parse(r)
	// 	buf := make([]byte, contentSize)
	// 	_, e = r.Read(buf)
	// case "mdia":
	// 	// mdia := NewMdiaBox(base)
	// 	// e = mdia.Parse(r)
	// 	buf := make([]byte, contentSize)
	// 	_, e = r.Read(buf)

	return nil
}
