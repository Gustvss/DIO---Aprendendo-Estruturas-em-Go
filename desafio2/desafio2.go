//CRIAR APLICATIVO COM BASE EM ANTIGA BRINCADEIRA 'PIM PAM'
//SEMPRE QUE APARECER UM MÚLTIPLO DE 3, DEVE ESCREVER 'PIN'
//SEMPRE QUE APARECER UM MULTIPLO DE 5, DEVE FALAR 'PAN'
//IMPRIMIR NÚMEROS DE 1 A 100 E NOS CASOS CITADOS, NÃO DEVEM APARECER OS NÚMEROS, MAS SIM, STRINGS.

package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		}

	}

}
