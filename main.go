package main

import (
	"flag"
	"log"
	"os"
	"watchd/executor"
	"watchd/watcher"
)

func main() {
	commandFlag := flag.String("command", "", "Command to execute...")
	pathFlag := flag.String("path", "", "Path to execution file...")
	flag.Parse()

	if (*commandFlag == "" && *pathFlag == "") || (*commandFlag != "" && *pathFlag != "") {
		log.Fatalf("You must specify exactly one flag: either --command or --path")
	}

	events := make(chan string)

	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Folder not found: %v,\n", err)
	}

	go watcher.Subscribe(events, path)
	go executor.Execute(events, *commandFlag, *pathFlag)

	select {} // block main to keep program working
}
