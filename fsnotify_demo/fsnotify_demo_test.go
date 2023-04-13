package fsnotify_demo

import (
	"fmt"
	"log"
	"testing"

	"github.com/fsnotify/fsnotify"
)

func TestFsNotify(t *testing.T) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Printf("%s %s\n", event.Name, event.Op)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("service.log")
	if err != nil {
		log.Fatal("Add failed:", err)
	}
	<-done

}
