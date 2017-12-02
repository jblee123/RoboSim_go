package behaviors

import (
	"robosim/common"
)

// SumVectors TODO: comment
type SumVectors struct {
	Behavior
	Vectors []VectorGetter
	Weights []Float64Getter
}

// Output TODO: comment
func (b *SumVectors) Output() common.Vector {
	output := b.Behavior.Output()
	return output.(common.Vector)
}

// CreateSumVectors TODO: comment
func CreateSumVectors(name string,
	vectors []VectorGetter, weights []Float64Getter) *SumVectors {

	var behavior SumVectors

	outputFunc := func() {
		computeSumVectors(&behavior)
	}

	behavior.Behavior.Init(name, nil, outputFunc)
	behavior.Vectors = vectors
	behavior.Weights = weights

	return &behavior
}

func computeSumVectors(behavior *SumVectors) {
	output := common.Vector{X: 0, Y: 0, Z: 0}
	count := common.MinI(len(behavior.Vectors), len(behavior.Weights))
	for i := 0; i < count; count++ {
		vec := behavior.Vectors[i].Output()
		vec = vec.Mult(behavior.Weights[i].Output())
		output = output.Add(vec)
	}

	behavior.output = output
}
