package main

func main() {
	requireSudo()
	titlePrinter.Println("Welcome to Sijmen's Cloud Server Installer for Arch Linux on TransIP!")
	titlePrinter.Println("Use this script to install arch on the /dev/vda disk.")
	confirm("THE /dev/vda/ VOLUME WILL BE WIPED! Are you sure you want to continue?")

	titlePrinter.Println("[1/5] Initializing network...")
	startNetwork()

	titlePrinter.Println("[2/5] Setting up mirrors...")
	mirrors()

	titlePrinter.Println("[3/5] Partitioning...")
	partition()

	titlePrinter.Println("[4/5] Mounting...")
	mount()

	titlePrinter.Println("[5/5] Installing Arch Linux...")
	install()

	titlePrinter.Println("Finished! DO NOT REBOOT YET!")
	titlePrinter.Println("Now use `arch-chroot /mnt` to get into the new system and run `setup-os` in `/root`.")
}

func partition(){
	cmdSecure("sfdisk /dev/vda < fdisk.layout")
	cmd("mkfs.vfat", "/dev/vda1", "-n", "boot")
	cmd("mkfs.ext4", "/dev/vda5", "-L", "rootfs")
	cmd("mkfs.ext4", "/dev/vda6", "-L", "datafs")
	cmd("mkswap", "/dev/vda7", "-L", "swap")
}

func mount() {
	cmd("mount", "/dev/vda5", "/mnt")
	cmd("mkdir", "/mnt/boot")
	cmd("mount", "/dev/vda1", "/mnt/boot")
	cmd("mkdir", "/mnt/data")
	cmd("mount", "/dev/vda6", "/mnt/data")
}

func mirrors(){
	tmpfile := "./mirrorlist"
	cmd("wget", "-O", tmpfile, `https://www.archlinux.org/mirrorlist/?country=NL&protocol=https`)
	replace("^#Server", "Server", tmpfile)
	cmdSecure("rankmirrors -n 6 "+ tmpfile + " > /etc/pacman.d/mirrorlist")
	cmd("rm", tmpfile)
	cmd("pacman", "-Sy")
}

func install(){
	cmd("pacstrap", "/mnt", "base")
	cmdSecure("genfstab -U /mnt >> /mnt/etc/fstab")
	cmd("cp", "/etc/pacman.d/mirrorlist", "/mnt/etc/pacman.d/mirrorlist")
	cmd("cp", "setup-os", "/mnt/root/setup-os")
}