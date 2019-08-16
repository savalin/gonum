#!/bin/bash

go generate github.com/savalin/gonum/blas
go generate github.com/savalin/gonum/blas/gonum
go generate github.com/savalin/gonum/unit
go generate github.com/savalin/gonum/unit/constant
go generate github.com/savalin/gonum/graph/formats/dot

if [ -n "$(git diff)" ]; then	
	git diff
	exit 1
fi
