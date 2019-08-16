package path

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/savalin/gonum/blas/blas64"
	"github.com/savalin/gonum/graph"
	"github.com/savalin/gonum/graph/simple"
	"github.com/savalin/gonum/mat"
)

func NewAllShortestByDump(dump []byte) (AllShortest, error) {
	var v1Marshaler = &allShortestDumperV1{}

	return v1Marshaler.Unmarshal(dump)
}

func (g *AllShortest) Marshal() ([]byte, error) {
	var v1Marshaler = &allShortestDumperV1{}

	return v1Marshaler.Marshal(g)
}

const allShortestDumperVersion1 = 1

type allShortestDumperV1 struct {
	Version int
	Nodes   []int64
	IndexOf map[int64]int
	Dist    *dense
	Next    [][]int
	Forward bool
}

// Dense is a dense matrix representation.
type dense struct {
	Mat blas64.General

	CapRows, CapCols int
}

func (d *allShortestDumperV1) Marshal(g *AllShortest) ([]byte, error) {
	d.Version = allShortestDumperVersion1

	d.Nodes = make([]int64, len(g.nodes))
	for i, node := range g.nodes {
		d.Nodes[i] = node.ID()
	}

	d.IndexOf = make(map[int64]int, len(g.indexOf))
	for k, v := range g.indexOf {
		d.IndexOf[k] = v
	}

	if g.dist != nil {
		v := reflect.ValueOf(*(g.dist))
		mat := v.FieldByName("mat")

		dataField := mat.FieldByName("Data")
		data := make([]float64, dataField.Len())
		for i := 0; i < dataField.Len(); i++ {
			val := dataField.Index(i).Float()
			if val < 0 {
				fmt.Printf("DATA < 0 FOUND %d\n", val)
			} else if val == math.Inf(1) {
				//fmt.Printf("DATA MAX FOUND\n")
				val = -1
			}
			data[i] = val
		}

		d.Dist = &dense{
			Mat: blas64.General{

				Rows:   int(mat.FieldByName("Rows").Int()),
				Cols:   int(mat.FieldByName("Cols").Int()),
				Stride: int(mat.FieldByName("Stride").Int()),
				Data:   data,
			},

			CapRows: int(v.FieldByName("capRows").Int()),
			CapCols: int(v.FieldByName("capCols").Int()),
		}
	}

	d.Next = make([][]int, len(g.next))
	for i, s := range g.next {
		d.Next[i] = s
	}

	d.Forward = g.forward

	return json.Marshal(d)
}

func (d *allShortestDumperV1) Unmarshal(data []byte) (result AllShortest, err error) {
	if err = json.Unmarshal(data, d); err != nil {
		return
	}

	if d.Version != allShortestDumperVersion1 {
		err = errors.New("dump marshaled by other dumper version")
		return
	}

	result.indexOf = d.IndexOf
	result.next = d.Next
	result.forward = d.Forward

	result.nodes = make([]graph.Node, len(d.Nodes))
	for i, node := range d.Nodes {
		result.nodes[i] = simple.Node(node)
	}

	if d.Dist != nil {
		for i, v := range d.Dist.Mat.Data {
			if v == -1 {
				d.Dist.Mat.Data[i] = math.Inf(1)
			}
		}

		result.dist = mat.NewDense(
			d.Dist.CapRows,
			d.Dist.CapCols,
			d.Dist.Mat.Data,
		)
	}

	return
}
