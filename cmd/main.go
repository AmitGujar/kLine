package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AmitGujar/kLine/pkg/control"
	"github.com/AmitGujar/kLine/pkg/info"
	"github.com/AmitGujar/kLine/pkg/namespaces"
)

func main() {
	help := flag.Bool("h", false, "Displays help message")
	version := flag.Bool("version", false, "Version Info")

	flag.Parse()
	if *help {
		info.DisplayHelp()
		os.Exit(0)
	}
	if *version {
		info.VersionInfo()
		os.Exit(0)
	}

	command, namespace, clusterName := control.ParseFlags()
	clientset, config := control.CreateClientset()

	switch command {
	case "use":
		namespaces.CheckoutNamespace(clientset, config, namespace)
	case "rm":
		namespaces.DeleteNamespace(clientset, namespace)
	case "stop":
		control.ClusterStop(clusterName)
	case "start":
		control.ClusterStart(clusterName)
	case "get":
		control.GetAdminCredentials(clusterName)
	case "status":
		control.GetClusterState(clusterName)
	default:
		fmt.Println("use '-h' for more information about a command")
	}
}
