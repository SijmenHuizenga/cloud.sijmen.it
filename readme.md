## Install a new TransIP VPS with Arch Linux

Download the script:
```
 $ wget https://github.com/SijmenHuizenga/sijmeninstaller/releases/download/v1.0/sijmeninstaller && chmod +x sijmeninstaller
```

Run the os installer to install arch to /dev/vda. Be careful! Because everything will be whiped!
```
 $ ./sijmeninstaller transip-install-os all |& tee -a installer-log.txt
```

Now, BEFORE REBOOT, use `arch-chroot /mnt` to get into the new system and run the following command to finish the installation
```
 $ cd /root
 $ ./sijmeninstaller transip-setup all |& tee -a installer-log.txt
```

Finished! Now reboot and ssh into the machine with the `sijmen` account.

Next: Tooling!

```
 $ ./sijmeninstaller extra-tools all
```



You can also run induvidual steps of the installation process. Look at the result of the following to get the exact commands:
```
 $ sijmeninstaller transip-install-os --help
```