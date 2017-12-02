package behaviors

import (
	"robosim/common"
	"robosim/robot/robot_interfaces"
)

// GetPosition TODO: comment
type GetPosition Behavior

// Output TODO: comment
func (b *GetPosition) Output() common.Placement {
	output := (*Behavior)(b).Output()
	return output.(common.Placement)
}

// CreateGetPosition TODO: comment
func CreateGetPosition(name string) *GetPosition {
	var behavior GetPosition

	outputFunc := func() {
		behavior.output = robot_interfaces.TheRobot.GetPosition()
	}

	(*Behavior)(&behavior).Init(name, nil, outputFunc)

	return &behavior
}
