package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/admpub/gohls-server/utils"
)

var tsFile string

func main() {
	flag.StringVar(&tsFile, `f`, ``, `-f demo.ts`)
	flag.Parse()

	if _, err := os.Stat(tsFile); err != nil {
		fmt.Println(err)
		return
	}
	saveFile := strings.TrimSuffix(tsFile, `.ts`) + `.mp4`
	err := utils.ConvertToMP4(tsFile, saveFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(`success:`, saveFile)
}
