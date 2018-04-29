package commands

import (
	"github.com/urfave/cli"
	"github.com/sijmen/sijmeninstaller/actions"
)

var ExtraTools = cli.Command{
	Name:  "extra-tools",
	Usage: "Setup all the tools that are needed to make a new system usable. Run this as the sijmen user without sudo!",
	Subcommands: actions.MakeSubCommands([]actions.Action{
		actions.InstallBaseDevel,
		actions.InstallAurman,
		actions.InstallZSH,
	}, actions.ConfirmBefore),
}