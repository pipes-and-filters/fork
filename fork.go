package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"

	"github.com/pipes-and-filters/filters"
)

var (
	chainsFile string
)

func init() {
	flag.StringVar(
		&chainsFile,
		"chains",
		"",
		"Chains file for forking process.",
	)
}

func main() {
	flag.Parse()
	if chainsFile == "" {
		log.Fatal("No chains file specified")
	}
	cs, err := filters.ChainsFile(chainsFile)
	if err != nil {
		log.Fatal(err)
	}
	var main filters.Chain
	main, err = cs.Get("Main")
	if err != nil {
		log.Fatal(err)
	}
	var fork filters.Chain
	fork, err = cs.Get("Fork")
	if err != nil {
		log.Fatal(err)
	}

	maine, err := main.Exec()
	if err != nil {
		log.Fatal(err)
	}
	forke, err := fork.Exec()
	if err != nil {
		log.Fatal(err)
	}
	var forki bytes.Buffer
	maini := io.TeeReader(os.Stdin, &forki)
	maine.SetInput(maini)
	maine.SetOutput(os.Stdout)
	err = maine.Run()
	if err != nil {
		log.Fatal(err)
	}
	forke.SetInput(&forki)
	forke.Detach()
	go forke.Run()

	err = maine.Run()
	if err != nil {
		log.Fatal(err)
	}
	wg.Wait()

}
