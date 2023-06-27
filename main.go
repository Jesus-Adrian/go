package main

import (
	"fmt"

	"github.com/jesus-adrian/go/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1855)
	fmt.Println(estado)
	fmt.Println(texto)
}
