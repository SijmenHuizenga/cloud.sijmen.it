package actions

import "github.com/sijmen/sijmeninstaller/util"

var InstallBaseDevel = Action {
	"tooling",
	"installs base-devel",
	func() error {
		return util.Cmd("sudo pacman --noconfirm -Sy base-devel")
	},
}

var InstallAurman = Action {
	"aurman",
	"install aurman to the system",
	func() error {
		return util.Cmds(
			"git clone https://aur.archlinux.org/aurman.git",
			"cd aurman && makepkg -si --noconfirm",
		)
	},
}