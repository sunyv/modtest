package modtest

import (
	"fmt"

	"github.com/google/uuid"
)

func GOMODTEST() {
	id := uuid.New().String()
	fmt.Println("UUID: ", id)

	p := message.NewPrinter(language.BritishEnglish)
	p.Printf("Number format: %v.\n", 1500)

	p = message.NewPrinter(language.Greek)
	p.Printf("Number format: %v.\n", 1500)
}