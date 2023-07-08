package main

import (
	"os"

	"github.com/starter-go/starter"
)

func main() {
	mod := theModule()
	i := starter.Init(os.Args)
	i.MainModule(mod).WithPanic(true)
	i.Run()
}
