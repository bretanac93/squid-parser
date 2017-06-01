package main

import (
	"github.com/go-fsnotify/fsnotify"
	"log"
)

// NewWatcher Doc
func NewWatcher() {
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("EVENT: ", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Modified file: ", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("ERROR: ", err)
			}
		}
	}()
	err := watcher.Add("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func main() {
	NewWatcher()
}
