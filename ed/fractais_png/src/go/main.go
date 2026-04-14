package main

import (
	"fmt"
	"math/rand"
)

func ri(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}

func triangulo(pen *Pen, tamanho float64, nivel int ){
	if nivel == 0{
		return
	}
	for range 3{
		pen.Walk(tamanho)
		pen.Left(120)
	}

	metade := tamanho/2

	triangulo(pen, metade, nivel-1)

	pen.Walk(metade)
	triangulo(pen, metade, nivel-1)

	pen.Left(120)
    pen.Walk(metade)
    pen.Right(120)
    triangulo(pen, metade, nivel-1)

	pen.Right(120)
    pen.Walk(metade)
    pen.Left(120)

}


func circulos(pen *Pen, raio float64, nivel int){
	if nivel == 0 {
		return
	}
	pen.DrawCircle(raio)

	for range 6 {
		pen.Right(60)

		pen.Up()
		pen.Walk(raio)
		pen.Down()

		pen.Right(90)

		circulos(pen, raio*0.33, nivel-1)
		
		pen.Up()
		pen.Left(90)
		pen.Walk(-raio)
		pen.Down()
	}
}

func arvere(pen *Pen, dist float64) {
	if dist < 10 {
		if ri(0, 50) == 0 {
			pen.SetRGB(255, 0, 0)
			pen.FillCircle(10)
		}
		return
	}
	ang_dir := ri(10, 40)
	ang_esq := ri(10, 40)

	pen.SetLineWidth(dist / 5)
	pen.SetRGB(0, 0, 0)
	pen.Walk(dist)
	pen.Right(ang_dir)
	arvere(pen, dist*(ri(80, 85)/100))
	pen.Left(ang_dir + ang_esq)
	arvere(pen, dist*(ri(80, 85)/100))
	pen.Right(ang_esq)
	pen.SetRGB(0, 0, 0)
	pen.Walk(-dist)
}

func main() {
	pen := NewPen(800, 800)
	pen.SetHeading(60)
	pen.SetPosition(400, 400)
	triangulo(pen, 350,5)
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}

