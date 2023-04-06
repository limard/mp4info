package mp4info

import (
	"flag"
	"fmt"
	"log"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	flag.Parse()
	pfile, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Panic(err.Error())
		return
	}

	box, e := ParseBox(pfile)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(box.FileType.String())
	fmt.Println(box.Moive.MovieHeader.String())
}
