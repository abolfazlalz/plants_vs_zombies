package objects

import (
	"image"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type Plane struct {
	frames []*ebiten.Image
	mu     sync.Mutex
}

func NewPlane() *Plane {
	return &Plane{
		mu:     sync.Mutex{},
		frames: make([]*ebiten.Image, 3),
	}
}

// Split the sprite sheet into frames (3x1 grid)
func (p *Plane) LoadFrames() {
	p.mu.Lock()
	defer p.mu.Unlock()

	spriteSheet := loadSpriteSheet("./assets/plane.png")

	// Update frame dimensions to 500x500 pixels
	frameWidth, frameHeight := 100, 100

	for i := 0; i < 3; i++ {
		rect := image.Rect(i*frameWidth, 0, (i+1)*frameWidth, frameHeight)
		p.frames[i] = spriteSheet.SubImage(rect).(*ebiten.Image)
	}
}

func (p *Plane) Frame(x, y int) *ebiten.Image {
	i := (x + y) % 3
	return p.frames[i]
}

func (p *Plane) Draw(screen *ebiten.Image) {

}
