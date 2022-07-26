package main

import (
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float32
	keliling() float32
}

type hitung3d interface {
	volume() float32
}

type Hitung interface {
	hitung2d
	hitung3d
}

type Kubus struct {
	sisi float32
}

func (k *Kubus) volume() float32 {
	return float32(math.Pow(float64(k.sisi), 3))
	// return math.Pow(k.sisi, 3)
}

func (k *Kubus) luas() float32 {
	return float32(math.Pow(float64(k.sisi), 3*6))
}

func (k *Kubus) keliling() float32 {
	return float32(math.Pow(float64(k.sisi), 12))
}
func main() {
	var bangunRuang Hitung = &Kubus{4}

	fmt.Println("======== Kubus")
	fmt.Println("Volume		: ", bangunRuang.volume())
	fmt.Println("Luas		: ", bangunRuang.luas())
	fmt.Println("Keliling	: ", bangunRuang.keliling())
}