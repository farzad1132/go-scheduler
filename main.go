package main

import (
	"example.com/scheduler/environment"
)

func main() {
	env := environment.SimpleEnv{}
	env.Initialize()

	env.Run(50)
}
