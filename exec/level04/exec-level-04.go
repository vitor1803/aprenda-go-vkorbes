package level04

import "fmt"

func Ex_01() {
	fmt.Println("<<ex 01>>")
	var list [5]int = [5]int{10, 20, 30, 40, 50}
	for i, v := range list {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
	fmt.Printf("%T\n", list)
}

func Ex_02() {
	fmt.Println("<<ex 02>>")
	var list []int = []int{11, 21, 31, 41, 51, 61, 71, 81, 91, 101}
	for i := 0; i < len(list); i++ {
		fmt.Printf("Index: %d, Value: %d\n", i, list[i])
	}
	fmt.Printf("%T\n", list)
}

func Ex_03() {
	fmt.Println("<<ex 03>>")
	var list []int = []int{11, 21, 31, 41, 51, 61, 71, 81, 91, 101}
	fmt.Printf("%v\n", list[:3])
	fmt.Printf("%v\n", list[4:])
	fmt.Printf("%v\n", list[1:7])
	fmt.Printf("%v\n", list[2:9])
	fmt.Printf("%v\n", list[2:len(list)-1])
}

func Ex_04() {
	fmt.Println("<<ex 04>>")
	var list []int = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	list = append(list, 52)
	fmt.Printf("%v\n", list)
	list = append(list, 53, 54, 55)
	fmt.Printf("%v\n", list)
	var list2 []int = []int{56, 57, 58, 59, 60}
	list = append(list, list2...)
	fmt.Printf("%v\n", list)
}

func Ex_05() {
	fmt.Println("<<ex 05>>")
	var list []int = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(list)
	fmt.Println(len(list))
	fmt.Println(cap(list))

	list = append(list[:3], list[6:]...)
	fmt.Println(list)
	fmt.Println(len(list))
	fmt.Println(cap(list))

}

func Ex_06() {
	fmt.Println("<<ex 06>>")

	var estados []string = make([]string, 26, 26)
	fmt.Printf("%T %d %d\n", estados, cap(estados), len(estados))
	estados = []string{
		"Acre",
		"Alagoas",
		"Amapá",
		"Amazonas",
		"Bahia",
		"Ceará",
		"Espírito Santo",
		"Goiás",
		"Maranhão",
		"Mato Grosso",
		"Mato Grosso do Sul",
		"Minas Gerais",
		"Pará",
		"Paraíba",
		"Paraná",
		"Pernambuco",
		"Piauí",
		"Rio de Janeiro",
		"Rio Grande do Norte",
		"Rio Grande do Sul",
		"Rondônia",
		"Roraima",
		"Santa Catarina",
		"São Paulo",
		"Sergipe",
		"Tocantins",
	}

	fmt.Printf("%T %d %d\n", estados, cap(estados), len(estados))
	for i := 0; i < len(estados); i++ {
		fmt.Println(i, estados[i])
	}
}

func Ex_07() {
	fmt.Println("<<ex 07>>")

	var pessoas [][]string = [][]string{}
	pessoas = append(pessoas, []string{"João", "Silva", "Futebol"})
	pessoas = append(pessoas, []string{"Ana", "Julia", "Pintura"})
	pessoas = append(pessoas, []string{"Maria", "Santos", "Artes"})
	fmt.Println(pessoas)
}

func Ex_08() {
	fmt.Println("<<ex 08>>")
	var pessoas map[string][]string = map[string][]string{
		"João": {"Silva", "Futebol"},
		"Ana":  {"Julia", "Pintura"},
	}
	pessoas["Maria"] = []string{"Santos", "Artes"}
	fmt.Println(pessoas)
	for i, _ := range pessoas {
		fmt.Printf("%#v\n", i)
	}
}

func Ex_09() {
	fmt.Println("<<ex 09>>")
	var pessoas map[string][]string = map[string][]string{
		"João": {"Silva", "Futebol"},
		"Ana":  {"Julia", "Pintura"},
	}
	pessoas["Maria"] = []string{"Santos", "Artes"}
	pessoas["Adiantar"] = []string{"Segunda", "Entrada"}
	fmt.Println(pessoas)
}

func Ex_10() {
	fmt.Println("<<ex 10>>")
	var pessoas map[string][]string = map[string][]string{
		"João": {"Silva", "Futebol"},
		"Ana":  {"Julia", "Pintura"},
	}
	pessoas["Maria"] = []string{"Santos", "Artes"}
	pessoas["Adiantar"] = []string{"Segunda", "Entrada"}

	delete(pessoas, "Ana")
	fmt.Println(pessoas)
}
