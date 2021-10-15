package main

import (
	"fmt"
)

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
	zero:=0
	switch zero {
	case 0:
		for k:=-5; k<6; k++ {
			if k<0 {
				fmt.Println("k - отрицательное число")
				fmt.Println(k, "\t")
			}
			if k==0 {
				fmt.Println("k равно", k, ", продолжаем")
				continue
			}
			if k>0 {
				break
			}
		}
		fmt.Println("k больше нуля")
	case 4:
		fallthrough
	default:
		fmt.Println("По умолчанию")


	}

}
