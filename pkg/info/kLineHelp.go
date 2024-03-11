package info

import (
	"fmt"
)

func DisplayHelp() {
	fmt.Println(`
  _    _     _      
 | | _| |   (_)_ __   ___ 
 | |/ / |   | | '_ \ / _ \
 |   <| |___| | | | |  __/
 |_|\_\_____|_|_| |_|\___|
                          
kLine - A command-line tool for managing Kubernetes clusters and namespaces with ease.
Author - Amit Gujar

Usage:
  kLine [command] [flags]

Supported Flags: namespace, clustername

Available Commands:
  use                  Use a specific namespace
  rm                   Delete a specific namespace
  stop                 Stop a specific cluster
  start                Start a specific cluster
  status               Get the status of a specific cluster
  version              Get the version of kLine

NOTE - Please use & when using stop and start commands

Flags:
  -h, --help   help for kLine

Use "kLine [command] --help" for more information about a command.`)
}
