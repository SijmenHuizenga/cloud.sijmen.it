package commands

import (
	"github.com/urfave/cli"
	"github.com/sijmen/sijmeninstaller/actions"
	"github.com/sijmen/sijmeninstaller/actions/transip"
)

var TransipSetup = cli.Command{
	Name:  "transip-setup",
	Usage: "Setup all the software for a transip-vps on arch linux.",
	Subcommands: actions.MakeSubCommands([]actions.Action{
		actions.Network,
		transip.KernelOptions,
		transip.Bootloader,
		actions.Timezone,
		actions.Locale,
		actions.Hostname,
		actions.CreateUserSijmen,
		actions.SetupSSH,
	}, actions.SudoBefore),
}