package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/admpub/gohls-server/utils"
)

var tsFile string
var version = `0.0.1`
var showVer bool

func main() {
	flag.StringVar(&tsFile, `f`, ``, `-f demo.ts`)
	flag.BoolVar(&showVer, `v`, false, `-v`)
	flag.Parse()
	if showVer {
		fmt.Println(`v` + version)
		return
	}

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
