package watcher

import (
	"log"
	"time"

	"github.com/radovskyb/watcher"
)

const (
	InitEvent = "INIT"
)

func Subscribe(channel chan string, path string) {
	w := watcher.New()
	defer w.Close()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Rename, watcher.Move, watcher.Write, watcher.Remove, watcher.Create)

	if err := w.AddRecursive(path); err != nil {
		log.Fatalf("Adding a new path to watcher failed: %v\n", err)
	}

	go func() {
		if err := w.Start(time.Millisecond * 100); err != nil {
			log.Fatalln(err)
		}
	}()

	// send init event to run command on start
	channel <- InitEvent

	for {
		select {
		case event := <-w.Event:
			channel <- event.Op.String()
		case err := <-w.Error:
			log.Println("Watcher error:", err)
		case <-w.Closed:
			return
		}
	}
}
