package kennel

import (
	"github.com/cucumber/messages-go/v10"
)

type Suite interface {
	Step(interface{}, interface{})
	BeforeScenario(func(*messages.Pickle))
	BeforeSuite(func())
	AfterSuite(func())
}

type Stepper interface {
	StepUp(Suite)
}

var steppers []Stepper

func Register(s Stepper) {
	steppers = append(steppers, s)
}

func StepUp(s Suite) {
	for _, stepper := range steppers {
		stepper.StepUp(s)
	}
}
