package behaviors

import (
	"robosim/common"
)

// GlobalToEgocentric TODO: comment
type GlobalToEgocentric struct {
	Behavior
	RobotPos  PlacementGetter
	GlobalPos VectorGetter
}

// Output TODO: comment
func (b *GlobalToEgocentric) Output() common.Vector {
	output := b.Behavior.Output()
	return output.(common.Vector)
}

// CreateGlobalToEgocentric TODO: comment
func CreateGlobalToEgocentric(name string,
	robotPos PlacementGetter, globalPos VectorGetter) *GlobalToEgocentric {

	var behavior GlobalToEgocentric

	outputFunc := func() {
		computeGlobalToEgocentric(&behavior)
	}

	behavior.Behavior.Init(name, nil, outputFunc)
	behavior.RobotPos = robotPos
	behavior.GlobalPos = globalPos

	return &behavior
}

func computeGlobalToEgocentric(behavior *GlobalToEgocentric) {
	robotPos := behavior.RobotPos.Output()
	globalPos := behavior.GlobalPos.Output()

	output := globalPos.Sub(robotPos.Pos)
	output = output.RotateZ(common.DegToRad(robotPos.Heading * -1))

	behavior.output = output
}
