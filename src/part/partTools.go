package part

import (
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
)

func GetDistance(pos1, pos2 box2d.Vec2) (box2d.Vec2, float64) {
	distVec := pos2.Sub(pos1)
	return distVec, distVec.Length()
}
