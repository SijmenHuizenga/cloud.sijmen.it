package transip

import (
	"github.com/sijmen/sijmeninstaller/util"
	"github.com/sijmen/sijmeninstaller/actions"
	"github.com/pkg/errors"
)
var PrivateNetwork = actions.Action {
	"private-network",
	"Setup the network interfaces to connecto to the private network",
	func() error {
		ip := util.Ask("What is the ip address of this machine on the ens7 interface?")
		if ip == "" {
			return errors.New("Ip must not be null")
		}

		return util.Cmds(
			"sudo ip addr add "+ip+"/255.255.0.0 dev ens7",
			"sudo ip link set ens7 up",
		)
	},
}