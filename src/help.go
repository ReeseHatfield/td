package main

import (
	_ "embed"
	"fmt"
)

// embed doc file into binary
//
//go:embed Docs.md
var staticDocs string

func Help() {
	fmt.Println(staticDocs)
}
