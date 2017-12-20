package main

import (
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	flag.Parse()
	blob, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	newcontents := "MOO!\n\n" + string(blob)

	if err := ioutil.WriteFile(flag.Arg(1), []byte(newcontents), 0644); err != nil {
		log.Fatal(err)
	}
}
