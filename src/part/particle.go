// Base of particle.
package part

import (
	"constants"
	"github.com/gen2brain/raylib-go/raylib"
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
	"id"
)

var (
	Particles []*Particle
	World     *box2d.World // Given by main when started.
	DebugV    bool
	DebugF    bool
)

type Particle struct {
	box2d.Body // Base class is box2d body.
	ID         id.ID
	Radius     float64
	Chr        float64 // Charge
	Anti       bool    // True if antiparticle
	clr        rl.Color
}

func (self *Particle) Draw() {
	rl.DrawCircleLines(int32(self.Body.Position.X), int32(self.Body.Position.Y), float32(self.Radius), self.clr)

	if DebugV {
		self.debugVel()
	}

	if DebugF {
		self.debugForce()
	}
}

func (self *Particle) Update(done chan<- bool) {
	self.Body.Force = self.getEMForce()
	done<- true
}

func (self *Particle) getEMForce() box2d.Vec2 {
	var fTotal = box2d.Vec2{0, 0}
	for i := 0; i < len(Particles); i++ {
		if Particles[i].ID.Num != self.ID.Num {
			distVec, dist := GetDistance(self.Body.Position, Particles[i].Position)
			f := -constants.ChrMult * (float64(self.Chr*Particles[i].Chr) / (dist * dist))
			fTotal.X += f * distVec.X
			fTotal.Y += f * distVec.Y
		}
	}
	return fTotal
}

func NewProton(newIDnum int, x, y, vx, vy float64) *Particle {
	p := Particle{
		ID:     id.ID{Num: newIDnum, Type: "p"},
		Radius: constants.ProtonR,
		Chr:    constants.ProtonChr,
		Anti:   false,
		clr:    constants.ProtonClr,
	}
	p.Set(&box2d.Vec2{float64(constants.ProtonDiam), float64(constants.ProtonDiam)}, float64(constants.ProtonM))
	p.Body.Position = box2d.Vec2{x, y}
	p.Body.Velocity = box2d.Vec2{vx * constants.SCALE * constants.SCALE, vy * constants.SCALE * constants.SCALE}
	println(p.Body.Velocity.X)
	World.AddBody(&p.Body)
	return &p
}

func NewAntiProton(newIDnum int, x, y, vx, vy float64) *Particle {
	p := Particle{
		ID:     id.ID{Num: newIDnum, Type: "_p"},
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
