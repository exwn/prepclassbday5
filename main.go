package main

import "fmt"

func main() {
	//fmt.Println("hello wirld")
	// loopWithRange()
	// oddEven()
	addToArray()
}

func loopWithRange() {
	inputPertama := 4
	inputKedua := 8

	for i := inputPertama; i <= inputKedua; i++ {
		fmt.Println(i)
	}
}

func oddEven() {
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
		//fmt.Println(odd)
	}
}

func addToArray() {
	var stuff = [...]string{"Meja", "Buku", "Topi", "Baju", "Kayu"}
	fmt.Println("stuff \t:", stuff)

}
