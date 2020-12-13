package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Limard/mp4info/mov"
)

func main() {
	pfile, err := os.Open(`F:\Photo\2019\19-10-01 乌拉盖\Mi9\VID_20191002_204806.mp4`)
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
