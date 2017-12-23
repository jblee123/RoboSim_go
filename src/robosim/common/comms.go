package common

import (
	"bytes"

	"encoding/binary"

	"image/color"
)

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

const DefConsolePort = 50000

var defByteOrder = binary.BigEndian

////////////////////////////////////////////////////////////////////////////////

type AliveMsg struct {
	MsgType       int32
	Id            int32
	Place         Placement
	Color         color.RGBA
	MaxVel        float64
	MaxAngularVel float64
	Radius        float64
}

func (msg *AliveMsg) Marshall(dest *bytes.Buffer) *bytes.Buffer {
	if dest == nil {
		dest = new(bytes.Buffer)
	}

	dest.Reset()
	binary.Write(dest, defByteOrder, msg.MsgType)
	binary.Write(dest, defByteOrder, msg.Id)
	binary.Write(dest, defByteOrder, msg.Place.Pos.X)
	binary.Write(dest, defByteOrder, msg.Place.Pos.Y)
	binary.Write(dest, defByteOrder, msg.Place.Pos.Z)
	binary.Write(dest, defByteOrder, msg.Place.Heading)
	binary.Write(dest, defByteOrder, msg.Color.R)
	binary.Write(dest, defByteOrder, msg.Color.G)
	binary.Write(dest, defByteOrder, msg.Color.B)
	binary.Write(dest, defByteOrder, msg.Color.A)
	binary.Write(dest, defByteOrder, msg.MaxVel)
	binary.Write(dest, defByteOrder, msg.MaxAngularVel)
	binary.Write(dest, defByteOrder, msg.Radius)

	return dest
}

func (msg *AliveMsg) Unmarshall(src *bytes.Buffer) {
	binary.Read(src, defByteOrder, &msg.MsgType)
	binary.Read(src, defByteOrder, &msg.Id)
	binary.Read(src, defByteOrder, &msg.Place.Pos.X)
	binary.Read(src, defByteOrder, &msg.Place.Pos.Y)
	binary.Read(src, defByteOrder, &msg.Place.Pos.Z)
	binary.Read(src, defByteOrder, &msg.Place.Heading)
	binary.Read(src, defByteOrder, &msg.Color.R)
	binary.Read(src, defByteOrder, &msg.Color.G)
	binary.Read(src, defByteOrder, &msg.Color.B)
	binary.Read(src, defByteOrder, &msg.Color.A)
	binary.Read(src, defByteOrder, &msg.MaxVel)
	binary.Read(src, defByteOrder, &msg.MaxAngularVel)
	binary.Read(src, defByteOrder, &msg.Radius)
}

////////////////////////////////////////////////////////////////////////////////

type StartMsg struct {
	MsgType int32
}

func (msg *StartMsg) Marshall(dest *bytes.Buffer) *bytes.Buffer {
	if dest == nil {
		dest = new(bytes.Buffer)
	}

	dest.Reset()
	binary.Write(dest, defByteOrder, msg.MsgType)

	return dest
}

func (msg *StartMsg) Unmarshall(src *bytes.Buffer) {
	binary.Read(src, defByteOrder, &msg.MsgType)
}

////////////////////////////////////////////////////////////////////////////////

type RequestPositionMsg struct {
	MsgType int32
	Id      int32
}

func (msg *RequestPositionMsg) Marshall(dest *bytes.Buffer) *bytes.Buffer {
	if dest == nil {
		dest = new(bytes.Buffer)
	}

	dest.Reset()
	binary.Write(dest, defByteOrder, msg.MsgType)
	binary.Write(dest, defByteOrder, msg.Id)

	return dest
}

func (msg *RequestPositionMsg) Unmarshall(src *bytes.Buffer) {
	binary.Read(src, defByteOrder, &msg.MsgType)
	binary.Read(src, defByteOrder, &msg.Id)
}

////////////////////////////////////////////////////////////////////////////////

type PositionMsg struct {
	MsgType int32
	Place   Placement
}

