package behaviors

import "robosim/common"

// MoveTo TODO: comment
type MoveTo struct {
	Behavior
	Target VectorGetter
}

// Output TODO: comment
func (b *MoveTo) Output() common.Vector {
	output := b.Behavior.Output()
	return output.(common.Vector)
}

// CreateMoveTo TODO: comment
func CreateMoveTo(name string, target VectorGetter) *MoveTo {

	var behavior MoveTo

	outputFunc := func() {
		behavior.output = behavior.Target.Output().Unit()
	}

	behavior.Behavior.Init(name, nil, outputFunc)
	behavior.Target = target

	return &behavior
}
