package main

import (
	"log"
	
	"github.com/spf13/cobra"
	scribe "github.com/xchacha20-poly1305/TLS-scribe"
)

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
