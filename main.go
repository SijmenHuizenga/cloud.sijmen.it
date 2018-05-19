package main

import (
	"log"
	"os"
	"github.com/urfave/cli"
	"github.com/sijmen/sijmeninstaller/util"
	"github.com/sijmen/sijmeninstaller/commands"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Use this command to install new dektops or vps with everything"
	app.EnableBashCompletion = true
	app.Version = "v1.3"

	log.SetFlags(0)
	log.SetOutput(util.Errwriter{})

	app.Commands = []cli.Command{
		commands.TransipInstallOs,
		commands.TransipSetup,
		commands.InstallSaltMinion,
		commands.InstallSaltMaster,
		commands.ExtraTools,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
