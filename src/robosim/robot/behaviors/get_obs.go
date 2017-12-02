package behaviors

import (
	"robosim/common"
	"robosim/robot/robot_interfaces"
)

// GetObs TODO: comment
type GetObs Behavior

// Output TODO: comment
func (b *GetObs) Output() []common.Vector {
	output := (*Behavior)(b).Output()
	return output.([]common.Vector)
}

// CreateGetObs TODO: comment
func CreateGetObs(name string) *GetObs {
	var behavior GetObs

	outputFunc := func() {
		behavior.output = robot_interfaces.TheRobot.GetObsReadings()
	}

	(*Behavior)(&behavior).Init(name, nil, outputFunc)

	return &behavior
}
