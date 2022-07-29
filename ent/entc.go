//go:build ignore
// +build ignore

package main

import (
	"log"

	"ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
)

func main() {
	oas, err := entoas.NewExtension(entoas.Spec(spec))
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}

	spec := new(ogen.Spec)
	ogent, err := ogent.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}

	if err := entc.Generate(
		"./schema",
		&gen.Config{},
		entc.Extensions(
			ogent,
			oas,
		),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
