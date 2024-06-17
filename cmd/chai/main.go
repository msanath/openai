package main

import (
	"fmt"
)

func main() {
	cmd := newRootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
