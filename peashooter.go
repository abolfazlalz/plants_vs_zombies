package main

import (
	"image"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Peashooter struct with its own animation and state
type Peashooter struct {
	frames           [2][3]*ebiten.Image // 3 idle frames, 2 shot frames
	currentFrame     *ebiten.Image
	idleFrameIndex   int
	shotFrameIndex   int
	isShot           bool
	lastUpdateTime   time.Time
	shotAnimationEnd time.Time
	mu               sync.Mutex
}

// Split the sprite sheet into frames (3x2 grid)
func (p *Peashooter) LoadFrames() {
	p.mu.Lock()
	defer p.mu.Unlock()

	spriteSheet := loadSpriteSheet("./assets/characters/peashooter.png")

	// Update frame dimensions to 500x500 pixels
	frameWidth, frameHeight := 500, 500

	// Loop through the sprite sheet and split into 6 frames (3x2 grid)
	for j := 0; j < 3; j++ {
		for i := 0; i < 2; i++ {
			rect := image.Rect(j*frameWidth, i*frameHeight, (j+1)*frameWidth, (i+1)*frameHeight)
			p.frames[i][j] = spriteSheet.SubImage(rect).(*ebiten.Image)
		}
	}

	// Set the initial frame to the first idle frame
	p.currentFrame = p.frames[0][0]
}

// Idle method for Peashooter
func (p *Peashooter) Idle() {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Change frame every 200ms
	if time.Since(p.lastUpdateTime) > time.Millisecond*200 {
		p.lastUpdateTime = time.Now()
		p.idleFrameIndex = (p.idleFrameIndex + 1) % 3 // Loop between 0, 1, and 2
		p.currentFrame = p.frames[0][p.idleFrameIndex]
	}
}

// Shot method for Peashooter
func (p *Peashooter) Shot() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if time.Since(p.shotAnimationEnd) > time.Millisecond*200 { // Change frame every 150ms
		p.currentFrame = p.frames[1][p.shotFrameIndex]
		p.shotFrameIndex++
		if p.shotFrameIndex == 2 { // After the shot animation ends, reset
			p.shotFrameIndex = 0
			p.isShot = false // End shot animation and go back to idle
		}
		p.shotAnimationEnd = time.Now()
	}

}
