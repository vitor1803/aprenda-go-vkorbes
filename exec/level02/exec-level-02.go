package level02

import "fmt"

func Ex_01() {
	fmt.Println("<<ex 01>>")
	var n int = 30
	fmt.Printf("%#b, %d, %#X\n", n, n, n)
}

func Ex_02() {
	fmt.Println("<<ex 02>>")
	var b1 bool = 10 == 5
	var b2 bool = 10 != 5
	var b3 bool = 10 > 5
	var b4 bool = 10 < 5
	var b5 bool = 10 >= 5
	var b6 bool = 10 <= 5

	fmt.Printf("%v, %v, %v, %v, %v, %v\n", b1, b2, b3, b4, b5, b6)

}

func Ex_03() {
	fmt.Println("<<ex 03>>")
	const x int = 10
	const y = 5
	fmt.Printf("x = %d, y = %d\n", x, y)
}

func Ex_04() {
	fmt.Println("<<ex 04>>")
	var n int = 200
	fmt.Printf("n: %d, %#x, %#b\n", n, n, n)
	var m = n << 1
	fmt.Printf("m: %d, %#x, %#b\n", m, m, m)

}

func Ex_05() {
	fmt.Println("<<ex 05>>")
	var s string = ` Hello,\n World!`
	fmt.Println(s)
}

func Ex_06() {
	const (
		currentYear = 2026 + iota
		next1Year
		next2Year
		next3Year
		next4Year
	)
	fmt.Println("<<ex 06>>")
	fmt.Println("Current Year:", currentYear)
	fmt.Println("Next Year:", next1Year)
	fmt.Println("Next Year:", next2Year)
	fmt.Println("Next Year:", next3Year)
	fmt.Println("Next Year:", next4Year)
}
