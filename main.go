package main

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/dasbd72/rfsnotify"
)

var (
	root_dir string
)

func init() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: rfsnotify <dir>")
	}
	root_dir = os.Args[1]

	// path exist
	if _, err := os.Stat(root_dir); os.IsNotExist(err) {
		log.Fatal("Path not exist")
	}
}

func main() {
	go exec()
	watch()
}

func watch() {
	log.Println("Starting watcher")
	watcher, err := rfsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Add(root_dir)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case ev := <-watcher.Events:
			log.Println("event:", ev)
		case err := <-watcher.Errors:
			log.Println("error:", err)
		}
		log.Println(watcher.WatchList())
	}
}

func exec() {
	// sleep 1
	time.Sleep(2 * time.Second)

	os.RemoveAll(path.Join(root_dir, "test"))
	os.Mkdir(path.Join(root_dir, "test"), 0755)
	os.Mkdir(path.Join(root_dir, "test", "1"), 0755)
	os.Mkdir(path.Join(root_dir, "test", "1", "2"), 0755)
	os.Mkdir(path.Join(root_dir, "test", "1", "3"), 0755)

	// write file
	f, err := os.Create(path.Join(root_dir, "test", "1", "2", "word.txt"))
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("hello world")
	f.Close()
}
