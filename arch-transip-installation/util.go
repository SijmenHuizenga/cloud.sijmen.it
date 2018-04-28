package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
	)

var titlePrinter = color.New(color.FgWhite, color.Bold)
var cmdPrinter = color.New(color.FgBlue, color.Bold)
var outPrinter = color.New(color.FgGreen, color.Bold)
var warnPrinter = color.New(color.FgYellow, color.Bold)
var errPrinter = color.New(color.FgRed, color.Bold)

func replace(original string, new string, file string){
	cmd("sed", "-i", "s/"+original+"/"+new+"/g", file)
}

func startNetwork(){
	cmd("systemctl", "enable", "dhcpcd")
	cmd("systemctl", "start", "dhcpcd")
	cmd("ping", "-c", "2", "google.com")
}

func requireSudo() {
	cmd := exec.Command("id", "-u")
	output, err := cmd.Output()

	if err != nil {
		fatalError(err)
	}
	i, err := strconv.Atoi(string(output[:len(output)-1]))

	if err != nil {
		fatalError(err)
	}

	if i != 0 {
		fatal("This program must be run as root! (sudo)")
	}
}

func cmd(name string, args ...string){
	cmdPrinter.Println(name, args)

	cmd := exec.Command(name, args...)
	color.Set(color.FgYellow)
	cmd.Stdout = outputwriter{}
	cmd.Stderr = warnwriter{}
	color.Unset()

	if err := cmd.Run(); err != nil {
		fatalError(err)
	}
}
func cmdSecure(name string){
	cmdPrinter.Println(name)

	cmd := exec.Command("sh", "-c", name)
	color.Set(color.FgYellow)
	cmd.Stdout = outputwriter{}
	cmd.Stderr = warnwriter{}
	color.Unset()

	if err := cmd.Run(); err != nil {
		fatalError(err)
	}
}

type outputwriter struct {}
func (o outputwriter) Write(p []byte) (n int, err error) {
	writeTo(p, outPrinter)
	return len(p), nil
}
type warnwriter struct {}
func (o warnwriter) Write(p []byte) (n int, err error) {
	writeTo(p, warnPrinter)
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

func fatalError(err error) {
	errPrinter.Println(err)
	os.Exit(1)
}
func fatal(msg string) {
	errPrinter.Println(msg)
	os.Exit(1)
}

func ask(s string) string {
	r := bufio.NewReader(os.Stdin)
	tries := 3

	for ; tries > 0; tries-- {
		titlePrinter.Printf("%s: ", s)

		res, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Empty input (i.e. "\n")
		if len(res) < 2 {
			continue
		}

		return strings.TrimSpace(res)
	}

	return ""
}

func confirm(s string) {
	r := bufio.NewReader(os.Stdin)
	tries := 3

	for ; tries > 0; tries-- {
		titlePrinter.Printf("%s [y/n]: ", s)

		res, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Empty input (i.e. "\n")
		if len(res) < 2 {
			continue
		}

		if strings.ToLower(strings.TrimSpace(res))[0] != 'y' {
			fatal("")
		}
		return
	}

	fatal("")
}