package main

import (
	"flag"
	"fmt"
	"github.com/alexellis/hmac"
	"strings"
)

func main() {
	var inputVar string
	var secretVar string

	flag.StringVar(&inputVar, "message", "", "message to create a digest from")
	flag.StringVar(&secretVar, "secret", "", "secret for the digest")
	flag.Parse()

	if len(strings.TrimSpace(secretVar)) == 0 {
		panic("--secret is required")
	}

	fmt.Printf("Computing hash for: %q\nSecret: %q\n\n", inputVar, secretVar)

	digest := hmac.Sign([]byte(inputVar), []byte(secretVar))
	fmt.Printf("Digested hash: %x\n", digest)
}