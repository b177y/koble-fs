package main

import (
	"fmt"
	"os"

	"github.com/b177y/koble-fs/pkg/startup"
	colorable "github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf, err := startup.LoadConf()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not load kstart config: %v\n", err)
	}
	if conf.Quiet {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
	log.SetOutput(colorable.NewColorableStdout())
	err = rootCmd.Execute()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
