package human

import (
	"sync"

	"github.com/amyadzuki/amygolib/maths"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/math32"
)

type Builder struct {
	age, gender, muscle, weight float64

	F, M *Human

	*core.Node

	sync.Mutex

	finalized bool
}

func New(f, m *obj.Decoder) (b *Builder, err error) {
	b = new(Builder)
	err = b.Init(f, m)
	return
}

func (b *Builder) Female() bool {
	return b.gender <= 0.5
}

func (b *Builder) Finalize() *Human {
	b.Lock() ; defer b.Unlock()
	if !b.finalized {
		b.update_unlocked(true)
		b.finalized = true
	}
	if b.Female() {
		return b.F
	} else {
		return b.M
	}
}

func (b *Builder) Finalized() bool {
	return b.finalized
}

func (b *Builder) Init(f, m *obj.Decoder) (err error) {
	b.Lock() ; defer b.Unlock()
	skinDelta := &math32.Vector4{0.5, 0.5, 0.5, 0.25}
	eyeColor := &math32.Color4{1.0/3.0, 2.0/3.0, 1, 1}
	uwF := &math32.Color4{1, 1, 1, 1}
	uwD := &math32.Color4{0.875, 0.875, 0.875, 0.5}
	uwT := &math32.Color4{0xff/255.0, 0xb6/255.0, 0xc1/255.0, 1}
	b.F, err = NewHuman(f, SkinDarkF, SkinLightF, skinDelta, Eyes, eyeColor, UnderwearF, uwF, uwD, uwT)
	if err != nil {
		return
	}
	b.M, err = NewHuman(m, SkinDarkM, SkinLightM, skinDelta, Eyes, eyeColor, UnderwearM, uwF, uwD, uwT)
	if err != nil {
		return
	}
	b.Node = core.NewNode()
	b.Node.Add(b.F)
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

func (b *Builder) Male() bool {
	return !b.Female()
}

func (b *Builder) Params() (float64, float64, float64, float64) {
	return b.age, b.gender, b.muscle, b.weight
}

func (b *Builder) Update(age, gender, muscle, weight float64) *Builder {
	b.Lock() ; defer b.Unlock()
	if !b.finalized {
		b.age = maths.ClampFloat64(age, 0, 1)
		if b.Female() {
			b.Node.Remove(b.F)
		} else {
			b.Node.Remove(b.M)
		}
		b.gender = maths.ClampFloat64(gender, 0, 1)
		if b.Female() {
			b.Node.Add(b.F)
		} else {
			b.Node.Add(b.M)
		}
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
