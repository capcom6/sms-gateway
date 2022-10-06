package config

import (
	"fmt"
	l "log"
	"os"
	"reflect"
	"strings"
)

type empty struct{}

var packageName = strings.Split(reflect.TypeOf(empty{}).PkgPath(), "/")
var logPrefix = fmt.Sprintf("[%s] ", packageName[len(packageName)-1])
var log = l.New(os.Stdout, logPrefix, l.Ldate|l.Ltime|l.Lshortfile)
var errorLog = l.New(os.Stderr, logPrefix, l.Ldate|l.Ltime|l.Lshortfile)
