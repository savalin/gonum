// Code generated by "go generate github.com/savalin/gonum/unit/constant”; DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package constant

import (
	"fmt"

	"github.com/savalin/gonum/unit"
)

// Faraday is the Faraday constant, the magnitude of electric charge per mole of electrons.
// The dimensions of Faraday are A s mol^-1. The standard uncertainty of the constant is 0.00059 A s mol^-1.
const Faraday = faradayUnits(96485.33289)

type faradayUnits float64

// Unit converts the faradayUnits to a *unit.Unit
func (cnst faradayUnits) Unit() *unit.Unit {
	return unit.New(float64(cnst), unit.Dimensions{
		unit.CurrentDim: 1,
		unit.TimeDim:    1,
		unit.MoleDim:    -1,
	})
}

func (cnst faradayUnits) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", cnst, float64(cnst))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), w, p, cnst.Unit())
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, cnst.Unit())
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), w, cnst.Unit())
		default:
			fmt.Fprintf(fs, "%"+string(c), cnst.Unit())
		}
	default:
		fmt.Fprintf(fs, "%%!"+string(c)+"(constant.faradayUnits=%v A s mol^-1)", float64(cnst))
	}
}
