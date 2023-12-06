package main

import (
	"fmt"
	"go-reloaded/internal/handle"
)

func main() {
	if err := reloaded.Run(); err != nil {
		fmt.Println(err)
	}
}
