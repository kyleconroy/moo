package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	fmt.Printf("moo %s moo\n", flag.Arg(0))
}
