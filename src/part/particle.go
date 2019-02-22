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
	self.checkInBounds()
	done<- true
}

func (self *Particle) checkInBounds() {
	upBoundX := float64(constants.SCREEN_W) - self.Radius
	upBoundY := float64(constants.SCREEN_H) - self.Radius

	if self.Position.X < self.Radius {
		self.Velocity.X = -self.Velocity.X
	}
	if self.Position.X > upBoundX {
		self.Velocity.X = -self.Velocity.X
	}

	if self.Position.Y < self.Radius {
		self.Velocity.Y = -self.Velocity.Y
	}
	if self.Position.Y > upBoundY {
		self.Velocity.Y = -self.Velocity.Y
	}
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
