package part

import (
	"constants"
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
)

func GetDistance(pos1, pos2 box2d.Vec2) (box2d.Vec2, float64) {
	distVec := pos2.Sub(pos1)
	distVec.X = distVec.X / constants.SCALE
	distVec.Y = distVec.Y / constants.SCALE

	return distVec, distVec.Length() / constants.SCALE
}
