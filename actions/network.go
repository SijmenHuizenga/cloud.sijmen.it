package actions

import "github.com/sijmen/sijmeninstaller/util"

var Network = Action {
	"start-network",
	"Initializing network",
	func() error {
		return util.Cmds(
			"systemctl enable dhcpcd",
			"systemctl start dhcpcd",
			"ping -c 2 google.com")
	},
}