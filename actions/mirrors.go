package actions

import (
	"github.com/sijmen/sijmeninstaller/util"
)

var Mirrors = Action {
	"mirrors",
	"Setting up mirrors",
	func() error {
		tmpfile := "./mirrorlist"

		e := util.Cmd("wget -O " + tmpfile + " https://www.archlinux.org/mirrorlist/?country=NL&protocol=https")
		if e != nil {
			return e
		}

		e = util.ReplaceInFile("^#Server", "Server", tmpfile)
		if e != nil {
			return e
		}

		return util.Cmds(
			"rankmirrors -n 6 "+ tmpfile + " > /etc/pacman.d/mirrorlist",
			"rm " +tmpfile,
			"pacman -Sy",
		)
	},
}

