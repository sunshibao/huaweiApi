package main

import (
	"os"

	"huaweiApi/pkg/application"
)

func main() {

	if err := application.GetCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
