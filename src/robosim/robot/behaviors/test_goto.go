package behaviors

import (
	"robosim/common"
)

// GoTo TODO: comment
type GoTo struct {
	Behavior
	moveRobot *MoveRobot
}

// CreateGoTo TODO: comment
func CreateGoTo(name string) *GoTo {
	var behavior GoTo

	outputFunc := func() {
		behavior.moveRobot.Output()
	}

	behavior.Behavior.Init(name, nil, outputFunc)

	behavior.moveRobot = CreateMoveRobot("MoveRobot",
		CreateSumVectors("movement",
			[]VectorGetter{
				CreateMoveTo("move_to",
					CreateGlobalToEgocentric("g2e",
						CreateGetPosition("get_pos"),
						CreateVectorLiteral("target_pos",
							common.Vector{X: 49, Y: 49, Z: 0}))),
				CreateAvoidObs("avoid_obs",
					CreateGetObs("get_obs"),
					CreateNumLiteral("safety_margin", 1.5),
					CreateNumLiteral("sphere_of_influence", 5)),
				CreateWander("wander",
					CreateIntLiteral("persistence", 10)),
			},
			[]Float64Getter{
				CreateNumLiteral("move_weight", 1),
				CreateNumLiteral("avoid_obs_weight", 1),
				CreateNumLiteral("wander_weight", 1),
			}),
		CreateNumLiteral("base_speed", 1),
		CreateNumLiteral("max_speed", 1))

	return &behavior
}
