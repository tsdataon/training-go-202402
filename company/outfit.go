package company

import "fmt"

type Outfit interface {
	Pakai()
}

type Kemeja struct {
}

func (k Kemeja) Pakai() {
	fmt.Println("Pakai Kemeja")
}

type Kaos struct {
}

func (k Kaos) Pakai() {
	fmt.Println("Pakai Kaos")
}

type Jacket struct {
}

func (j Jacket) Pakai() {
	fmt.Println("Pakai Jacket")
}

type TV struct{}

func (t TV) Play() {
	fmt.Println("Pakai Play")
}
