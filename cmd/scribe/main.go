package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	scribe "github.com/xchacha20-poly1305/TLS-scribe"
)

var version string = "Unknown"

var (
	showVersion = false
	serverName  string
)

func init() {
	flag.BoolVar(&showVersion, "v", false, "Show version")
	flag.StringVar(&serverName, "sni", "", "Server name. Default to use server address")

	flag.Usage = func() {
		name := os.Args[0]
		fmt.Printf("Usage of %s:\n", name)
		fmt.Printf("%s <URL> <flags>\n", name)
		fmt.Println()
		fmt.Println("URL:  Target")
		fmt.Println()
		fmt.Println("flags:")
		flag.PrintDefaults()
	}

	flag.Parse()
}

func main() {
	if showVersion {
		fmt.Printf("Version: %s\n", version)
		os.Exit(0)
		return
	}

	target := flag.Arg(0)
	cert, err := scribe.Execute(target, serverName)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Print(cert)
}
