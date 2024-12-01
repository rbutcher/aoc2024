package main

import (
	"github.com/rbutcher/aoc2024/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
