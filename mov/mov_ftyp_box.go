package mov

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

// ftyp 文件类型的描述
type MovFtypBox struct {
	MovBaseBox
	MovFtypBoxContent
}

type MovFtypBoxContent struct {
	MajorBrand       [FTYP_MAJORBRAND_SIZE]byte
	MinorVersion     int32
	CompatibleBrands [FTYP_COMPATIBLE_BRANDS_SIZE]byte
}

// 申请一个ftyp类型的box
func NewFtypBox(head *MovBaseBox, r io.ReadSeeker) (ftyp *MovFtypBox, e error) {
	ftyp = new(MovFtypBox)
	ftyp.BoxSize = head.BoxSize
	ftyp.BoxType = head.BoxType

	buf := make([]byte, ftyp.BoxSize-8)
	_, e = r.Read(buf)
	if e != nil {
		log.Println(e)
		return
	}

	content := MovFtypBoxContent{}
	e = binary.Read(bytes.NewReader(buf), binary.BigEndian, &content)
	if e != nil {
		log.Println(e)
		return
	}

	ftyp.MovFtypBoxContent = content
	return
}

// 解析ftyp的box
// func (ftyp *MovFtypBox) Parse(r io.ReadSeeker) (err error) {

// 	ftyp.MajorBrand = string(buf[:FTYP_MAJORBRAND_SIZE])
// 	buf = buf[FTYP_MAJORBRAND_SIZE:]

// 	ftyp.MinorVersion, err = comm.BytesToInt(buf[:FTYP_MINORVERSION_SIZE])
// 	if err != nil {
// 		return err
// 	}
// 	buf = buf[FTYP_MINORVERSION_SIZE:]

// 	for {
// 		if len(buf) < FTYP_COMPATIBLE_BRANDS_SIZE {
// 			break
// 		}
// 		ftyp.CompatibleBrands = append(ftyp.CompatibleBrands, string(buf[:FTYP_COMPATIBLE_BRANDS_SIZE]))
// 		buf = buf[FTYP_COMPATIBLE_BRANDS_SIZE:]
// 	}

// 	ftyp.Show()
// 	return nil
// }

func (ftyp *MovFtypBox) String() string {
	str := "== ftyp == \n"
	str += fmt.Sprintf("Major Band:%s\n", ftyp.MajorBrand)
	str += fmt.Sprintf("Monor Version:%d\n", ftyp.MinorVersion)
	str += fmt.Sprintf("Compatible Brands:%s\n", ftyp.CompatibleBrands)
	return str
}
