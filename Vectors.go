package imgui

// #include "imguiWrapperTypes.h"
import "C"

// Vec2 represents a two-dimensional vector.
type Vec2 struct {
	X float32
	Y float32
}

func (vec *Vec2) wrapped() (out *C.IggVec2, finisher func()) {
	if vec != nil {
		out = &C.IggVec2{
			x: C.float(vec.X),
			y: C.float(vec.Y),
		}
		finisher = func() {
			vec.X = float32(out.x) // nolint: gotype
			vec.Y = float32(out.y) // nolint: gotype
		}
	} else {
		finisher = func() {}
	}
	return
}

// Vec4 represents a four-dimensional vector.
type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func (vec *Vec4) wrapped() (out *C.IggVec4, finisher func()) {
	if vec != nil {
		out = &C.IggVec4{
			x: C.float(vec.X),
			y: C.float(vec.Y),
			z: C.float(vec.Z),
			w: C.float(vec.W),
		}
		finisher = func() {
			vec.X = float32(out.x) // nolint: gotype
			vec.Y = float32(out.y) // nolint: gotype
			vec.Z = float32(out.z) // nolint: gotype
			vec.W = float32(out.w) // nolint: gotype
		}
	} else {
		finisher = func() {}
	}
	return
}
