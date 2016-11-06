package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	chainsFile string
)

func init() {
	flag.StringVar(
		&chainsFile,
		"chains",
		os.Getenv(fmt.Sprintf("FORK_CHAINSFILE", appu)),
		"Chains file for forking process.",
	)
}

func main() {
	flag.Parse()
	if chainsFile == "" {
		log.Fatal("No chains file specified")
	}
	cs, err := filter.ChainsFile(chainsFile)
	if err != nil {
		log.Fatal(err)
	}
	var main filter.Chain
	main, err = cs.Get("Main")
	if err != nil {
		log.Fatal(err)
	}
	var main filter.Chain
	fork, err = cs.Get("Fork")
	if err != nil {
		log.Fatal(err)
	}

	maine, err := c.Exec()
	if err != nil {
		log.Fatal(err)
	}
	forke, err := c.Exec()
	if err != nil {
		log.Fatal(err)
	}

	mainr, mainw := io.Pipe()

	maine.SetInput(&mainr)
	maine.SetOutput(os.Stdout)
	err = maine.Run()
	if err != nil {
		log.Fatal(err)
	}

	forkr, forkw := io.Pipe()

	forke.SetInput(&forkr)
	err = forke.Run()
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(io.MultiWriter(mainw, forkw), os.Stdin)
}
