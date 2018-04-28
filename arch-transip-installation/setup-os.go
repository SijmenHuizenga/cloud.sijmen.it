package main

func main() {
	requireSudo()
	titlePrinter.Println("Welcome to Sijmen's Cloud Server Installer for Arch Linux on TransIP!")
	titlePrinter.Println("Use this script to install everything you need on a new Arch server.")
	confirm("Are you sure you want to continue?")

	titlePrinter.Println("[1/8] Starting network...")
	startNetwork()

	titlePrinter.Println("[2/8] Adjusting kernal options...")
	setupKernelOptions()

	titlePrinter.Println("[3/8] Configuring bootloader...")
	setupBootloader()

	titlePrinter.Println("[4/8] Setting up time...")
	setupTime()

	titlePrinter.Println("[5/8] Setting up locale...")
	setupLocale()

	titlePrinter.Println("[6/8] Setting up hostname...")
	setupHostname()

	titlePrinter.Println("[7/8] Creating user Sijmen...")
	setupSijmen()

	titlePrinter.Println("[8/8] Setting up SSH server...")
	setupSSh()

	titlePrinter.Println("Finished! Ready for reboot.")
}

func setupKernelOptions(){
	replace("MODULES=()", "MODULES=(virtio virtio_blk virtio_pci virtio_net virtio_ring)", "/etc/mkinitcpio.conf")
	cmd("mkinitcpio", "-p", "linux")
}

func setupBootloader(){
	cmd("pacman", "--noconfirm", "-S", "grub")
	cmd("grub-install", "--target=i386-pc", "/dev/vda")
	cmd("grub-mkconfig", "-o", "/boot/grub/grub.cfg")
}

func setupTime(){
	cmd("ln", "-sf", "/usr/share/zoneinfo/Europe/Amsterdam", "/etc/localtime")
	cmd("hwclock", "--systohc")
}

func setupLocale(){
	replace("#en_US.UTF-8 UTF-8", "en_US.UTF-8 UTF-8", "/etc/locale.gen")
	cmd("locale-gen")
}

func setupHostname(){
	hostname := ask("What is the hostname of this machine?")
	if hostname == "" {
		fatal("Host name not provided.")
	}

	cmd("touch", "/etc/hostname")
	cmdSecure(`echo "` + hostname + `" > /etc/hostname`)
}
func setupSijmen(){
	username := "sijmen"

	password := ask("What is the password of user '"+username+"'? Cannot be empty! ")
	if password == "" {
		fatal("Password must not be empty.")
	}

	cmd("pacman" ,"--noconfirm", "-S", "sudo")
	cmd("groupadd", "sudo")
	replace("# %sudo	ALL=(ALL) ALL", "%sudo	ALL=(ALL) ALL", "/etc/sudoers")

	cmd("useradd", "-m", username)
	cmdSecure(`echo -e "` + password + `\n` + password + `" | passwd ` + username)
	cmd("usermod", "-a", "-G", "sudo", username)
}

func setupSSh(){
	userdir := "/home/sijmen"
	publickey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC/XA0bPMDydcKHmcFtikPDy/PQPxsx6L9fSA2bR8VqG50H1kS6gbdGCcTL+sPxD5yaprqabQ6Q6auj8vKoWMBUXAftvfyNy1mB6uanD1LBXi+GwfrtegeKg8FK0hcxL4rWJsrTS/LAEFCaYXW0jEb3HoyYhmHpzyNtKu3VDvb3hR7jcWD4epTBETpd+vmHv6Ak0ULnqIdJ2E9i4F9W0gmZMuCvcFPbge1YJO2gglzJ89luG9xvxgJ24bY07U1LH8eR5Y5YMrGPKdO8chtT4nzruYhHIxfcpGxT+WMSmFq3+D+MAGfqD3+dBgdbHsHx+JxSNO9fDTbgqon2Uf0IaNPAVKQ7xLUxw/CNkCgepvOV8qdNV8RZXG68u9NyVg/I1y53E0wz7VvsHcewnsDIxGidQ1m2fqwJV2Ybaz4VPVhizwYK5InGWuCZ6cRHMPGeno8DBFFiDfjRSiNkbYWqOytZ4WcrEydpbGrWL/VvpiXj609to9RxRMGjEVwQ6FJuFACcABTjXDIQ+tKR5JcDTJg2xC7hvesuisJTZdaBWqMmG3bHZUONtYPGRNiqq3VYYHXoIUJ44yLoVhPrz4RsfGisbjSK6hltLRrsH3S5LXgFtz9GArNW77aPw0E8rQPEekKG204Ssp3r6M9FxfozUHiBFtIy4OjAHxNCdwe3PSu2UQ== sijmenhuizenga@gmail.com"

	cmd("pacman", "--noconfirm", "-S", "openssh")

	cmd("mkdir", userdir+"/.ssh")
	cmd("touch", userdir+"/.ssh/authorized_keys")

	cmdSecure(`echo "` + publickey + `" > `+userdir+`/.ssh/authorized_keys`)

	cmd("chown", "-R", "sijmen:sijmen", userdir+"/.ssh")
	cmd("chmod", "400", userdir+"/.ssh/authorized_keys")
	cmd("chmod", "700", userdir+"/.ssh")

	sshdconfig := "/etc/ssh/sshd_config"
	replace("UsePAM yes", "UsePAM no", sshdconfig)
	replace("#PermitEmptyPasswords no", "PermitEmptyPasswords no", sshdconfig)
	replace("#PasswordAuthentication yes", "PasswordAuthentication no", sshdconfig)
	replace("PermitRootLogin without-password", "PermitRootLogin no", sshdconfig)
	replace("#PubkeyAuthentication yes", "PubkeyAuthentication yes", sshdconfig)

	cmd("systemctl", "enable", "sshd")
}