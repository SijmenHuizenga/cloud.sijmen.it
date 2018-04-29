package actions

import (
	"github.com/sijmen/sijmeninstaller/util"
	"errors"
)
var Hostname = Action{
	"hostname",
	"Set the hostname",
	func() error {
		hostname := util.Ask("What is the hostname of this machine?")
		if hostname == "" {
			return errors.New("host name not provided")
		}

		return util.Cmds(
			"touch /etc/hostname",
			`echo "`+hostname+`" > /etc/hostname`,
		)
	},
}