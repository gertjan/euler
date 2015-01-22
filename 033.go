package euler

import "fmt"

func P033() string {
	denproduct := 1
	nomproduct := 1

	for i := 1; i < 10; i++ {
		for den := 1; den < i; den++ {
			for nom := 1; nom < den; nom++ {
				if (nom*10+i)*den == nom*(i*10+den) {
					denproduct *= den
					nomproduct *= nom
				}
			}
		}
	}
	denproduct /= gcd(nomproduct, denproduct)
	return fmt.Sprint(denproduct)
}
