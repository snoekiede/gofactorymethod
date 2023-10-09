package main

import (
	"errors"
	"fmt"
)

type VehicleBrand int

const (
	BrandVolkswagen VehicleBrand = iota
	BrandRenault
)

type VehicleFactory interface {
	createCar(color string) AbstractCar
	createBike(numberOfWheels int16) AbstractBike
}

type AbstractCar interface {
	GetDescription() string
}

type AbstractBike interface {
	GetDescription() string
}

type VolkswagenCar struct {
	make  string
	color string
}

func (c VolkswagenCar) GetDescription() string {
	return fmt.Sprintf("Make: %s, color: %s", c.make, c.color)
}

type VolkswagenBike struct {
	make           string
	numberOfWheels int16
}

func (b VolkswagenBike) GetDescription() string {
	return fmt.Sprintf("Make: %s, numberOfWheels: %d", b.make, b.numberOfWheels)
}

type RenaultCar struct {
	make  string
	color string
}

func (r RenaultCar) GetDescription() string {
	return fmt.Sprintf("Make: %s, color: %s", r.make, r.color)
}

type RenaultBike struct {
	make           string
	numberOfWheels int16
}

func (b RenaultBike) GetDescription() string {
	return fmt.Sprintf("Make: %s, numberOfWheels: %d", b.make, b.numberOfWheels)
}

type VolkswagenFactory struct{}

func (f VolkswagenFactory) createCar(color string) AbstractCar {
	return VolkswagenCar{color: color, make: "Volkswagen"}
}

func (f VolkswagenFactory) createBike(numberOfWheels int16) AbstractBike {
	return VolkswagenBike{numberOfWheels: numberOfWheels, make: "Volkswagen"}
}

type RenaultFactory struct{}

func (r RenaultFactory) createCar(color string) AbstractCar {
	return RenaultCar{color: color, make: "Renault"}
}

func (r RenaultFactory) createBike(numberOfWheels int16) AbstractBike {
	return RenaultBike{numberOfWheels: numberOfWheels, make: "Renault"}
}

type IVehicleCreator interface {
	createCar(brand VehicleBrand, color string) (AbstractCar, error)
	createBike(brand VehicleBrand, numberOfWheels int16) (AbstractBike, error)
}

type VehicleCreator struct{}

func (vc VehicleCreator) createCar(brand VehicleBrand, color string) (AbstractCar, error) {
	switch brand {
	case BrandVolkswagen:
		{
			factory := VolkswagenFactory{}
			return factory.createCar(color), nil
		}
	case BrandRenault:
		{
			factory := RenaultFactory{}
			return factory.createCar(color), nil
		}
	default:
		return nil, errors.New("unknown brand")
	}
}

func (vc VehicleCreator) createBike(brand VehicleBrand, numberOfWheels int16) (AbstractBike, error) {
	switch brand {
	case BrandVolkswagen:
		{
			factory := VolkswagenFactory{}
			return factory.createBike(numberOfWheels), nil
		}
	case BrandRenault:
		{
			factory := RenaultFactory{}
			return factory.createBike(numberOfWheels), nil
		}
	default:
		return nil, errors.New("unknown brand")
	}
}

func main() {
	var creator IVehicleCreator
	creator = VehicleCreator{}
	car, carError := creator.createCar(BrandVolkswagen, "red")
	if carError != nil {
		fmt.Println("Could not create Volkswagen car")
		return
	}
	fmt.Println(car.GetDescription())
	nextBike, nextError := creator.createBike(BrandRenault, 3)
	if nextError != nil {
		fmt.Println("Could not create car: ", nextError)
		return
	}
	fmt.Println(nextBike.GetDescription())
}
