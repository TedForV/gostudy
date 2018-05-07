package absfactory

import (
	"errors"
	"fmt"
)

type Motorbike interface {
	GetMotorbikeType() int
}

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct {
}

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}

type SportMotorbike struct {
}

func (s *SportMotorbike) NumWheels() int {
	return 2
}

func (s *SportMotorbike) NumSeats() int {
	return 1
}

func (s *SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}

type CruiseMotorbike struct {
}

func (s *CruiseMotorbike) NumWheels() int {
	return 2
}

func (s *CruiseMotorbike) NumSeats() int {
	return 2
}

func (s *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}
