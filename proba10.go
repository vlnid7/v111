package main
import "fmt"

func main () {
	a:=9
	b:=4
	d:=a+b;	e:=a-b; f:=a*b; g:=a/b
	fmt.Println(d,e,f,g)
	l:=a&b; m:=a|b
	fmt.Println(l,m)
	var (
		h bool=true
		i bool=false
		j=h&&i; k=h||i
	)
	fmt.Print(j,k)
}
