package actions

import "github.com/sijmen/sijmeninstaller/util"

var InstallOS = Action {
	"install-arch",
	"Installing Arch Linux",
	func() error {
		return util.Cmds(
			"pacstrap /mnt base",
			"genfstab -U /mnt >> /mnt/etc/fstab",
			"cp /etc/pacman.d/mirrorlist /mnt/etc/pacman.d/mirrorlist",
			"cp sijmeninstaller /mnt/root/sijmeninstaller",
		)
	},
}