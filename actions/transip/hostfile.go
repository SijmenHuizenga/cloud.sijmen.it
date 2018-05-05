package transip

import (
	"github.com/sijmen/sijmeninstaller/actions"
	"github.com/sijmen/sijmeninstaller/util"
)

var HostFile = actions.Action {
	"hostfile",
	"Copying hostfile...",
	func() error {
		return util.CopyResourceToDisk("transip-hostfile", "/etc/hosts")
	},
}