package main

import (
	"strings"

	"github.com/urfave/cli/v2"
)

type config struct {
	benchCmd       string
	base           string
	compare        []string
	benchArgs      []string
	onlyDegression bool
	threshold      float64
}

func newConfig(c *cli.Context) config {
	return config{
		onlyDegression: c.Bool("only-degression"),
		threshold:      c.Float64("threshold"),
		base:           c.String("base"),
		compare:        strings.Split(c.String("compare"), ","),
		benchCmd:       c.String("bench-cmd"),
		benchArgs:      strings.Fields(c.String("bench-args")),
	}
}
