package main

import (
	"flag"
	"fmt"

	"github.com/google/uuid"
)

var n = 1

func init() {
	flag.IntVar(&n, "n", n, "Number of UUIDs to generate")
	flag.Parse()
}

func main() {
	for n > 0 {
		fmt.Println(uuid.New())
		n--
	}
}
