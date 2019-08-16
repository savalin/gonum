// Code generated by "go generate github.com/savalin/gonum/unit”; DO NOT EDIT.

// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"errors"
	"fmt"
	"math"
	"unicode/utf8"
)

// Volume represents a volume in cubic metres.
type Volume float64

const Litre Volume = 1e-3

// Unit converts the Volume to a *Unit
func (v Volume) Unit() *Unit {
	return New(float64(v), Dimensions{
		LengthDim: 3,
	})
}

// Volume allows Volume to implement a Volumeer interface
func (v Volume) Volume() Volume {
	return v
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (v *Volume) From(u Uniter) error {
	if !DimensionsMatch(u, Litre) {
		*v = Volume(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*v = Volume(u.Unit().Value())
	return nil
}

func (v Volume) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", v, float64(v))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " m^3"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(v))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(v))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(v))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(v))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g m^3)", c, v, float64(v))
	}
}
