package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
	"strconv"

	"constants"
	"part"
)

var (
	timeMult float64 = 0.0000000000000001 / constants.SCALE // time multiplier
	coreNum int
	paused bool = false
	World *box2d.World = box2d.NewWorld(box2d.Vec2{0, 0}, 20)
	done               = make(chan bool) // For updating
)

func Draw() {
	for i := 0; i < len(part.Particles); i++ {
		part.Particles[i].Draw()
	}
}

func Update(dt float64) {
	dt = dt * timeMult
	World.Step(dt)

	for i := 0; i < len(part.Particles); i++ {
		go part.Particles[i].Update(done)
	}

	for i := 0; i < len(part.Particles); i++ {
		<-done
	}
}

func checkInputs(mouseDown, paused *bool, startMX, startMY, endMX, endMY *int32) {
	if (rl.IsMouseButtonDown(0) || rl.IsMouseButtonDown(1)) && !(*mouseDown) {
		*mouseDown = true
		*startMX, *startMY = rl.GetMouseX(), rl.GetMouseY()
	}

	if rl.IsMouseButtonReleased(0) || rl.IsMouseButtonReleased(1) {
		*mouseDown = false
		*endMX, *endMY = rl.GetMouseX(), rl.GetMouseY()
		if rl.IsMouseButtonReleased(0) {
			part.Particles = append(part.Particles, part.NewProton(len(part.Particles), float64(*startMX), float64(*startMY), float64(*startMX-*endMX)*2, float64(*startMY-*endMY)*2))
		} else {
			part.Particles = append(part.Particles, part.NewAntiProton(len(part.Particles), float64(*startMX), float64(*startMY), float64(*startMX-*endMX)*2, float64(*startMY-*endMY)*2))
		}
	}

	if *mouseDown {
		rl.DrawLineEx(rl.NewVector2(float32(*startMX), float32(*startMY)), rl.NewVector2(float32(rl.GetMouseX()), float32(rl.GetMouseY())), 2, rl.Lime)
	}

	if rl.IsKeyPressed(rl.KeyP) {
		*paused = !(*paused)
	}
	if rl.IsKeyPressed(rl.KeyV) {
		part.DebugV = !part.DebugV
	}
	if rl.IsKeyPressed(rl.KeyF) {
		part.DebugF = !part.DebugF
	}
	if rl.IsKeyPressed(rl.KeyR) {
		reset()
	}

	if rl.IsKeyPressed(rl.KeyS) {
		makeSqr(float64(rl.GetMouseX()), float64(rl.GetMouseY()), 4, 4)
	}
}

func makeSqr(offsetX, offsetY float64, w, h int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			println(i, j)
			if (i+j)%2 == 0 {
				part.Particles = append(part.Particles, part.NewAntiProton(len(part.Particles), ((float64(i)*constants.ProtonDiam)+offsetX), (float64(j)*constants.ProtonDiam)+offsetY, 0, 0))
			} else {
				part.Particles = append(part.Particles, part.NewProton(len(part.Particles), ((float64(i)*constants.ProtonDiam)+offsetX), (float64(j)*constants.ProtonDiam)+offsetY, 0, 0))
			}
		}
	}
}

func loadTest() {
	part.Particles = append(part.Particles, part.NewProton(len(part.Particles), 520, 200, 0, 0))
	part.Particles = append(part.Particles, part.NewAntiProton(len(part.Particles), 500, 300, 0, 0))
	part.Particles = append(part.Particles, part.NewProton(len(part.Particles), 600, 500, 0, 0))

	part.Particles = append(part.Particles, part.NewProton(len(part.Particles), 300, 200, 0, 0))
	part.Particles = append(part.Particles, part.NewAntiProton(len(part.Particles), 350, 300, 0, 0))
	part.Particles = append(part.Particles, part.NewProton(len(part.Particles), 200, 500, 0, 0))
	makeSqr(200, 200, 10, 8)
}

func reset() {
	World.Clear()
	part.Particles = []*part.Particle{}
}

func drawPaused() {
	rl.DrawRectangle(int32(constants.SCREEN_W-40), 8, 10, 40, rl.RayWhite)
	rl.DrawRectangle(int32(constants.SCREEN_W-20), 8, 10, 40, rl.RayWhite)
}

func main() {
	rl.InitWindow(int32(constants.SCREEN_W), int32(constants.SCREEN_H), "Particles")
	rl.SetTargetFPS(60)

	var (
		mouseDown bool  = false
		startMX   int32 = 0
		startMY   int32 = 0
		endMX     int32 = 0
		endMY     int32 = 0
	)

	part.World = World
	part.DebugV = false
	part.DebugF = false

	println("SCALE: ", constants.SCALE)
	println("Time multiplier: ", timeMult)

	loadTest()

	//part.Particles = append(part.Particles, part.NewProton(len(part.Particles), 0, 400, 700, 0))

	//part.Particles = append(part.Particles, part.NewProton(len(part.Particles), 800, 400, -700, 40))
	//part.Particles = append(part.Particles, part.NewElectron(len(part.Particles), 800, 700, 0, 0))

	for !rl.WindowShouldClose() {
		if !paused {
			Update(float64(rl.GetFrameTime()))
		}
		checkInputs(&mouseDown, &paused, &startMX, &startMY, &endMX, &endMY)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawFPS(10, 10)
		rl.DrawText("Particle count: "+strconv.Itoa(len(part.Particles)), 10, 40, 12, rl.RayWhite)
		rl.DrawText(strconv.FormatFloat(1/timeMult, 'e', 6, 64)+" times slower than real time.", 560, 10, 20, rl.RayWhite)

		if paused {
			drawPaused()
		}

		Draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
