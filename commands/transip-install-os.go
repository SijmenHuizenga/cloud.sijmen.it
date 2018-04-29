package commands

import (
	"github.com/sijmen/sijmeninstaller/actions"
	"github.com/sijmen/sijmeninstaller/actions/transip"
	"github.com/urfave/cli"
)

var TransipInstallOs = cli.Command{
	Name:  "transip-install-os",
	Usage: "Install the arch os on a brandnew transip-vps.",
	Subcommands: actions.MakeSubCommands([]actions.Action{
		actions.Network,
		actions.Mirrors,
		transip.PartitionAction,
		transip.Mount,
		actions.InstallOS,
	}),
}