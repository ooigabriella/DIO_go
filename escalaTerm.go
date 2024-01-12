package main

import (
	"fmt"
)

const ebulicaoK = 373.2

func main() {

	tempk := ebulicaoK        //variável do valor da temperatura
	var tempC = (tempk - 273) //conversão de K para C

	fmt.Printf("A temperatura de ebulição da água em °K = %g  e a temperatura de ebulição da água em °C = %g  ", tempk, tempC)

}
