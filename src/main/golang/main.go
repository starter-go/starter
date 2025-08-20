package main

import (
	"os"

	"github.com/starter-go/starter"
)

func main() {
	m := starter.Module()
	i := starter.Init(os.Args)
	i.MainModule(m).WithPanic(true)
	i.Run()
}
