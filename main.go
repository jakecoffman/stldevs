package main

import (
	"encoding/json"
	"log"
	"os"

	"runtime"

	"github.com/jakecoffman/stldevs/web"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	config := web.Config{}
	f, err := os.Open("config.json")
	if err != nil {
		log.Println("Couldn't find dev_config.json")
		return
	}

	json.NewDecoder(f).Decode(&config)

	web.Run(config)
}
