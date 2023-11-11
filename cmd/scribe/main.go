package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	scribe "github.com/xchacha20-poly1305/TLS-scribe"
)

var (
	version     string = "Unknown"
	showVersion bool   = false
)

var (
	mainCommand = &cobra.Command{
		Use: "scribe",
		// Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Show version
			if showVersion {
				fmt.Printf("Version: %s", version)
				os.Exit(0)
				return
			}

			cert, err := scribe.Execute(args[0], serverName)
			if err != nil {
				log.Println(err)
			}

			print(cert)
		},
	}
)

var (
	serverName string
)

func init() {
	mainCommand.Flags().StringVarP(&serverName, "sni", "s", "", "Server name. (Default: as your server address)")
	mainCommand.Flags().BoolVarP(&showVersion, "version", "v", false, "Print the version")
}

func main() {
	mainCommand.Execute()
}
