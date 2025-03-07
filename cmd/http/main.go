package main

import (
	"github.com/hedon-go-road/ddd-archetype-go/cmd"
	apihttp "github.com/hedon-go-road/ddd-archetype-go/internal/api/httpapi"
)

func main() {
	defer cmd.StopSafe()
	apihttp.SetupHTTPServer()
}
