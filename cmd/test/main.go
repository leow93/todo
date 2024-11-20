package main

import (
	"fmt"

	"github.com/leow93/todo/cfg"
)

func main() {
	c, err := cfg.New()
	if err != nil {
		panic(err)
	}

	fmt.Println("got a config", c)
}
