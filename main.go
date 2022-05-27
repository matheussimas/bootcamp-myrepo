package main

import "fmt"

// eu consigo criar varias interfaces dentro de uma struct, e subdividr ainda mais criando um array de interfaces
type ListaDeFuncs struct {
	Data []interface{}
}

func main() {
	l := ListaDeFuncs{}
	l.Data = append(l.Data, 1)
	l.Data = append(l.Data, "oi!")
	l.Data = append(l.Data, false)

	fmt.Println(l.Data)

}
