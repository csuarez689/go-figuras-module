package figuras

import "fmt"

type geom interface {
	area() float64
	perimetro() float64
}

func Area(g geom) float64 {
	return g.area()
}

func Perimetro(g geom) float64 {
	return g.perimetro()
}

func Medidas(g geom) {
	fmt.Print("\n==============================\n")
	fmt.Println("Medidas:", g)
	fmt.Println("Área:", Area(g))
	fmt.Println("Perímetro:", Perimetro(g))
	fmt.Print("==============================\n")

}
