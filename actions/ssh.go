package actions

import "github.com/sijmen/sijmeninstaller/util"

var SetupSSH = Action {
	"ssh",
	"setup the sijmen's ssh key, ssh-private-key-server security and enable the sshd service",
	func() error {
		e := installSshKey()
		if e != nil { return e }

		e = sshServerSecurity()
		if e != nil { return e }

		return sshServiceEnable()
	},
}


func installSshKey() error{
	userdir := "/home/sijmen"
	publickey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC/XA0bPMDydcKHmcFtikPDy/PQPxsx6L9fSA2bR8VqG50H1kS6gbdGCcTL+sPxD5yaprqabQ6Q6auj8vKoWMBUXAftvfyNy1mB6uanD1LBXi+GwfrtegeKg8FK0hcxL4rWJsrTS/LAEFCaYXW0jEb3HoyYhmHpzyNtKu3VDvb3hR7jcWD4epTBETpd+vmHv6Ak0ULnqIdJ2E9i4F9W0gmZMuCvcFPbge1YJO2gglzJ89luG9xvxgJ24bY07U1LH8eR5Y5YMrGPKdO8chtT4nzruYhHIxfcpGxT+WMSmFq3+D+MAGfqD3+dBgdbHsHx+JxSNO9fDTbgqon2Uf0IaNPAVKQ7xLUxw/CNkCgepvOV8qdNV8RZXG68u9NyVg/I1y53E0wz7VvsHcewnsDIxGidQ1m2fqwJV2Ybaz4VPVhizwYK5InGWuCZ6cRHMPGeno8DBFFiDfjRSiNkbYWqOytZ4WcrEydpbGrWL/VvpiXj609to9RxRMGjEVwQ6FJuFACcABTjXDIQ+tKR5JcDTJg2xC7hvesuisJTZdaBWqMmG3bHZUONtYPGRNiqq3VYYHXoIUJ44yLoVhPrz4RsfGisbjSK6hltLRrsH3S5LXgFtz9GArNW77aPw0E8rQPEekKG204Ssp3r6M9FxfozUHiBFtIy4OjAHxNCdwe3PSu2UQ== sijmenhuizenga@gmail.com"

	e := util.Cmds(
		"pacman --noconfirm -S openssh",
		"mkdir " + userdir + "/.ssh",
		"touch " + userdir + "/.ssh/authorized_keys",
	)
	if e != nil { return e }

	e = util.Cmd(`echo "` + publickey + `" > `+userdir+`/.ssh/authorized_keys`)
	if e != nil { return e }

	return util.Cmds(
		"chown -R sijmen:sijmen " + userdir + "/.ssh",
		"chmod 400 " + userdir+"/.ssh/authorized_keys",
		"chmod 700 " + userdir+"/.ssh",
	)
}

func sshServerSecurity() error {
	sshdconfig := "/etc/ssh/sshd_config"
	return util.ReplaceInFiles(
		[3]string{"UsePAM yes", "UsePAM no", sshdconfig},
		[3]string{"#PermitEmptyPasswords no", "PermitEmptyPasswords no", sshdconfig},
		[3]string{"#PasswordAuthentication yes", "PasswordAuthentication no", sshdconfig},
		[3]string{"PermitRootLogin without-password", "PermitRootLogin no", sshdconfig},
		[3]string{"#PubkeyAuthentication yes", "PubkeyAuthentication yes", sshdconfig},
	)
}

func sshServiceEnable() error {
	return util.Cmd("systemctl enable sshd")
}