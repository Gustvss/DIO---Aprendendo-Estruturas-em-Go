//CRIE UM CÓDIGO QUE EXIBA TODOS OS NÚMEROS ENTRE 1 E 100 QUE SÃO DIVISÍVEIS POR 3.

package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
