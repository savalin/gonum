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

// AbsorbedRadioactiveDose is a measure of absorbed dose of ionizing radiation in grays.
type AbsorbedRadioactiveDose float64

const Gray AbsorbedRadioactiveDose = 1

// Unit converts the AbsorbedRadioactiveDose to a *Unit.
func (a AbsorbedRadioactiveDose) Unit() *Unit {
	return New(float64(a), Dimensions{
		LengthDim: 2,
		TimeDim:   -2,
	})
}

// AbsorbedRadioactiveDose allows AbsorbedRadioactiveDose to implement a AbsorbedRadioactiveDoseer interface.
func (a AbsorbedRadioactiveDose) AbsorbedRadioactiveDose() AbsorbedRadioactiveDose {
	return a
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension.
func (a *AbsorbedRadioactiveDose) From(u Uniter) error {
	if !DimensionsMatch(u, Gray) {
		*a = AbsorbedRadioactiveDose(math.NaN())
		return errors.New("unit: dimension mismatch")
	}
	*a = AbsorbedRadioactiveDose(u.Unit().Value())
	return nil
}

func (a AbsorbedRadioactiveDose) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", a, float64(a))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " Gy"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(a))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(a))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(a))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(a))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g Gy)", c, a, float64(a))
	}
}
