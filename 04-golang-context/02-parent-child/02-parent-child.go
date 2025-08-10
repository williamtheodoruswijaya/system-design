package main

import (
	"context"
	"fmt"
)

func main() {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	// Print the context family tree
	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// Get Value
	fmt.Println(contextF.Value("f")) // dapet karena di context sendiri
	fmt.Println(contextF.Value("c")) // dapet tapi milik parent
	fmt.Println(contextF.Value("b")) // ga dapet beda parent

	fmt.Println(contextE.Value("d")) // ga dapet punya sibling-nya

	fmt.Println(contextA.Value("b")) // parent gabisa mengambil data child
}
