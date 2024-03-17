package main

import (
	"github.com/Woutjeee/racing_cli/internal"
	"github.com/Woutjeee/racing_cli/internal/configuration"
)

func main() {
	config := configuration.Config{}

	internal.StartReplLoop(&config)
}
