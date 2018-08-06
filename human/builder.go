package human

import (
	"sync"

	"github.com/amyadzuki/amygolib/maths"

	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/math32"
)

type Human struct {
	age, gender, muscle, weight float64

	*Embed

	sync.Mutex

	finalized bool
}

func New(dec *obj.Decoder) (h *Human, err error) {
	h = new(Human)
	err = h.Init(dec)
	return
}

func (h *Human) Finalize() *Human {
	h.Lock() ; defer h.Unlock()
	if !h.finalized {
		h.update_unlocked(true)
		h.finalized = true
	}
	return h
}

func (h *Human) Finalized() bool {
	return h.finalized
}

func (h *Human) Init(dec *obj.Decoder) (err error) {
	h.Lock() ; defer h.Unlock()
	skin := &math32.Vector4{0.5, 0.5, 0.5, 0.25}
	eyes := &math32.Color4{1.0/3.0, 2.0/3.0, 1, 1}
	uwF := &math32.Color4{1, 1, 1, 1}
	uwD := &math32.Color4{0.875, 0.875, 0.875, 0.5}
	uwT := &math32.Color4{0xff/255.0, 0xb6/255.0, 0xc1/255.0, 1}
	if h.Embed, err = NewEmbed(dec, SkinDark, SkinLight, skin, Eyes, eyes, Underwear, uwF, uwD, uwT); err != nil {
		return
	}
	h.age, h.gender, h.muscle, h.weight = 0.5, 0.125, 0.5, 0.5
	h.finalized = false
	if HumanInit != nil {
		HumanInit(h)
	}
	if HumanUpdate != nil {
		HumanUpdate(h, false)
	}
	return
}

func (h *Human) Params() (float64, float64, float64, float64) {
	return h.age, h.gender, h.muscle, h.weight
}

func (h *Human) Update(age, gender, muscle, weight float64) *Human {
	h.Lock() ; defer h.Unlock()
	if !h.finalized {
		h.age = maths.ClampFloat64(age, 0, 1)
		h.gender = maths.ClampFloat64(gender, 0, 1)
		h.muscle = maths.ClampFloat64(muscle, 0, 1)
		h.weight = maths.ClampFloat64(weight, 0, 1)
	}
	h.update_unlocked(false)
	return h
}

func (h *Human) update_unlocked(final bool) {
	if h.finalized {
		return
	}
	if HumanUpdate != nil {
		HumanUpdate(h, final)
	}
}

var HumanInit func(*Human)
var HumanUpdate func(*Human, bool)
