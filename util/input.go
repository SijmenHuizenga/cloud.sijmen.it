package util

import (
	"bufio"
	"os"
	"log"
	"strings"
)

func Ask(s string) string {
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

func Confirm(s string) bool{
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

		if strings.ToLower(strings.TrimSpace(res))[0] == 'y' {
			return true
		}
	}

	return false
}