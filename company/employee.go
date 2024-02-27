package company

import "fmt"

type Employee struct {
	Name     string
	Position string
	Outfit   Outfit
}

// public
func (e Employee) GetEmployeeName() string {
	return e.Name
}

func (e Employee) GetOutfit() {
	e.Outfit.Pakai()
}

// private
func (e Employee) present() {
	fmt.Println("Present")
}
