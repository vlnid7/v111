package main

import "fmt"

func main () {
	var b int
	for i:=19; i>=0; i-- {
		if i>9 {
			a := i * i
			fmt.Println("Квадрат числа", i, "равен", a)
		} else {
				b+=i
			}
	}
	fmt.Println("Сумма чисел от 0 до 9 равна", b)
}
