package main

import (
  "github.com/gen2brain/raylib-go/raylib"
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
	"strconv"

	"part"
	"constants"
)


var (
	World *box2d.World = box2d.NewWorld(box2d.Vec2{0, 0}, 100)
	done = make(chan bool) // For updating
)


func Draw() {
	for i := 0; i < len(part.Particles); i++ {
		part.Particles[i].Draw()
	}
}

func Update(dt float64) {
	World.Step(dt)

	for i := 0; i < len(part.Particles); i++ {
		go part.Particles[i].Update(done)
	}

	for i := 0; i < len(part.Particles); i++ {
		<-done
	}
}

func checkInputs() {
	if rl.IsKeyPressed(rl.KeyV) {
		part.DebugV = !part.DebugV
	}
	if rl.IsKeyPressed(rl.KeyF) {
		part.DebugF = !part.DebugF
	}
}

func main() {
  rl.InitWindow(int32(constants.SCREEN_W), int32(constants.SCREEN_H), "Particles")
  rl.SetTargetFPS(60)

	part.World = World
	part.DebugV = false
	part.DebugF = false

	part.Particles = append(part.Particles,	part.NewProton(len(part.Particles), 520, 200, 0, 0))
	part.Particles = append(part.Particles,	part.NewAntiProton(len(part.Particles), 500, 300, 0, 0))
	part.Particles = append(part.Particles,	part.NewProton(len(part.Particles), 600, 500, 0, 0))

	part.Particles = append(part.Particles,	part.NewProton(len(part.Particles), 300, 200, 0, 0))
	part.Particles = append(part.Particles,	part.NewAntiProton(len(part.Particles), 350, 300, 0, 0))
	part.Particles = append(part.Particles,	part.NewProton(len(part.Particles), 200, 500, 0, 0))

	for i := 0; i < 50; i++ {
		if i % 2 == 0 {
			part.Particles = append(part.Particles,	part.NewAntiProton(len(part.Particles), float64((i*10)+200), 400, 0, 0))
		} else {
			part.Particles = append(part.Particles,	part.NewProton(len(part.Particles), float64((i*10)+200), 400, 0, 0))
		}
	}

  for !rl.WindowShouldClose() {
		Update(float64(rl.GetFrameTime()))
		checkInputs()

    rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawFPS(10, 10)
		rl.DrawText("Particle count: "+strconv.Itoa(len(part.Particles)), 10, 40, 12, rl.RayWhite)

		Draw()

		rl.EndDrawing()
  }

  rl.CloseWindow()
}
