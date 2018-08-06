package human

import (
	"sync"
)

type Builder struct {
	F, M           Human
	F0, F1, F2, F3 float64

	sync.Mutex

	finalized, male bool
}

func New(f, m *obj.Decoder) *Builder {
	return new(Builder).Init(f, m)
}

func (b *Builder) Finalize() *Human {
	b.Lock() ; defer b.Unlock()
	if !b.finalized {
		b.update_unlocked(b.F0, b.F1, b.F2, b.F3, true)
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
	skinDelta = &math32.Vector4{0.5, 0.5, 0.5, 0.25}
	eyeColor = &math32.Color4{1.0/3.0, 2.0/3.0, 1, 1}
	uwF = &math32.Color4{1, 1, 1, 1}
	uwD = &math32.Color4{0.875, 0.875, 0.875, 0.5}
	uwT = &math32.Color4{0xff/255.0, 0xb6/255.0, 0xc1/255.0, 1}
	b.F.Init(f, SkinDarkF, SkinLightF, skinDelta, Eyes, eyeColor, UnderwearF, uwF, uwD, uwT)
	b.M.Init(m, SkinDarkM, SkinLightM, skinDelta, Eyes, eyeColor, UnderwearM, uwF, uwD, uwT)
	b.Update = updateBuilder
	b.finalized, b.male = false, false
	if BuilderInit != nil {
		BuilderInit(b)
	}
	return b
}

func (b *Builder) Update(f0, f1, f2, f3 float64) *Builder {
	b.Lock() ; defer b.Unlock()
	update_unlocked(f0, f1, f2, f3)
}

func (b *Builder) update_unlocked(f0, f1, f2, f3 float64, final bool) *Builder {
	if b.finalized {
		return
	}
	if !final {
		b.F0, b.F1, b.F2, b.F3 = f0, f1, f2, f3
	}
	if BuilderUpdate != nil {
		BuilderUpdate(b, f0, f1, f2, f3, final)
	}
}

var BuilderInit func(*Builder)
var BuilderUpdate func(*Builder, float64, float64, float64, float64, bool)
