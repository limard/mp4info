package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Limard/mp4info/mov"
)

func main() {
	pfile, err := os.Open(`D:\Photo\20-10-08 湖北\GoPro\GH010307.MP4`)
	if err != nil {
		log.Panic(err.Error())
		return
	}

	box, e := mov.ParseBox(pfile)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(box.FileType.String())
	fmt.Println(box.Moive.MovieHeader.String())
}
