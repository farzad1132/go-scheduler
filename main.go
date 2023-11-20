package main

import (
	"example.com/scheduler/environment"
)

func main() {
	env := environment.MonolithicEnv{}
	env.Initialize()

	env.Run(100)
}
