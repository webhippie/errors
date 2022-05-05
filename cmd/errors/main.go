package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/webhippie/errors/pkg/command"
)

func main() {
	if env := os.Getenv("ERRORS_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
