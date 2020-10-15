package figuras

import "math"

type FiguraMultiple struct {
	Figuras []Figura
}

func (fm *FiguraMultiple) Area() float64 {
	var area float64
	for _, f := range fm.Figuras {
		area += f.Area()
	}
	return area
}

type Figura interface {
	Area() float64
}

type Circulo struct {
	Radio float64
}

func (c *Circulo) Area() float64 {
	return c.Radio * c.Radio * math.Pi
}

type Rectangulo struct {
	Base float64
	Altura float64
}

func (r *Rectangulo) Area() float64 {
	return r.Base * r.Altura
}