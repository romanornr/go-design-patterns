package abstract_factory

import (
	"github.com/go-errors/errors"
	"fmt"
)

const (
	LuxuaryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct {}

func (c *CarFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuaryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamiliarCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n"))
	}
}
