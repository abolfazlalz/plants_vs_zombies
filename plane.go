package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Plane struct {
	frames [2][3]*ebiten.Image // 3 idle frames, 2 shot frames
}

// Split the sprite sheet into frames (3x1 grid)
func (p *Plane) LoadFrames() {
	p.mu.Lock()
	defer p.mu.Unlock()

	spriteSheet := loadSpriteSheet("./assets/characters/plane.png")

	// Update frame dimensions to 500x500 pixels
	frameWidth, frameHeight := 500, 500

	// Loop through the sprite sheet and split into 6 frames (3x2 grid)
	for j := 0; j < 3; j++ {
		rect := image.Rect(j*frameWidth, i*frameHeight, (j+1)*frameWidth, (i+1)*frameHeight)
		p.frames[i][j] = spriteSheet.SubImage(rect).(*ebiten.Image)
	}

	// Set the initial frame to the first idle frame
	p.currentFrame = p.frames[0][0]
}
