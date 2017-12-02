package behaviors

import (
	"robosim/robot/robot_interfaces"
)

// MoveRobot TODO: comment
type MoveRobot struct {
	Behavior
	Movement  VectorGetter
	BaseSpeed Float64Getter
	MaxSpeed  Float64Getter
}

// CreateMoveRobot TODO: comment
func CreateMoveRobot(name string,
	movement VectorGetter, baseSpeed Float64Getter,
	maxSpeed Float64Getter) *MoveRobot {

	var behavior MoveRobot

	outputFunc := func() {
		computeMoveRobot(&behavior)
	}

	behavior.Behavior.Init(name, nil, outputFunc)

	return &behavior
}

func computeMoveRobot(behavior *MoveRobot) {
	movement := behavior.Movement.Output()
	baseSpeed := behavior.BaseSpeed.Output()
	maxSpeed := behavior.MaxSpeed.Output()

	movement = movement.Mult(baseSpeed)
	if movement.Length() > maxSpeed {
		movement = movement.Unit().Mult(maxSpeed)
	}

	robot_interfaces.TheRobot.Move(movement)
}
