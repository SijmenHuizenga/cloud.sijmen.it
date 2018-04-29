package actions

import (
	"github.com/urfave/cli"
	"github.com/sijmen/sijmeninstaller/util"
	"fmt"
	"errors"
)

type Key string
type Description string
type Runner func() error
type Action struct {Key; Description; Runner}

func MakeSubCommands(acts []Action) []cli.Command {
	vsm := make([]cli.Command, len(acts)+1)
	vsm[0] = mkRunAllCommand(acts)
	for i, v := range acts {
		vsm[i+1] = mkCommand(i, v)
	}
	return vsm
}

func mkCommand(i int,  v Action) cli.Command{
	return cli.Command {
		Name:  fmt.Sprintf("%d-%s", i, v.Key),
		Usage: fmt.Sprintf("%s", v.Description),
		Action: func(c *cli.Context) error { return v.Runner() },
		Before: globalBefore,
	}
}

func mkRunAllCommand(acts []Action) cli.Command{
	return cli.Command{
		Name: "all",
		Usage: "Run all steps",
		Before: globalBefore,
		Action: func(c *cli.Context) error {
			for index, element := range acts {
				util.Title(fmt.Sprintf("[%d/%d] %s...\n", index, len(acts), element.Description))
				e := element.Runner()
				if e != nil {
					return e
				}
			}
			return nil
		},
	}
}

func globalBefore(_ *cli.Context) error {
	if !util.IsSudo() {
		return errors.New("this program must be run as root (sudo)")
	}

	util.Title("Hi Sijmen! This script might do scary things!!")
	if !util.Confirm("Are you sure you want to continue? Did you type the correct command?") {
		return errors.New("you must type 'y' to continue")
	}

	return nil
}