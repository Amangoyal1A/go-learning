package main

import (
	"fmt"
	"sync"
)

/* -------------------------------------------------------------------------- */
/*                                 single instace                             */
/* -------------------------------------------------------------------------- */

type singleton struct {
	timestamp int64
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{timestamp: 123456789}
	})
	return instance
}

// func main() {
// 	a := GetInstance()
// 	b := GetInstance()
// 	fmt.Println(a == b) // true
// }

/* -------------------------------------------------------------------------- */
/*                                 factory pattern                            */
/* -------------------------------------------------------------------------- */

type Vehicle interface { Drive() }

type Car struct{}
func (c Car) Drive() { fmt.Println("Driving a car") }

type Truck struct{}
func (t Truck) Drive() { fmt.Println("Driving a truck") }

func VehicleFactory(t string) Vehicle {
    if t == "car" { return Car{} }
    if t == "truck" { return Truck{} }
    return nil
}

func main() {
    v := VehicleFactory("car")
    v.Drive()
}