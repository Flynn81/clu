package main

import (
	"flynn81/clu/api"
	"fmt"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()
	fmt.Println("Hello, World!")

	w := api.BasicWalker{}
	w.Walk("./")
}
