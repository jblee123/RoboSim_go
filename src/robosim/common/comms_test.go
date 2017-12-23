package common

import (
	"testing"

	"image/color"
)

func TestMarshallUnmarshallAliveMsg(t *testing.T) {
	msg1 := AliveMsg{
		1,
		2,
		Placement{Vector{6, 7, 8}, 20},
		color.RGBA{25, 26, 27, 28},
		50, 51, 52,
	}

	buf := msg1.Marshall(nil)

	var msg2 AliveMsg
	msg2.Unmarshall(buf)

	if msg1.MsgType != msg2.MsgType {
		t.Errorf("AliveMsg marshall/unmarshall (%v): got %v", msg1, msg2)
	}
}

func TestMarshallUnmarshallStartMsg(t *testing.T) {
	msg1 := StartMsg{1}

	buf := msg1.Marshall(nil)

	var msg2 StartMsg
	msg2.Unmarshall(buf)

	if msg1.MsgType != msg2.MsgType {
		t.Errorf("StartMsg marshall/unmarshall (%v): got %v", msg1, msg2)
	}
}

func TestMarshallUnmarshallRequestPositionMsg(t *testing.T) {
	msg1 := RequestPositionMsg{1, 2}

	buf := msg1.Marshall(nil)

	var msg2 RequestPositionMsg
	msg2.Unmarshall(buf)

	if msg1.MsgType != msg2.MsgType {
		t.Errorf("RequestPositionMsg marshall/unmarshall (%v): got %v", msg1, msg2)
	}
}

func TestMarshallUnmarshallPositionMsg(t *testing.T) {
	msg1 := PositionMsg{
		1,
		Placement{Vector{6, 7, 8}, 20},
	}

	buf := msg1.Marshall(nil)

	var msg2 PositionMsg
	msg2.Unmarshall(buf)

	if msg1.MsgType != msg2.MsgType {
		t.Errorf("PositionMsg marshall/unmarshall (%v): got %v", msg1, msg2)
	}
}

func TestMarshallUnmarshallKillMsg(t *testing.T) {
	msg1 := KillMsg{1}

	buf := msg1.Marshall(nil)

	var msg2 KillMsg
	msg2.Unmarshall(buf)

	if msg1.MsgType != msg2.MsgType {
		t.Errorf("KillMsg marshall/unmarshall (%v): got %v", msg1, msg2)
	}
}

func TestMarshallUnmarshallRobotDyingMsg(t *testing.T) {
	msg1 := RobotDyingMsg{1, 2}

	buf := msg1.Marshall(nil)

	var msg2 RobotDyingMsg
	msg2.Unmarshall(buf)

	if msg1.MsgType != msg2.MsgType {
		t.Errorf("RobotDyingMsg marshall/unmarshall (%v): got %v", msg1, msg2)
	}
}

func TestMarshallUnmarshallGetObstaclesMsg(t *testing.T) {
	msg1 := GetObstaclesMsg{1, 2}

	buf := msg1.Marshall(nil)

	var msg2 GetObstaclesMsg
	msg2.Unmarshall(buf)

	if msg1.MsgType != msg2.MsgType {
		t.Errorf("GetObstaclesMsg marshall/unmarshall (%v): got %v", msg1, msg2)
	}
}
