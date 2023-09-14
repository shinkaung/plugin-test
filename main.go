package main

import (
	"test/core"
	"test/plugins/helloworld"
)

func main() {
	orchestrator := core.NewOrchestrator()

	// Register HelloWorld plugin
	orchestrator.RegisterPlugin("hello", &helloworld.HelloWorldPlugin{})

	// Simulating command execution
	orchestrator.Execute("hello", nil)
}
