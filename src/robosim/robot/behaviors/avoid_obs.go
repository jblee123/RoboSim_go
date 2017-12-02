package behaviors

import (
	"robosim/common"
)

// AvoidObs TODO: comment
type AvoidObs struct {
	Behavior
	ObsList           VectorListGetter
	SafetyMargin      Float64Getter
	SphereOfInfluence Float64Getter
}

// Output TODO: comment
func (b *AvoidObs) Output() common.Vector {
	output := b.Behavior.Output()
	return output.(common.Vector)
}

// CreateAvoidObs TODO: comment
func CreateAvoidObs(name string,
	obsList VectorListGetter, safetyMargin Float64Getter,
	sphereOfInfluence Float64Getter) *AvoidObs {

	var behavior AvoidObs

	outputFunc := func() {
		computeAvoidObs(&behavior)
	}

	behavior.Behavior.Init(name, nil, outputFunc)
	behavior.ObsList = obsList
	behavior.SafetyMargin = safetyMargin
	behavior.SphereOfInfluence = sphereOfInfluence

	return &behavior
}

func computeAvoidObs(behavior *AvoidObs) {
	soi := behavior.SphereOfInfluence.Output()
	output := common.Vector{X: 0, Y: 0, Z: 0}

	for _, obs := range behavior.ObsList.Output() {
		length := obs.Length()
		if length < soi {
			obs = obs.Mult(100000)
		}
		output = output.Add(obs.Mult(-1))
	}

	behavior.output = output
}
