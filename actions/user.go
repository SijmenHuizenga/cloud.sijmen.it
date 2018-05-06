package actions

import (
	"github.com/sijmen/sijmeninstaller/util"
	"errors"
)

var CreateUserSijmen = Action {
	"user-sijmen",
	"Add the sijmen-user to the system",
	func() error {
		username := "sijmen"

		password := util.Ask("What is the password of user '"+username+"'? Cannot be empty! ")
		if password == "" {
			return errors.New("password must not be empty")
		}

		e := util.Cmds(
			"pacman --noconfirm -S sudo",
			"groupadd sudo",
			"useradd -m " + username,
			"echo " + username + `:` + password + " | chpasswd",
			"usermod -a -G sudo " + username,

			//disable sudo login
			"passwd -l root",
		)
		if e != nil { return e }

		return util.ReplaceInFile("# %sudo	ALL=(ALL) ALL", "%sudo	ALL=(ALL) ALL", "/etc/sudoers")
	},
}
