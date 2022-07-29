//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	oas, err := entoas.NewExtension()
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}

	if err := entc.Generate(
		"./schema",
		&gen.Config{},
		entc.Extensions(
			oas,
		),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
