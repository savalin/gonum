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

// Voltage represents a voltage in Volts.
type Voltage float64

const Volt Voltage = 1

// Unit converts the Voltage to a *Unit
func (v Voltage) Unit() *Unit {
	return New(float64(v), Dimensions{
		CurrentDim: -1,
		LengthDim:  2,
		MassDim:    1,
		TimeDim:    -3,
	})
}

// Voltage allows Voltage to implement a Voltager interface
func (v Voltage) Voltage() Voltage {
	return v
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (v *Voltage) From(u Uniter) error {
	if !DimensionsMatch(u, Volt) {
		*v = Voltage(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*v = Voltage(u.Unit().Value())
	return nil
}

func (v Voltage) Format(fs fmt.State, c rune) {
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
		const unit = " V"
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
		fmt.Fprintf(fs, "%%!%c(%T=%g V)", c, v, float64(v))
	}
}
