package main

import (
	"fmt"
	reloaded "go-reloaded/internal/handle"
)

func main() {
	if err := reloaded.Run(); err != nil {
		//nolint:forbidigo
		fmt.Printf("%s\n\n%s\n", reloaded.Error(), err)
	}
}
