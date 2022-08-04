package main

import (
	"fmt"
)

type slovo struct {
	c float64
	z float64
	y float64
}

func Sqrt(x, c, z, y float64) float64 {

	perv := c
	perv -= (perv*perv - x) / (2 * perv)
	vtr := z
	vtr -= (vtr*vtr - x) / (2 * vtr)
	thrd := y
	thrd -= (thrd*thrd - x) / (2 * thrd)
	if perv > x && perv < vtr && perv < thrd {
		fmt.Printf("Вариант %.2f наиболее близкий к числу %.0f - %.2f\nВторой вариант - %.2f\nТретий вариант - %.2f", c, x, perv, vtr, thrd) //Первый вариант ближе всех
	} else if vtr > x && vtr < thrd && vtr < perv {
		fmt.Printf("Вариант %.2f наиболее близкий к числу %.0f - %.2f\nПервый вариант - %.2f\nТретий вариант - %.2f", z, x, vtr, perv, thrd) //Второй Вариант ближе всех

	} else {
		fmt.Printf("Вариант %.2f наиболее близкий к числу %.0f - %.2f\nПервый вариант - %.2f\nВторой вариант - %.2f", y, x, thrd, perv, vtr) //Третий вариант ближе всех
	}

	return x
}

func main() {
	chislo := 0.00
	c1 := 0.00
	z1 := 0.00
	y1 := 0.00
	fmt.Println("Введите к какому числу вы хотите проверить близость")
	fmt.Scanln(&chislo)
	fmt.Println("\nВведите первое значение")
	fmt.Scanln(&c1)
	fmt.Println("\nВведите второе значение")
	fmt.Scanln(&z1)
	fmt.Println("\nВведите третье значение")
	fmt.Scanln(&y1)
	var slv = slovo{c: c1, z: c1, y: y1}
	fmt.Println(Sqrt(chislo, slv.c, slv.z, slv.y))
}
