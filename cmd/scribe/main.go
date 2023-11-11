package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	scribe "github.com/xchacha20-poly1305/TLS-scribe"
)

var version string = "Unknown"

var (
	mainCommand = &cobra.Command{
		Use:  "scribe",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cert, err := scribe.Execute(args[0], serverName)
			if err != nil {
				log.Println(err)
			}

			print(cert)
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s", version)
		},
	}
)

var (
	serverName string
)

func init() {
	mainCommand.Flags().StringVarP(&serverName, "sni", "s", "", "Server name. (Default: as your server address)")
}

func main() {
	mainCommand.Execute()
}
