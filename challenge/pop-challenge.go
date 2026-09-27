package challenge

import "fmt"

func PopChallenge() {
	fmt.Println("Desafio surpresa")

	for i := rune(33); i < 123; i++ {
		fmt.Printf("%#U - %U - %v\n", i, i, string(i))
	}
}
