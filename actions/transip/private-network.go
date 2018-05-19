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

		e := util.Cmd("mkdir /opt/transip-privatenetwork")
		if e != nil {return e}

		serviceFilePath := "/etc/systemd/system/transip-privatenetwork.service"
		exeFilePath := "/opt/transip-privatenetwork/setup.sh"

		e = util.CopyResourceToDisk("transip-privatenetwork.service", serviceFilePath)
		if e != nil {return e}

		e = util.CopyResourceToDisk("transip-privatenetwork.sh", exeFilePath)
		if e != nil {return e}

		e = util.ReplaceInFile("THEIP", ip, exeFilePath)
		if e != nil {return e}

		return util.Cmds(
			"chmod +x "+ exeFilePath,
			"systemctl enable transip-privatenetwork.service",
			"systemctl start transip-privatenetwork.service",
		)
	},
}