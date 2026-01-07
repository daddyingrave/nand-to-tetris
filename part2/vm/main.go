package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"vm/code"
	"vm/parser"
)

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	var vmFiles []string

	targetFileName := ""
	if fileInfo.IsDir() {
		targetFileName = filepath.Join(path, fileInfo.Name()+".asm")
		files, err := os.ReadDir(fileInfo.Name())
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".vm") {
				vmFiles = append(vmFiles, filepath.Join(path, file.Name()))
			}
		}
	} else if strings.HasSuffix(fileInfo.Name(), ".vm") {
		targetFileName = strings.ReplaceAll(path, ".vm", ".asm")
		vmFiles = append(vmFiles, path)
	} else {
		log.Fatalf("Unsupported file type: '%s'", fileInfo.Name())
	}

	asmFile, err := os.OpenFile(targetFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer asmFile.Close()

	for _, vmFile := range vmFiles {
		p := parser.NewParser(vmFile)
		writer, err := code.NewWriter(asmFile, vmFile)
		if err != nil {
			log.Fatal(err)
		}

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
}
