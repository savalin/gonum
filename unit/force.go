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

// Force represents a force in Newtons.
type Force float64

const Newton Force = 1

// Unit converts the Force to a *Unit
func (f Force) Unit() *Unit {
	return New(float64(f), Dimensions{
		LengthDim: 1,
		MassDim:   1,
		TimeDim:   -2,
	})
}

// Force allows Force to implement a Forcer interface
func (f Force) Force() Force {
	return f
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (f *Force) From(u Uniter) error {
	if !DimensionsMatch(u, Newton) {
		*f = Force(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*f = Force(u.Unit().Value())
	return nil
}

func (f Force) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", f, float64(f))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " N"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(f))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(f))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(f))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(f))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g N)", c, f, float64(f))
	}
}
