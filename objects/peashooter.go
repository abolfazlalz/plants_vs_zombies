package objects

import (
	"fmt"
	"image"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Peashooter struct with its own animation and state
type Peashooter struct {
	*Abstract
	frames           [2][3]*ebiten.Image // 3 idle frames, 2 shot frames
	currentFrame     *ebiten.Image
	idleFrameIndex   int
	shotFrameIndex   int
	isShot           bool
	lastUpdateTime   time.Time
	shotAnimationEnd time.Time
	mu               sync.Mutex
	width            int
	height           int
	lastUpdateLoc    time.Time

	bullet   []*Bullet
	lastShot time.Time
	health   int
}

func NewPeashooter() *Peashooter {
	return &Peashooter{
		Abstract: NewAbstract(),
		mu:       sync.Mutex{},
		width:    100,
		height:   100,
		health:   5,
	}
}

// Split the sprite sheet into frames (3x2 grid)
func (p *Peashooter) LoadFrames() {
	p.mu.Lock()
	defer p.mu.Unlock()

	spriteSheet := loadSpriteSheet("./assets/characters/peashooter.png")

	// Update frame dimensions to 100x100 pixels
	frameWidth, frameHeight := 100, 100

	// Loop through the sprite sheet and split into 6 frames (3x2 grid)
	for j := 0; j < 3; j++ {
		for i := 0; i < 2; i++ {
			rect := image.Rect((j+(i*3))*frameWidth, 0, (j+(i*3)+1)*frameWidth, frameHeight)
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
func (p *Peashooter) ShotAnimation() {
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

func (p *Peashooter) IsShot() bool {
	return p.isShot
}

func (p *Peashooter) Shot(shot bool) {
	p.isShot = shot
	if time.Since(p.lastShot) < time.Millisecond*500 {
		return
	}
	p.lastShot = time.Now()
	bullet := NewBullet(1, p.frames[1][2])
	bullet.SetPosition(p.x*150, p.y*100)
	p.bullet = append(p.bullet, bullet)
}

func (p *Peashooter) SetPosition(x, y int) {
	p.x, p.y = x*p.width, y*p.height
}

func (p *Peashooter) Frame() *ebiten.Image {
	return p.currentFrame
}

func (p *Peashooter) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	posX := float64(p.x)
	posY := float64(p.y)
	op.GeoM.Translate(posX, posY)
	screen.DrawImage(p.Frame(), op)

	for _, bullet := range p.bullet {
		bullet.Draw(screen)
	}
}

func (p *Peashooter) Animate() {
	// Change frame every 200ms
	// Handle idle or shot animation based on space key
	if !p.IsShot() {
		p.Idle()
	} else {
		p.ShotAnimation()
	}

	for _, bullet := range p.bullet {
		bullet.Animate()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if time.Since(p.lastUpdateLoc) > time.Millisecond*200 {
			p.lastUpdateLoc = time.Now()
			p.x += 1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.Shot(true)
	}
}

func (p *Peashooter) Collide(obj Object) {
	_, ok := obj.(*Zombie)
	if ok {
		fmt.Println("Zombie was eat me")
	}
}

func (p *Peashooter) Eat() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.health--
}
