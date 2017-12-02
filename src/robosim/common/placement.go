package common

import "fmt"

type Placement struct {
	Pos     Vector
	Heading Degrees
}

func (p Placement) String() string {
	return fmt.Sprintf("[%v, %.02f]", p.Pos, p.Heading)
}
