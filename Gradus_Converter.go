package main

import (
 "fmt"
)

func Get_celsius(temperature int) int{
   return (temperature - 32) * 5/9
}

func main() {
	var temperature_farengeit int
	fmt.Println("Изучение типов и функций ввода,вывода на языке Go by MK")
	fmt.Print("Введите температуру в Фарингейтах = ")
	fmt.Scanf("%d",&temperature_farengeit)
	fmt.Println("Температура в цельсиях = ",Get_celsius(temperature_farengeit),"°C")
}
