package main

import (
	"log"
	"shopd/core"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) /* configuring default logger to show additional info */
	ac := &core.AppContext{}
	ac.Init()
}
