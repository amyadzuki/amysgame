package human

import (
	"sync"

	"github.com/g3n/engine/loader/obj"
)

type Builder struct {
	f0, f1, f2, f3 float64

	F, M *Human

	sync.Mutex

	finalized, male bool
}

func New(f, m *obj.Decoder) *Builder {
	return new(Builder).Init(f, m)
}

func (b *Builder) Finalize() *Human {
	b.Lock() ; defer b.Unlock()
	if !b.finalized {
		b.update_unlocked(true)
		b.finalized = true
	}
	if b.male {
		return b.M
	} else {
		return b.F
	}
}

func (b *Builder) Finalized() bool {
	return b.finalized
}

func (b *Builder) Init(f, m *obj.Decoder) *Builder {
	b.Lock() ; defer b.Unlock()
	skinDelta = &math32.Vector4{0.5, 0.5, 0.5, 0.25}
	eyeColor = &math32.Color4{1.0/3.0, 2.0/3.0, 1, 1}
	uwF = &math32.Color4{1, 1, 1, 1}
	uwD = &math32.Color4{0.875, 0.875, 0.875, 0.5}
	uwT = &math32.Color4{0xff/255.0, 0xb6/255.0, 0xc1/255.0, 1}
	b.F = NewHuman(f, SkinDarkF, SkinLightF, skinDelta, Eyes, eyeColor, UnderwearF, uwF, uwD, uwT)
	b.M = NewHuman(m, SkinDarkM, SkinLightM, skinDelta, Eyes, eyeColor, UnderwearM, uwF, uwD, uwT)
	b.f0, b.f1, b.f2, b.f3 = 0.5, 0.125, 0.5, 0.5
	b.finalized, b.male = false, false
	if BuilderInit != nil {
		BuilderInit(b)
	}
	if BuilderUpdate != nil {
		BuilderUpdate(b, false)
	}
	return b
}

func (b *Builder) Params() (float64, float64, float64, float64) {
	return b.f0, b.f1, b.f2, b.f3
}

func (b *Builder) Update(f0, f1, f2, f3 float64) *Builder {
	b.Lock() ; defer b.Unlock()
	if !b.finalized {
		b.f0 = maths.ClampFloat64(f0, 0, 1)
		b.f1 = maths.ClampFloat64(f1, 0, 1)
		b.f2 = maths.ClampFloat64(f2, 0, 1)
		b.f3 = maths.ClampFloat64(f3, 0, 1)
	}
	update_unlocked(false)
}

func (b *Builder) update_unlocked(final bool) *Builder {
	if b.finalized {
		return
	}
	if BuilderUpdate != nil {
		BuilderUpdate(b, final)
	}
}

var BuilderInit func(*Builder)
var BuilderUpdate func(*Builder, bool)
