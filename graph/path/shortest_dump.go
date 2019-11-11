package path

import (
	"encoding/json"
	"errors"
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

type allShortestDumperV1 struct {}
type allShortestDump struct {
	Version int
	Nodes   []int64
	IndexOf map[int64]int
	Dist    *Dense
	Next    [][]int
	Forward bool
}

// Dense is a dense matrix representation.
type Dense struct {
	Mat blas64.General

	CapRows, CapCols int
}

func (d *allShortestDumperV1) Marshal(g *AllShortest) ([]byte, error) {
	var dump = &allShortestDump{}

	dump.Version = allShortestDumperVersion1

	dump.Nodes = make([]int64, len(g.nodes))
	for i, node := range g.nodes {
		dump.Nodes[i] = node.ID()
	}

	dump.IndexOf = make(map[int64]int, len(g.indexOf))
	for k, v := range g.indexOf {
		dump.IndexOf[k] = v
	}

	if g.dist != nil {
		v := reflect.ValueOf(*(g.dist))
		mat := v.FieldByName("mat")

		dataField := mat.FieldByName("Data")
		data := make([]float64, dataField.Len())
		for i := 0; i < dataField.Len(); i++ {
			val := dataField.Index(i).Float()
			if val == math.Inf(1) {
				val = -1
			}
			data[i] = val
		}

		dump.Dist = &Dense{
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

	dump.Next = make([][]int, len(g.next))
	for i, s := range g.next {
		dump.Next[i] = s
	}

	dump.Forward = g.forward

	return json.Marshal(dump)
}

func (d *allShortestDumperV1) Unmarshal(data []byte) (result AllShortest, err error) {
	var dump = &allShortestDump{}
	if err = json.Unmarshal(data, dump); err != nil {
		return
	}

	if dump.Version != allShortestDumperVersion1 {
		err = errors.New("dump marshaled by other dumper version")
		return
	}

	result.indexOf = dump.IndexOf
	result.next = dump.Next
	result.forward = dump.Forward

	result.nodes = make([]graph.Node, len(dump.Nodes))
	for i, node := range dump.Nodes {
		result.nodes[i] = simple.Node(node)
	}

	if dump.Dist != nil {
		for i, v := range dump.Dist.Mat.Data {
			if v == -1 {
				dump.Dist.Mat.Data[i] = math.Inf(1)
			}
		}

		result.dist = mat.NewDense(
			dump.Dist.CapRows,
			dump.Dist.CapCols,
			dump.Dist.Mat.Data,
		)
	}

	return
}
