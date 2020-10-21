package mov

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"time"
	"unsafe"

	"github.com/Limard/mp4info/comm"
)

type MovMvhdBox struct {
	MovBaseBox
	MovMvhdBoxContent
}

type MovMvhdBoxContent struct {
	Version         byte                       // 1 box版本，0/1，一般为0
	Flags           [MVHD_FLAGS_SIZE]byte      // 3
	CreationTime    uint32                     // 4 创建时间（相对于UTC时间1904 - 01 - 01零点的秒数）
	ModificaionTime uint32                     // 4 修改时间
	TimeScale       int32                      // 4 文件媒体在1秒内的刻度值，用duration和time_scale 值可以计算track时长
	Duration        int32                      // 4 该track的时间长度，用duration和time scale值可以计算track时长
	Rate            [MVHD_RATE_SIZE]byte       // 4 推荐播放速率，高16位和低16位分别为小数点整数部分和小数部分，即[16.16] 格式.该值为1.0（0x00010000）表示正常前向播放
	Volume          [MVHD_VOLUME_SIZE]byte     // 2 与rate类似，[8.8] 格式，1.0（0x0100）表示最大音量
	Reserved        [MVHD_RESERVED_SIZE]byte   // 10 保留位
	Matrix          [MVHD_MATRIX_SIZE]byte     // 36 视频变换矩阵
	PreDefined      [MVHD_PREDEFINFO_SIZE]byte // 24
	NextTrackID     int32                      // 4 下一个track使用的ID
}

func (t *MovMvhdBox) GetCreationTime() time.Time {
	return baseTime1904.Add(time.Duration(t.CreationTime) * time.Second)
}

func (t *MovMvhdBox) GetModificaionTime() time.Time {
	return baseTime1904.Add(time.Duration(t.ModificaionTime) * time.Second)
}

func (t *MovMvhdBox) GetVolume() float32 {
	f, _ := comm.BytesToFloat32Ex(t.Volume[:1], t.Volume[1:2])
	return f
}

func (t *MovMvhdBox) GetRate() float32 {
	f, _ := comm.BytesToFloat32Ex(t.Rate[:2], t.Rate[2:4])
	return f
}

func NewMvhdBox(head *MovBaseBox, r io.ReadSeeker) (mvhd *MovMvhdBox, e error) {
	mvhd = new(MovMvhdBox)
	mvhd.BoxSize = head.BoxSize
	mvhd.BoxType = head.BoxType

	buf := make([]byte, head.BoxSize-8)
	_, e = r.Read(buf)
	if e != nil {
		log.Println(e)
		return
	}

	v := MovMvhdBoxContent{}
	e = binary.Read(bytes.NewReader(buf), binary.BigEndian, &v)
	if e != nil {
		log.Println(e)
		log.Println("MovMvhdBoxContent:", unsafe.Sizeof(v))
		log.Println("head.BoxSize:", head.BoxSize)
		return
	}

	mvhd.MovMvhdBoxContent = v
	return
}

func (mvhd *MovMvhdBox) String() string {
	str := "== mvhd == \n"
	str += fmt.Sprintf("Version:%d \n", mvhd.Version)
	str += fmt.Sprintf("Flags:%d \n", mvhd.Flags)
	str += fmt.Sprintf("CreationTime:%v \n", mvhd.GetCreationTime())
	str += fmt.Sprintf("ModificaionTime:%v \n", mvhd.GetModificaionTime())
	str += fmt.Sprintf("TimeScale:%d \n", mvhd.TimeScale)
	str += fmt.Sprintf("Duration:%d \n", mvhd.Duration)
	str += fmt.Sprintf("Rate:%v \n", mvhd.GetRate())
	str += fmt.Sprintf("Volume:%v \n", mvhd.GetVolume())
	str += fmt.Sprintf("Matrix:%s \n", mvhd.Matrix)
	str += fmt.Sprintf("PreDefined:%s \n", mvhd.PreDefined)
	str += fmt.Sprintf("NextTrackID:%d \n", mvhd.NextTrackID)
	return str
}
