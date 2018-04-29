package actions

import "github.com/sijmen/sijmeninstaller/util"

var Timezone = Action {
	"time",
	"Set the timezone to europe/amsterdam",
	func() error {
		return util.Cmds(
			"ln -sf /usr/share/zoneinfo/Europe/Amsterdam /etc/localtime",
			"hwclock --systohc",
		)
	},
}

var Locale = Action {
	"locale",
	"Set the locale to en_US.UTF-8",
	func() error {
		e := util.ReplaceInFile("#en_US.UTF-8 UTF-8", "en_US.UTF-8 UTF-8", "/etc/locale.gen")
		if e != nil {return e}

		return util.Cmd("locale-gen")
	},
}