package part

import (
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
	"constants"
	"id"
)

func NewAntiProton(newIDnum int, x, y, vx, vy float64) *Particle {
	p := Particle{
		ID:     id.ID{Num: newIDnum, Type: "p"},
		Radius: constants.ProtonR,
		Chr:    -constants.ProtonChr,
		Anti:   true,
		clr:    constants.AntiProtonClr,
	}
	p.Set(&box2d.Vec2{float64(constants.ProtonDiam), float64(constants.ProtonDiam)}, float64(constants.ProtonM))
	p.Body.Position = box2d.Vec2{x, y}
	p.Body.Velocity = box2d.Vec2{vx * constants.SCALE * constants.SCALE, vy * constants.SCALE * constants.SCALE}
	World.AddBody(&p.Body)
	return &p
}
