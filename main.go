package main

import (
	"os"
)

func main() {
	a := App{}
	a.Init(os.Getenv("ENV"))
	a.Run()
}
