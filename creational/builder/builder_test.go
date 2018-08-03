package builder

import (
	"testing"
	"fmt"
)

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	CarBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(CarBuilder.SetWheels())
	manufacturingComplex.Construct()

	car := CarBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on car must be 4 but there were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure must be 'Car' but was %s\n", car.Structure)
	}

	BikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(BikeBuilder)
	manufacturingComplex.Construct()

	bike := BikeBuilder.GetVehicle()
	bike.Seats = 1

	if bike.Wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n", bike.Wheels)
	}
	fmt.Println(bike.Seats)
}
