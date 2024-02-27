package main

import (
	"errors"
	"fmt"
	"gotraining3/company"
	"sync"
	"time"
)

func main() {
	// fmt.Println("Hello, World!")

	// // variable declaration
	// var a int = 10
	// b := 40
	// var c, d int = 10, 20

	// // phi := 3.14

	// // changing the value of b
	// b = 100
	// fmt.Println(a, b, c, d)

	// // constant declaration
	// /* Ini comment */
	// const x = 3.14
	// fmt.Println(x)

	// // write var with all data type in go
	// var (
	// 	name   string  = "hello"
	// 	age    int     = -25
	// 	amount int64   = 20000000
	// 	phi    float64 = 3.14
	// 	isTrue bool    = true

	// 	// 0+
	// 	positiveInt uint = 100
	// )
	// fmt.Println(name, age, amount, phi, isTrue, positiveInt)

	// if isTrue && ((age < 0) || len(name) > 0) {
	// 	fmt.Println("ini true")
	// } else {
	// 	fmt.Println("ini false")
	// }

	// // switch case
	// switch {
	// case age < 13:
	// 	fmt.Println("ini masih anak-anak")
	// case age >= 12 && age < 20:
	// 	fmt.Println("ini remaja")
	// default:
	// 	fmt.Println("gak tau dah")
	// }

	// // array
	// var arr [5]int
	// arr[0] = 100
	// arr[1] = 200
	// arr[2] = 300
	// arr[3] = 400
	// arr[4] = 500
	// fmt.Println(arr)
	// arr[4] = 1000
	// fmt.Println(arr)

	// var buahBuahan = [3]string{"apel", "jeruk", "mangga"}
	// fmt.Println(buahBuahan)

	// // slice
	// var sliceBuah = []string{"apel", "jeruk", "mangga", "pepaya"}
	// fmt.Println(sliceBuah[2:])

	// var bebas = []any{1, "dua", 3.14, true}
	// fmt.Println(bebas)

	// // map
	// // gak berurutan
	// var person = map[string]string{
	// 	"name": "john",
	// 	"age":  "25",
	// }
	// fmt.Println(person)
	// person["name"] = "doe"
	// fmt.Println(person["name"])

	// // for loop
	// for key, value := range person {
	// 	fmt.Println(key, value) // name john, age 25
	// }
	// i := 0
	// for i < 5 {
	// 	fmt.Println(i) //  0 1 2 3 4
	// 	i++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i) // 0 1 2 3 4
	// }

	// for {
	// 	if i < 5 {
	// 		fmt.Println(i) // 0 1 2 3 4
	// 		i++
	// 	} else {
	// 		break
	// 	}
	// }

	// // function
	// greeting, name, err := sayHello("John", 25)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// } else {
	// 	fmt.Println(greeting, name) // Hello John
	// }

	// panic
	// fmt.Println(sayHello(0, 25))

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Error occured", r)
	// 	} else {
	// 		fmt.Println("Application running perfectly")
	// 	}
	// }()

	// panic
	// a := 0
	// fmt.Println(23 / a)

	// sayHello2() // Hello, end

	// sum := func(a, b int) int {
	// 	return a + b
	// }
	// fmt.Println(sum(3, 4)) // 7

	// method
	// agus := person{name: "Agus", age: 25}
	// agus.sayHello() // Hello Agus

	// agus.age = 30
	// fmt.Println(agus.age) // 30

	// pointer
	// copyanAgus := agus // pass by value (copy)
	// agus.name = "Agus Ganteng"
	// fmt.Println(copyanAgus.name) // Agus
	// fmt.Println(agus.name)       // Agus Ganteng

	// var agusShare *person = &agus // pass by reference
	// agus.name = "Agus Ganteng Banget"
	// fmt.Println(agusShare.name)

	// agusShare.name = "Agus Ganteng Sekali"
	// fmt.Println(agus.name)

	// panic
	// status(nil)

	// Goroutine()
	// Mutex()
	// Mutex()
	// Mutex()
	// Mutex()

	budi := company.Employee{
		Name:     "Budi",
		Position: "Manager",
	}
	fmt.Println(budi.GetEmployeeName())

	// budi.Outfit = company.Kaos{}
	// budi.Outfit = company.Jacket{}
	budi.Outfit = company.Kemeja{}
	budi.GetOutfit()
}

// func sayHello(name any, age int) (string, string, error) {
// 	if age < 0 {
// 		return "", "", errors.New("umur tidak boleh minus")
// 	}
// 	return "Hello", name.(string), nil // Hello John
// }

func sayHello2() {
	// stack = Last in First Out
	defer fmt.Println("end")
	defer fmt.Println("world")
	fmt.Println("Hello")
}

type person struct {
	name   string
	age    int
	status string
}

// method
func (p person) sayHello() {
	fmt.Println("Hello", p.name)
}

// mutasi data
func status(person *person) error {
	if person == nil {
		return errors.New("person tidak boleh nil")
	}

	if person.age < 20 {
		person.status = "belum menikah"
	} else {
		person.status = "menikah"
	}
	return nil
}

// sample func with goroutine processes
func Goroutine() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("World")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Satu")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Dua")
		wg.Done()
	}()

	fmt.Println("Hello")
	wg.Wait()
}

// Mutex
func Mutex() {
	mu := sync.Mutex{}

	wg := sync.WaitGroup{}
	var x = 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				x++
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(x)

	// saldo 1000 / 1500 / 700
	// topup 500  -- 1500  -- 1000 + 500 -- 1500 selesai
	// tarik 800  -- 700 -- |||| 1500 - 800 -- 700
	// 700
}
