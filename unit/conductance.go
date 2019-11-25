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

// Conductance represents an electrical conductance in Siemens.
type Conductance float64

const Siemens Conductance = 1

// Unit converts the Conductance to a *Unit.
func (co Conductance) Unit() *Unit {
	return New(float64(co), Dimensions{
		CurrentDim: 2,
		LengthDim:  -2,
		MassDim:    -1,
		TimeDim:    3,
	})
}

// Conductance allows Conductance to implement a Conductancer interface.
func (co Conductance) Conductance() Conductance {
	return co
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension.
func (co *Conductance) From(u Uniter) error {
	if !DimensionsMatch(u, Siemens) {
		*co = Conductance(math.NaN())
		return errors.New("unit: dimension mismatch")
	}
	*co = Conductance(u.Unit().Value())
	return nil
}

func (co Conductance) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", co, float64(co))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " S"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(co))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(co))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(co))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(co))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g S)", c, co, float64(co))
	}
}
