package main

import (
	"os"
	"fmt"
	"encoding/json"
	"flag"
	"github.com/espebra/libfileinfo"
)

var usage = `Usage: fileinfo [options...] FILE

Options:
  --help  This help text.

`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Fprint(os.Stderr, usage)
	}
	fpath := flag.Arg(0)

	f, err := libfileinfo.Open(fpath)
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
