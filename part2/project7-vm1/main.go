package main

import (
	"log"
	"os"
	"project7-vm1/code"
	"project7-vm1/parser"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("path must be provided")
	}
	path := os.Args[1]
	if path == "" {
		log.Fatal("path must be provided")
	}

	p := parser.NewParser(path)
	writer, err := code.NewWriter(path)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	for cmd, err := range p.Commands {
		if err != nil {
			log.Fatal(err)
		}
		switch cmd.Type {
		case parser.Push, parser.Pop:
			if err := writer.WritePushPop(cmd.Type, cmd.Arg1, cmd.Arg2); err != nil {
				log.Fatal(err)
			}
		case parser.Arithmetic:
			if err := writer.WriteArithmetic(cmd.Arg1); err != nil {
				log.Fatal(err)
			}
		default:
			log.Fatalf("unsupported command: %+v", *cmd)
		}
	}
}
