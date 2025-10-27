package main

import (
	"assembler/internal/assembler"
	"flag"
	"log"
)

var path string

func init() {
	flag.StringVar(&path, "path", "", "path to assembly file")
}

func main() {
	flag.Parse()
	if path == "" {
		log.Fatal("--path must be set")
	}

	assembler.Translate(path)
}