func (msg *PositionMsg) Marshall(dest *bytes.Buffer) *bytes.Buffer {
	if dest == nil {
		dest = new(bytes.Buffer)
	}

	dest.Reset()
	binary.Write(dest, defByteOrder, msg.MsgType)
	binary.Write(dest, defByteOrder, msg.Place.Pos.X)
	binary.Write(dest, defByteOrder, msg.Place.Pos.Y)
	binary.Write(dest, defByteOrder, msg.Place.Pos.Z)
	binary.Write(dest, defByteOrder, msg.Place.Heading)

	return dest
}

func (msg *PositionMsg) Unmarshall(src *bytes.Buffer) {
	binary.Read(src, defByteOrder, &msg.MsgType)
	binary.Read(src, defByteOrder, &msg.Place.Pos.X)
	binary.Read(src, defByteOrder, &msg.Place.Pos.Y)
	binary.Read(src, defByteOrder, &msg.Place.Pos.Z)
	binary.Read(src, defByteOrder, &msg.Place.Heading)
}

////////////////////////////////////////////////////////////////////////////////

type KillMsg struct {
	MsgType int32
}

func (msg *KillMsg) Marshall(dest *bytes.Buffer) *bytes.Buffer {
	if dest == nil {
		dest = new(bytes.Buffer)
	}

	dest.Reset()
	binary.Write(dest, defByteOrder, msg.MsgType)

	return dest
}

func (msg *KillMsg) Unmarshall(src *bytes.Buffer) {
	binary.Read(src, defByteOrder, &msg.MsgType)
}

////////////////////////////////////////////////////////////////////////////////

type RobotDyingMsg struct {
	MsgType int32
	Id      int32
}

func (msg *RobotDyingMsg) Marshall(dest *bytes.Buffer) *bytes.Buffer {
	if dest == nil {
		dest = new(bytes.Buffer)
	}

	dest.Reset()
	binary.Write(dest, defByteOrder, msg.MsgType)
	binary.Write(dest, defByteOrder, msg.Id)

	return dest
}

func (msg *RobotDyingMsg) Unmarshall(src *bytes.Buffer) {
	binary.Read(src, defByteOrder, &msg.MsgType)
	binary.Read(src, defByteOrder, &msg.Id)
}

////////////////////////////////////////////////////////////////////////////////

type GetObstaclesMsg struct {
	MsgType int32
	Id      int32
}

func (msg *GetObstaclesMsg) Marshall(dest *bytes.Buffer) *bytes.Buffer {
	if dest == nil {
		dest = new(bytes.Buffer)
	}

	dest.Reset()
	binary.Write(dest, defByteOrder, msg.MsgType)
	binary.Write(dest, defByteOrder, msg.Id)

	return dest
}

func (msg *GetObstaclesMsg) Unmarshall(src *bytes.Buffer) {
	binary.Read(src, defByteOrder, &msg.MsgType)
	binary.Read(src, defByteOrder, &msg.Id)
}

////////////////////////////////////////////////////////////////////////////////

type ObsReadingsMsg struct {
	MsgType int32
	Readings []struct{ X, Y int}
}

func (msg *ObsReadingsMsg) Marshall(dest *bytes.Buffer) *bytes.Buffer {
	if dest == nil {
		dest = new(bytes.Buffer)
	}

	dest.Reset()
	binary.Write(dest, defByteOrder, msg.MsgType)
	binary.Write(dest, defByteOrder, int32(len(msg.Readings)))
	for _, reading := range msg.Readings {
		binary.Write(dest, defByteOrder, reading.X)
		binary.Write(dest, defByteOrder, reading.Y)
	}

	return dest
}

func (msg *ObsReadingsMsg) Unmarshall(src *bytes.Buffer) {
	var readingCount int32

	binary.Read(src, defByteOrder, &msg.MsgType)
	binary.Read(src, defByteOrder, &readingCount)
}
