package objects

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	*Abstract
	Damage         int
	lastUpdateTime time.Time
	image          *ebiten.Image
}

func NewBullet(damage int, image *ebiten.Image) *Bullet {
	return &Bullet{
		Abstract: NewAbstract(),
		Damage:   damage,
		image:    image,
	}
}

func (b *Bullet) Animate() {
	if time.Since(b.lastUpdateTime) > time.Millisecond { // Change frame every 150ms
		b.SetPosition(b.x+10, b.y)
		b.lastUpdateTime = time.Now()
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(b.image, op)

}
