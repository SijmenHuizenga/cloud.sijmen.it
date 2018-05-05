package actions

import (
	"github.com/sijmen/sijmeninstaller/util"
)

var InstallZSH = Action {
	"zsh",
	"Run this command as the `sijmen` user without sudo and this will install the cool shell.",
	func() error {
		e := util.CopyResourceToDisk("zshrc", "/home/sijmen/.zshrc")
		if e != nil {return e}

		return util.Cmds(
			"aurman -S --noconfirm --noedit powerline-fonts-git",
			"sudo pacman -S --noconfirm zsh",

			//install zsh for sijmen
			`aurman -S --noconfirm --noedit antigen-git`,
			"sudo chsh -s /bin/zsh sijmen",
		)
	},
}
