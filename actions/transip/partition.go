package transip

import (
	"github.com/sijmen/sijmeninstaller/actions"
	"github.com/sijmen/sijmeninstaller/util"
)

var PartitionAction = actions.Action {
	"partition",
	"Partitioning...",
	func() error {
		e := util.CopyResourceToDisk("transip-fdisk.layout", "transip-fdisk.layout")
		if e != nil {return e}

		e = util.Cmds(
			"sfdisk /dev/vda < transip-fdisk.layout",
			"mkfs.vfat /dev/vda1 -n boot",
			"mkfs.ext4 /dev/vda5 -L rootfs",
			"mkfs.ext4 /dev/vda6 -L datafs",
			"mkswap /dev/vda7 -L swap",
		)
		if e != nil {return e}

		return util.DeleteFileOnDisk("transip-fdisk.layout")
	},

}