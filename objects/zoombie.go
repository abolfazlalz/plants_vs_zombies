package objects

import (
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Zombie struct {
	*Abstract
	width, height  int
	lastUpdateTime time.Time
	frames         []*ebiten.Image // 2 walk frames
	currentFrame   *ebiten.Image
	idleFrameIndex int
	speed          int
	eating         bool
}

func NewZombie() *Zombie {
	spriteSheet := loadSpriteSheet("./assets/characters/zombie.png")
	frameWidth, frameHeight := 100, 100

	frames := make([]*ebiten.Image, 2)

	// Loop through the sprite sheet and split into 6 frames (2x1 grid)
	for i := 0; i < 2; i++ {
		rect := image.Rect(i*frameWidth, 0, (i+1)*frameWidth, frameHeight)
		frames[i] = spriteSheet.SubImage(rect).(*ebiten.Image)
	}

	return &Zombie{
		Abstract:       NewAbstract(),
		width:          100,
		height:         100,
		lastUpdateTime: time.Now(),
		frames:         frames,
		currentFrame:   frames[0],
		idleFrameIndex: 0,
		speed:          5,
		eating:         false,
	}
}

func (z *Zombie) Animate() {
	z.mu.Lock()
	defer z.mu.Unlock()
	if time.Since(z.lastUpdateTime) < time.Millisecond*200 { // Change frame every 150ms
		return
	}
	z.lastUpdateTime = time.Now()
	if !z.eating {
		z.x -= z.speed

		z.idleFrameIndex = (z.idleFrameIndex + 1) % 2 // Loop between 0, 1, and 2
		z.currentFrame = z.frames[z.idleFrameIndex]
	}
}

func (z *Zombie) SetPosition(x, y int) {
	z.x, z.y = x*z.width, y*z.height-20
}

func (z *Zombie) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	posX := float64(z.x)
	posY := float64(z.y)
	op.GeoM.Scale(1.2, 1.2)
	op.GeoM.Translate(posX, posY)
	screen.DrawImage(z.currentFrame, op)
}
