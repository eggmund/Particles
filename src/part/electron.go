package part

import (
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
	"constants"
	"id"
)

func NewElectron(newIDnum int, x, y, vx, vy float64) *Particle {
	p := Particle{
		ID:     id.ID{Num: newIDnum, Type: "e"},
		Radius: constants.ElecR,
		Chr:    constants.ElecChr,
		Anti:   false,
		clr:    constants.ElecClr,
	}
	p.Set(&box2d.Vec2{float64(constants.ElecDiam), float64(constants.ElecDiam)}, float64(constants.ElecM))
	p.Body.Position = box2d.Vec2{x, y}
	p.Body.Velocity = box2d.Vec2{vx * constants.SCALE * constants.SCALE, vy * constants.SCALE * constants.SCALE}
	World.AddBody(&p.Body)
	return &p
}
