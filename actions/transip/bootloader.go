package transip

import (
	"github.com/sijmen/sijmeninstaller/util"
	"github.com/sijmen/sijmeninstaller/actions"
)

var KernelOptions = actions.Action {
	"transip-kernel-options",
	"Setup some kernel options so the vps will actually boot",
	func() error {
		e := util.ReplaceInFile("MODULES=()", "MODULES=(virtio virtio_blk virtio_pci virtio_net virtio_ring)", "/etc/mkinitcpio.conf")
		if e != nil {return e}

		return util.Cmd("mkinitcpio -p linux")
	},
}

var Bootloader = actions.Action {
	"transip-bootloader",
	"Setup the bootloader for a transip vps",
	func() error {
		return util.Cmds(
			"pacman --noconfirm -S grub",
			"grub-install --target=i386-pc /dev/vda",
			"grub-mkconfig -o /boot/grub/grub.cfg",
		)
	},
}