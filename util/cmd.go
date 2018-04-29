package util

import (
	"os/exec"
	"strconv"
	"fmt"
	"github.com/fatih/color"
)

func ReplaceInFile(original string, new string, file string) error {
	return Cmd(fmt.Sprintf("sed -i 's/%s/%s/g' %s", original, new, file))
}

func ReplaceInFiles(commands ...[3]string) error{
	for _, command := range commands {
		if err := ReplaceInFile(command[0], command[1], command[2]); err != nil {
			return err
		}
	}
	return nil
}

func IsSudo() bool {
	cmd := exec.Command("id", "-u")
	output, err := cmd.Output()

	if err != nil {
		return false
	}
	i, err := strconv.Atoi(string(output[:len(output)-1]))

	if err != nil {
		return false
	}

	return i == 0
}

func Cmd(command string) error{
	cmdPrinter.Println(command)

	cmd := exec.Command("sh", "-c", command)
	color.Set(color.FgYellow)
	cmd.Stdout = OutputWriter{}
	cmd.Stderr = WarnWriter{}
	color.Unset()

	return cmd.Run()
	return nil
}

func Cmds(commands ...string) error{
	for _, command := range commands {
		if err := Cmd(command); err != nil {
			return err
		}
	}
	return nil
}

