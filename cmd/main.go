package main

import (
	cli "ivan/CLI-Task-Tracker/internal/CLI"
	storage "ivan/CLI-Task-Tracker/internal/Storage"
	"log"
)

func main() {
	store, err := storage.NewJSONStorage("tasks.json")
	if err != nil {
		log.Fatal(err)
	}

	app := cli.NewCLI(store)
	if err = app.Run(); err != nil {
		log.Fatal(err)
	}
}
