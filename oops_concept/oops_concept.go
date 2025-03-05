package main

import (
	"fmt"
)

type Vahicle struct {
	make  string
	model string
	year  int
}

func (v *Vahicle) String() string {
	return fmt.Sprintf("%v %v (%v)", v.make, v.model, v.year)
}

type VehicleActions interface {
	Start()
	Stop()
	Drive()
}
type Car struct {
	NumberOfDoors int
	Vahicle
}

func (c Car) Start() {
	fmt.Println("Starting car", c.String())
}
func (c Car) Drive() {
	fmt.Println("Driving car", c.String())
}
func (c Car) Stop() {
	fmt.Println("Stoping car", c.String())
}

type Bike struct {
	HasSidecar string
	Vahicle
}

func (b Bike) Start() {
	fmt.Println("Starting bike", b.String())
}
func (b Bike) Drive() {
	fmt.Println("Driving bike", b.String())
}
func (b Bike) Stop() {
	fmt.Println("Stoping bike", b.String())
}

type Truck struct {
	PayloadCapacity string
	Vahicle
}

func (t Truck) Start() {
	fmt.Println("Starting Truck", t.String())
}
func (t Truck) Drive() {
	fmt.Println("Driving Truck", t.String())
}
func (t Truck) Stop() {
	fmt.Println("Stoping Truck", t.String())
}
func PerformVechleActions(v VehicleActions) {
	v.Start()
	v.Drive()
	v.Stop()
}

func main() {
	car := Car{Vahicle: Vahicle{"Toyata", "Corolla", 2022}, NumberOfDoors: 4}
	bike := Bike{Vahicle: Vahicle{"Harley Davidson", "Street 750", 2022},HasSidecar: "1"}

	PerformVechleActions(car)
	PerformVechleActions(bike)
}
