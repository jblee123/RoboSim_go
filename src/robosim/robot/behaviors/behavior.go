package behaviors

import (
	"fmt"

	"robosim/common"
	"robosim/robot/robot_interfaces"
)

// Float64Getter TODO: comment
type Float64Getter interface {
	Output() float64
}

// IntGetter TODO: comment
type IntGetter interface {
	Output() int
}

// PlacementGetter TODO: comment
type PlacementGetter interface {
	Output() common.Placement
}

// VectorGetter TODO: comment
type VectorGetter interface {
	Output() common.Vector
}

// VectorListGetter TODO: comment
type VectorListGetter interface {
	Output() []common.Vector
}

////////////////////////////////////////////////////////////////////////////////

// Output TODO: comment
type Output interface{}

// ActivateFunc TODO: comment
type ActivateFunc func()

// OutputFunc TODO: comment
type OutputFunc func()

var nextBehaviorAnonID = 1

// Behavior TODO: comment
type Behavior struct {
	name         string
	output       Output
	lastCycle    int
	activateFunc ActivateFunc
	outputFunc   OutputFunc
}

// Init TODO: comment
func (b *Behavior) Init(name string,
	activateFunc ActivateFunc, outputFunc OutputFunc) {

	b.name = name
	if b.name == "" {
		b.name = fmt.Sprintf("AN_%d", nextBehaviorAnonID)
		nextBehaviorAnonID++
	}

	b.lastCycle = -1

	b.activateFunc = activateFunc
	b.outputFunc = outputFunc
}

// Name TODO: comment
func (b *Behavior) Name() string {
	return b.name
}

// Output TODO: comment
func (b *Behavior) Output() Output {
	cycleNum := robot_interfaces.TheRobot.CycleNum()

	if (cycleNum-b.lastCycle > 2) && (b.activateFunc != nil) {
		b.activateFunc()
		b.activateFunc = nil
	}

	if b.lastCycle < cycleNum {
		if b.outputFunc != nil {
			b.outputFunc()
		}
		b.lastCycle = cycleNum
	}

	return b.output
}
