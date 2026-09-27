package level01

import "fmt"

func Ex_01() {
	fmt.Println("<<ex 01>>")
	n := 42
	s := "James Bond"
	b := true
	fmt.Println(n, s, b)
	fmt.Println(n)
	fmt.Println(s)
	fmt.Println(b)
}

func Ex_02() {
	fmt.Println("<<ex 02>>")
	var n int
	var s string
	var b bool

	fmt.Printf("%v, %v, %v\n", n, s, b)
}

func Ex_03() {
	fmt.Println("<<ex 03>>")

	var x int = 42
	var y string = "James Bond"
	var z bool = true

	var s string = fmt.Sprintf("%v, %v, %v\n", x, y, z)
	fmt.Print(s)
}

func Ex_04() {
	fmt.Println("<<ex 04>>")

	type numero int
	var n4 numero

	fmt.Printf("%v, %T\n", n4, n4)
	n4 = 42
	fmt.Printf("%v, %T\n", n4, n4)
}

func Ex_05() {
	fmt.Println("<<ex 05>>")

	type numero1 int
	var n5 numero1
	var n6 int = 7

	fmt.Printf("%v, %T\n", n5, n5)
	n5 = 42
	fmt.Printf("%v, %T\n", n5, n5)
	n6 = int(n5)
	fmt.Printf("%v, %T\n", n6, n6)
}
