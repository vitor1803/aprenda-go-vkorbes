package main

import (
	"fmt"
	"os"

	"github.com/vitor1803/aprenda-go-vkorbes/challenge"
	"github.com/vitor1803/aprenda-go-vkorbes/exec/level01"
	"github.com/vitor1803/aprenda-go-vkorbes/exec/level02"
	"github.com/vitor1803/aprenda-go-vkorbes/exec/level03"
	"github.com/vitor1803/aprenda-go-vkorbes/exec/level04"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run hello-world.go <exercise_number>")
		return
	}
	switch os.Args[1] {
	case "l1":
		fmt.Println("Level 1")
		level01.Ex_01()
		level01.Ex_02()
		level01.Ex_03()
		level01.Ex_04()
		level01.Ex_05()
	case "l2":
		fmt.Println("Level 2")
		level02.Ex_01()
		level02.Ex_02()
		level02.Ex_03()
		level02.Ex_04()
		level02.Ex_05()
		level02.Ex_06()
	case "l3":
		level03.Ex_01()
		level03.Ex_02()
		level03.Ex_03()
		level03.Ex_04()
		level03.Ex_05()
		level03.Ex_06()
		level03.Ex_07()
		level03.Ex_08()
		level03.Ex_09()
		level03.Ex_10()
	case "l4":
		level04.Ex_01()
		level04.Ex_02()
		level04.Ex_03()
		level04.Ex_04()
		level04.Ex_05()
		level04.Ex_06()
		level04.Ex_07()
		level04.Ex_08()
		level04.Ex_09()
		level04.Ex_10()

	case "c1":
		challenge.PopChallenge()

	default:
		fmt.Println("Invalid exercise number. Please provide a number between 1 and 6.")
	}

}
