package main

import (
	"fmt"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()
	fmt.Println("Hello, World!")
}
