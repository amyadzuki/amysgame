package human

import (
	"sync"

	"github.com/amyadzuki/amygolib/maths"

	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/math32"
)

type Builder struct {
	age, gender, muscle, weight float64

	human *Human

	sync.Mutex

	finalized bool
}

func New(dec *obj.Decoder) (b *Builder, err error) {
	b = new(Builder)
	err = b.Init(dec)
	return
}

func (b *Builder) Finalize() *Human {
	b.Lock() ; defer b.Unlock()
	if !b.finalized {
		b.update_unlocked(true)
		b.finalized = true
	}
	return b.human
}

func (b *Builder) Finalized() bool {
	return b.finalized
}

func (b *Builder) Init(dec *obj.Decoder) (err error) {
	b.Lock() ; defer b.Unlock()
	skin := &math32.Vector4{0.5, 0.5, 0.5, 0.25}
	eyes := &math32.Color4{1.0/3.0, 2.0/3.0, 1, 1}
	uwF := &math32.Color4{1, 1, 1, 1}
	uwD := &math32.Color4{0.875, 0.875, 0.875, 0.5}
	uwT := &math32.Color4{0xff/255.0, 0xb6/255.0, 0xc1/255.0, 1}
	if b.human, err = NewHuman(dec, SkinDark, SkinLight, skin, Eyes, eyes, Underwear, uwF, uwD, uwT); err != nil {
		return
	}
	b.age, b.gender, b.muscle, b.weight = 0.5, 0.125, 0.5, 0.5
	b.finalized = false
	if BuilderInit != nil {
		BuilderInit(b)
	}
	if BuilderUpdate != nil {
		BuilderUpdate(b, false)
	}
	return
}

func (b *Builder) Params() (float64, float64, float64, float64) {
	return b.age, b.gender, b.muscle, b.weight
}

func (b *Builder) Update(age, gender, muscle, weight float64) *Builder {
	b.Lock() ; defer b.Unlock()
	if !b.finalized {
		b.age = maths.ClampFloat64(age, 0, 1)
		b.gender = maths.ClampFloat64(gender, 0, 1)
		b.muscle = maths.ClampFloat64(muscle, 0, 1)
		b.weight = maths.ClampFloat64(weight, 0, 1)
	}
	b.update_unlocked(false)
	return b
}

func (b *Builder) update_unlocked(final bool) {
	if b.finalized {
		return
	}
	if BuilderUpdate != nil {
		BuilderUpdate(b, final)
	}
}

var BuilderInit func(*Builder)
var BuilderUpdate func(*Builder, bool)
