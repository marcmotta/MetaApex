// cmd/metaapex/main.go
package main

import (
	"flag"
	"log"
	"os"

	"metaapex/internal/metaapex"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := metaapex.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
