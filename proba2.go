package main
import "fmt"

func main () {
	const (
		a float32 = 25.7
		c = iota+0.7
		d

		f = 10
		e = iota-0.8
		h,l = iota,iota
		
		b = "Проверка"
		
)
		fmt.Println(a,c,d,b)
		fmt.Println(f,e,)
		fmt.Println(h,l)
		fmt.Println(b)
}
