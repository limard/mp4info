package mov

import (
	"io"
)

type MovMdiaBox struct {
	MovBaseBox

	MediaHeader      *MovMdhdBox // mdhd
	HandlerReference *MovHdlrBox // hdlr
	// minf(vmhd, dinf ...)
}

type MovMdiaBoxContent struct {
}

func NewMdiaBox(head *MovBaseBox, r io.ReadSeeker) (mdia *MovMdiaBox, e error) {
	mdia = new(MovMdiaBox)
	mdia.BoxSize = head.BoxSize
	mdia.BoxType = head.BoxType

	return
}

func (mdia *MovMdiaBox) Parse(r io.ReadSeeker) error {
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
	// case "mdhd":
	// 	// mdhd := NewMdhdBox(base)
	// 	// e = mdhd.Parse(r)
	// 	buf := make([]byte, contentSize)
	// 	_, e = r.Read(buf)
	// case "hdlr":
	// 	// hdlr := NewHdlrBox(base)
	// 	// e = hdlr.Parse(r)
	// 	buf := make([]byte, contentSize)
	// 	_, e = r.Read(buf)
	return nil
}
