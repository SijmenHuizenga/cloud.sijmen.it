package actions

import (
	"github.com/sijmen/sijmeninstaller/util"
)

var InstallZSH = Action {
	"zsh",
	"Add the sijmen-user to the system. Run this command as the `sijmen` user without sudo and this will install the cool shell for both sijmen as root.",
	func() error {
		e := util.Cmds(
			"aurman -S --noconfirm --noedit powerline-fonts-git",
			"sudo pacman -S --noconfirm zsh",

			//install zsh for sijmen
			"sudo chsh -s /bin/zsh sijmen",
			`sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"`,
			//"rm ~/.zshrc",

			//install zsh for root
			"sudo chsh -s /bin/zsh",
			`sudo sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"`,
			//"sudo rm /root/.zshrc",
		)
		if e != nil {return e}

		e = util.CopyResourceToDisk("zshrc", "/home/sijmen/.zshrc")
		if e != nil {return e}

		e = util.Cmd("sudo cp /home/sijmen/.zshrc /root/.zshrc")
		if e != nil {return e}

		return nil
	},
}
