package util

import (
	"strings"
	"github.com/fatih/color"
)

var titlePrinter = color.New(color.FgWhite, color.Bold)
var cmdPrinter = color.New(color.FgBlue, color.Bold)
var outPrinter = color.New(color.FgGreen, color.Bold)
var warnPrinter = color.New(color.FgYellow, color.Bold)
var errPrinter = color.New(color.FgRed, color.Bold)

func Title(title string){
	titlePrinter.Println(title)
}

type OutputWriter struct {}
func (o OutputWriter) Write(p []byte) (n int, err error) {
	writeTo(p, outPrinter)
	return len(p), nil
}
type WarnWriter struct {}
func (o WarnWriter) Write(p []byte) (n int, err error) {
	writeTo(p, warnPrinter)
	return len(p), nil
}

type Errwriter struct {}
func (o Errwriter) Write(p []byte) (n int, err error) {
	errPrinter.Print(string(p))
	return len(p), nil
}

func writeTo(p []byte, config *color.Color){
	s := string(p)

	suffix := ""
	prefix := ""

	if strings.HasSuffix(s, "\n") {
		s = s[:len(s)-1]
		suffix = "\n"
	}
	if strings.HasSuffix(s, "\n") {
		s = s[1:]
		prefix = "\n"
	}

	s = strings.Replace(s, "\n", "\n> ", -1)

	config.Print(prefix, "> ", s, suffix)
}