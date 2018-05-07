package composite

import "fmt"

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type Athlete struct {
}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type SwimmerImpl struct {
}

func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimming!")
}

type CompositeSwimmer struct {
	Trainer
	Swimmer
}

func ComposeAll() {
	swimmer := CompositeSwimmer{
		&Athlete{},
		&SwimmerImpl{},
	}
	swimmer.Train()
	swimmer.Swim()
}
