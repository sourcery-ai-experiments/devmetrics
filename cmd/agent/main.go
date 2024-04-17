package main

import (
	"log"

	"github.com/rybalka1/devmetrics/internal/agent"
	"github.com/rybalka1/devmetrics/internal/config"
)

func main() {
	var (
		addr                         string
		pollInterval, reportInterval int
	)
	config.SelectArgs(&addr, &pollInterval, &reportInterval)
	mAgent, err := agent.NewAgent(addr, pollInterval, reportInterval)
	if err != nil {
		log.Fatal(err)
	}
	mAgent.Start()
}
