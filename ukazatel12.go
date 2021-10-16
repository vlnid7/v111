package main
import "fmt"

func main () {
	var b float32 = 40.7
	var p *float32
	p=&b
	fmt.Println("Указатель p =",p)
	fmt.Println("Значение b = ",b)
	b=b+b
	c:=*p
	fmt.Println(c)
}