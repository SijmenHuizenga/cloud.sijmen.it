package actions

import (
	"github.com/sijmen/sijmeninstaller/util"
	"errors"
)

var Salt = Action {
	"salt",
	"Install the salt package",
	func() error {
		return util.Cmd("sudo pacman --noconfirm -S salt")
	},
}

var SaltMinion = Action {
	"salt-minion",
	"Start and enable the minion service",
	func() error {
		return util.Cmds(
			"sudo systemctl enable salt-minion",
			"sudo systemctl start salt-minion",
		)
	},
}

var SaltMaster = Action {
	"salt-master",
	"Start and enable the salt master",
	func() error {
		intergezicht := util.Ask("What interface does salt need to listen on? (For example 192.168.2.1): ")

		if intergezicht == "" {
			return errors.New("interface must not be empty")
		}

		util.ReplaceInFile("#interface: 0.0.0.0", "interface: "+intergezicht, "/etc/salt/master")

		return util.Cmds(
			"sudo systemctl enable salt-master",
			"sudo systemctl start salt-master",
		)
	},
}