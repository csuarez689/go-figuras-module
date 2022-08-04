package figuras

import "fmt"

type Geom interface {
	area() float64
	perimetro() float64
}

func Area(g Geom) float64 {
	return g.area()
}

func Perimetro(g Geom) float64 {
	return g.perimetro()
}

func Medidas(g Geom) {
	fmt.Print("\n==============================\n")
	fmt.Println("Medidas:", g)
	fmt.Println("Área:", Area(g))
	fmt.Println("Perímetro:", Perimetro(g))
	fmt.Print("==============================\n")

}
