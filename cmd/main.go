package main

import (
	"fmt"
	reloaded "go-reloaded/internal/handle"
)

func main() {
	if err := reloaded.Run(); err != nil {
		fmt.Printf("%s\n\n%s\n", reloaded.Error(), err)
	}
}
