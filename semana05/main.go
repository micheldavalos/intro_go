package main

import (
	"fmt"
	"./figuras"
)

func sumAreas(figs ...figuras.Figura) float64 {
	area_total := 0.0

	for _, f := range figs {
		area_total += f.Area()
	}
	return area_total
}

func main()  {
	c01 := figuras.Circulo{Radio: 6}
	fmt.Println(c01.Area())

	r01 := figuras.Rectangulo{Altura: 10, Base: 2}
	fmt.Println(r01.Area())

	fmt.Println(sumAreas(&c01, &r01))

	fm := figuras.FiguraMultiple{
		Figuras: []figuras.Figura{
			&c01,
			&r01,
			// &figuras.Circulo{2},
		},
	}
	fmt.Println(fm.Area())

	// c02 := Circulo{radio: 10}
	// c03 := Circulo{15}
	// c04 := new(Circulo)
	// c05 := &Circulo{100}


	// fmt.Println(c01)
	// fmt.Println(c02)
	// fmt.Println(c03)
	// fmt.Println(c04)
	// fmt.Println(c05)
}