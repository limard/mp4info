package mp4info

import "time"

const (
	//base box head
	BASEBOX_HEAD_SIZE_LEN = 4
	BASEBOX_HEAD_TYPE_LEN = 4
	BASEBOX_HEAD_LEN      = BASEBOX_HEAD_SIZE_LEN + BASEBOX_HEAD_TYPE_LEN
	//ftyp box
	FTYP_MAJORBRAND_SIZE        = 4
	FTYP_MINORVERSION_SIZE      = 4
	FTYP_COMPATIBLE_BRANDS_SIZE = 4
	//mvhd box
	MVHD_VERSION_SIZE      = 1
	MVHD_FLAGS_SIZE        = 3
	MVHD_CREATIONTIME_SIZE = 4
	MVHD_MODIFYTIME_SIZE   = 4
	MVHD_TRACKID_SIZE      = 4
	MVHD_TIMESCALE_SIZE    = 4
	MVHD_DURATION_SIZE     = 4
	MVHD_RATE_SIZE         = 4
	MVHD_VOLUME_SIZE       = 2
	MVHD_RESERVED_SIZE     = 10
	MVHD_MATRIX_SIZE       = 36
	MVHD_PREDEFINFO_SIZE   = 24
	MVHD_NEXTTRACKID_SIZE  = 4

	//tkhd
	TKHD_VERSION_SIZE      = 1
	TKHD_FLAGS_SIZE        = 3
	TKHD_CREATIONTIME_SIZE = 4
	TKHD_MODIFYTIME_SIZE   = 4
	TKHD_RESERVED1_SIZE    = 4
	TKHD_DURATION_SIZE     = 4
	TKHD_RESERVED2_SIZE    = 8
	TKHD_LAYER_SIZE        = 2
	TKHD_ALTERGROUP_SIZE   = 2
	TKHD_VOLUME_SIZE       = 2
	TKHD_RESERVED3_SIZE    = 2
	TKHD_MATRIX_SIZE       = 36
	TKHD_WIDTH_SIZE        = 4
	TKHD_HEIGHT_SIZE       = 4

	//mdhd
	MDHD_VERSION_SIZE      = 1
	MDHD_FLAGS_SIZE        = 3
	MDHD_CREATIONTIME_SIZE = 4
	MDHD_MODIFYTIME_SIZE   = 4
	MDHD_TIMESCALE_SIZE    = 4
	MDHD_DURATION_SIZE     = 4
	MDHD_LANGUAGE_SIZE     = 2
	MDHD_PREDEFINED_SIZE   = 2

	//hdlr
	HDLR_VERSION_SIZE     = 1
	HDLR_FLAGS_SIZE       = 3
	HDLR_PREDEFINED_SIZE  = 4
	HDLR_HANDLERTYPE_SIZE = 4
	HDLR_RESERVED_SIZE    = 12
	HDLR_NAME_SIZE        = 0 // 未知，不确定大小

	FTYP_BOX = "ftyp" //文件类型
	MOOV_BOX = "moov" //音视频数据的metadatat信息
	MVHD_BOX = "mvhd" //电影头文件
	TRAK_BOX = "trak" //流的track头
	TKHD_BOX = "tkhd" //流信息的track头
	TREF_BOX = "tref" //track参考容器
	EDTS_BOX = "edts" //edit list 容器
	ELST_BOX = "elst" //edit list 元素信息
	MDHD_BOX = "mdhd" //media info box 媒体信息容器
	HDLR_BOX = "hdlr" //hdlr解释了媒体的播放过程信息，该box也可以被包含在meta box（meta）中。
)

var (
	baseTime1904 time.Time
)

func init() {
	baseTime1904, _ = time.Parse("2006-01-02", "1904-01-01")
}
