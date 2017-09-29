package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"os"
	"strings"

	"github.com/malud/temgo/temgo"
)

var envVars = make(temgo.EnvVars)
var inlineFlag = flag.String("i", "", "-i filename")
var strictFlag = flag.Bool("s", false, "-s")

func init() {
	for _, e := range os.Environ() {
		string := strings.Split(e, "=")
		envVars[string[0]] = string[1]
	}

	if !flag.Parsed() {
		flag.Parse()
	}
}

func main() {
	var rw *bufio.ReadWriter
	var file *os.File
	if *inlineFlag != "" {
		var err error
		file, err = os.OpenFile(*inlineFlag, os.O_RDWR, 644)
		must(err)
		defer file.Close()
		rw = bufio.NewReadWriter(bufio.NewReader(file), bufio.NewWriter(file))
	} else {
		rw = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	}

	replace(rw, file)
}

func replace(rw *bufio.ReadWriter, file *os.File) {
	bytes, err := ioutil.ReadAll(rw)
	must(err)
	write := func(b []byte) {
		_, err := rw.Write(b)
		must(err)
		must(rw.Flush())
	}

	tg := temgo.New(envVars, *strictFlag)
	if tg.ContainsVariable(bytes) {
		str, err := tg.ReplaceVariables(bytes)
		must(err)
		if file != nil {
			truncate(file)
		}
		write(str)
	} else if file == nil {
		write(bytes)
	}
}

// fatal
func must(err error) {
	if err != nil {
		println("Error:", err.Error())
		os.Exit(1)
	}
}

func truncate(file *os.File) {
	err := file.Truncate(0)
	must(err)
	_, err = file.Seek(0, 0)
	must(err)
}
