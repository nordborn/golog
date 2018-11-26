package golog

import (
	"log"
	"os"
)

// FlagsDefault provides messages format: "2018/11/26 16:57:49 golog.go:61"
const FlagsDefault = log.Ldate | log.Ltime | log.Lshortfile
const customPrefixDefault = "[main]:"

var OutDefault = os.Stdout
var ErrDefault = os.Stderr
