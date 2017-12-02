package behaviors

import (
	"robosim/common"
)

// Literal TODO: comment
type Literal Behavior

// CreateLiteral TODO: comment
func CreateLiteral(name string, val interface{}) *Literal {
	var behavior Literal

	(*Behavior)(&behavior).Init(name, nil, nil)
	behavior.output = val

	return &behavior
}

////////////////////////////////////////////////////////////////////////////////

// NumLiteral TODO: comment
type NumLiteral Literal

// CreateNumLiteral TODO: comment
func CreateNumLiteral(name string, val float64) *NumLiteral {
	return (*NumLiteral)(CreateLiteral(name, val))
}

// Output TODO: comment
func (b *NumLiteral) Output() float64 {
	output := (*Behavior)(b).Output()
	return output.(float64)
}

////////////////////////////////////////////////////////////////////////////////

// IntLiteral TODO: comment
type IntLiteral Literal

// Output TODO: comment
func (b *IntLiteral) Output() int {
	output := (*Behavior)(b).Output()
	return output.(int)
}

// CreateIntLiteral TODO: comment
func CreateIntLiteral(name string, val int) *IntLiteral {
	return (*IntLiteral)(CreateLiteral(name, val))
}

////////////////////////////////////////////////////////////////////////////////

// VectorLiteral TODO: comment
type VectorLiteral Behavior

// Output TODO: comment
func (b *VectorLiteral) Output() common.Vector {
	output := (*Behavior)(b).Output()
	return output.(common.Vector)
}

// CreateVectorLiteral TODO: comment
func CreateVectorLiteral(
	name string, val common.Vector) *VectorLiteral {

	return (*VectorLiteral)(CreateLiteral(name, val))
}

////////////////////////////////////////////////////////////////////////////////

// VectorListLiteral TODO: comment
type VectorListLiteral Behavior

// Output TODO: comment
func (b *VectorListLiteral) Output() []common.Vector {
	output := (*Behavior)(b).Output()
	return output.([]common.Vector)
}

// CreateVectorListLiteral TODO: comment
func CreateVectorListLiteral(
	name string, val []common.Vector) *VectorListLiteral {

	return (*VectorListLiteral)(CreateLiteral(name, val))
}
