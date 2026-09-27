package level03

import "fmt"

func Ex_01() {
	fmt.Println("<<ex 01>>")
	for i := 1; i <= 10_000; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func Ex_02() {
	fmt.Println("<<ex 02>>")
	for il := rune(65); il <= 90; il++ {
		for in := 0; in < 3; in++ {
			fmt.Printf("%d %#U\n", (in + 1), il)
		}
		fmt.Println()
	}
}

func Ex_03() {
	fmt.Println("<<ex 03>>")
	birth_year := rune(1998)
	current_year := rune(2026)

	year := birth_year

	for year <= current_year {
		fmt.Println(year)
		year++
	}
}

func Ex_04() {
	fmt.Println("<<ex 04>>")
	birth_year := rune(1998)
	current_year := rune(2026)

	year := birth_year

	for {
		fmt.Println(year)
		if year == current_year {
			break
		}
		year++
	}
}

func Ex_05() {
	fmt.Println("<<ex 05>>")
	for i := 10; i <= 100; i++ {
		fmt.Printf("%d %% 4 = %d\n", i, i%4)
	}
}

func Ex_06() {
	fmt.Println("<<ex 06>>")
	var i int = 2
	if i > 1 {
		fmt.Println("i is greater than 1")
	}
}

func Ex_07() {
	fmt.Println("<<ex 07>>")
	var i int = 20
	if i > 5 {
		fmt.Println("i is greater than 5")
	} else if i > 1 {
		fmt.Println("i is greater than 1")
	} else {
		fmt.Println("i is less than or equal to 1")
	}
}

func Ex_08() {
	fmt.Println("<<ex 08>>")
	var i int = -3
	switch {
	case i > 5:
		fmt.Println("i is greater than 5")
	case i > 1:
		fmt.Println("i is greater than 1")
	default:
		fmt.Println("i is less than or equal to 1")
	}
}

func Ex_09() {
	fmt.Println("<<ex 09>>")
	var sport string = "tennis"
	switch sport {
	case "soccer":
		fmt.Println("Soccer is a popular sport worldwide.")
	case "basketball":
		fmt.Println("Basketball is played with a ball and a hoop.")
	case "tennis":
		fmt.Println("Tennis is played on a court with rackets.")
	default:
		fmt.Println("Unknown sport.")
	}
}

func Ex_10() {
	fmt.Println("<<ex 10>>")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
