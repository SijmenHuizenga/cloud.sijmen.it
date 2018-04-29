package transip

import (
	"github.com/sijmen/sijmeninstaller/actions"
	"github.com/sijmen/sijmeninstaller/util"
)

var Mount = actions.Action {
	"mount",
	"Mounting filesystems",
	func() error {
		return util.Cmds(
			"mount /dev/vda5 /mnt",
			"mkdir /mnt/boot",
			"mount /dev/vda1 /mnt/boot",
			"mkdir /mnt/data",
			"mount /dev/vda6 /mnt/data",
		)
	},
}