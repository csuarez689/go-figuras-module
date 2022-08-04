package figuras

type Rectangulo struct {
	Ancho float64
	Alto  float64
}

func (c *Rectangulo) area() float64 {
	return c.Alto * c.Ancho
}

func (c *Rectangulo) perimetro() float64 {
	return 2*c.Alto + 2*c.Ancho
}
