package area

import "math"

const Pi = 3.1416

// Circ é responsável em calcular a área de um círculo
func Circ(raio float64) float64{
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável em calcular a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TriaguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}