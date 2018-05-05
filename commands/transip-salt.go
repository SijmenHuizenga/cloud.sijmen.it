package commands

import (
	"github.com/sijmen/sijmeninstaller/actions"
	"github.com/urfave/cli"
	"github.com/sijmen/sijmeninstaller/actions/transip"
)

var InstallSaltMinion = cli.Command{
	Name:  "transip-minion",
	Usage: "Installs the salt-minion on the machine.",
	Subcommands: actions.MakeSubCommands([]actions.Action{
		transip.PrivateNetwork,
		actions.Salt,
		actions.SaltMinion,
	}, actions.SudoBefore),
}

var InstallSaltMaster = cli.Command{
	Name:  "transip-master",
	Usage: "Installs the salt-master on the machine.",
	Subcommands: actions.MakeSubCommands([]actions.Action{
		transip.PrivateNetwork,
		actions.Salt,
		actions.SaltMaster,
	}, actions.SudoBefore),
}