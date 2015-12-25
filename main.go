package main

import (
	"os"
	"fmt"
	"encoding/json"
	"flag"
	"github.com/espebra/libfileinfo"
)

var (
	fpath	= flag.String("file", "", "")
)

var usage = `Usage: fileinfo [options...] --file FILE

Options:
  --file  File to process.
  --help  This help text.

`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}
	flag.Parse()

	f, err := libfileinfo.Open(*fpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	b, err := json.MarshalIndent(f, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	os.Stdout.Write(b)
	fmt.Println()
}
