package robot_interfaces

import (
    "robosim/common"
)

type RobotInterface interface {
    CycleNum() int
    GetPosition() common.Placement
    Move(movement common.Vector)
    GetObsReadings() []common.Vector
}

var TheRobot RobotInterface
