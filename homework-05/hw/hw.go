package hw

import (
	"fmt"
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.
type Geom struct {
	X1, Y1, X2, Y2 float64
}

// Конструктор для создания объекта Geom
func NewGeom(x1, y1, x2, y2 float64) *Geom {
	return &Geom{
		X1: x1,
		Y1: y1,
		X2: x2,
		Y2: y2,
	}
}

func (g *Geom) CalculateDistance() (distance float64, err error) {
	if g.X1 < 0 || g.X2 < 0 || g.Y1 < 0 || g.Y2 < 0 {
		fmt.Println("Координаты не могут быть меньше нуля")
		return -1, fmt.Errorf("координаты не могут быть меньше нуля")
	}
	
	distance = math.Sqrt(math.Pow(g.X2-g.X1, 2) + math.Pow(g.Y2-g.Y1, 2))
	return distance, nil
}

/*
	Вообще здесь можно было бы не использовать ООП подход, 
	а просто создать функцию для вычисления дистанции.
	Также вместо простого возврата -1 при ошибке целесообразно 
	в добавок создать ошибку и возвращать её.
*/