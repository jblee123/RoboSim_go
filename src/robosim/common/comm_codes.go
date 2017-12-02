package common

type RobotMsgId int

const (
	Alive           = RobotMsgId(1)
	Start           = RobotMsgId(2)
	RequestPosition = RobotMsgId(3)
	Position        = RobotMsgId(4)
	Kill            = RobotMsgId(5)
	RobotDying      = RobotMsgId(6)
	GetObstacles    = RobotMsgId(7)
	ObsReadings     = RobotMsgId(8)
	Pause           = RobotMsgId(9)
	Move            = RobotMsgId(10)
	Spin            = RobotMsgId(11)
)
