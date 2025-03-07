package main

import (
	"github.com/hedon-go-road/ddd-archetype-go/cmd"
	launcher "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-launcher"
)

func main() {
	defer cmd.StopSafe()
	launcher.SetupHTTPServer("config.yaml")
}
