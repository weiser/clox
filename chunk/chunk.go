package chunk

// import "common"
import "github.com/weiser/clox/value"

const (
	OP_RETURN = iota
	OP_CONSTANT
	OP_NEGATE
	OP_ADD
	OP_SUBTRACT
	OP_MULTIPLY
	OP_DIVIDE
)

// if we were really itching about performance, Lines and Code would be pointers to slices and not slices.
type Chunk struct {
	Code      []uint8
	Lines     []int
	Constants value.ValueArray
}

func InitChunk(chunk *Chunk) {
	chunk = &Chunk{Code: make([]uint8, 0), Constants: make(value.ValueArray, 0), Lines: make([]int, 0)}
}

func WriteChunk(chunk *Chunk, blob uint8, lineno int) {
	chunk.Code = append(chunk.Code, blob)
	chunk.Lines = append(chunk.Lines, lineno)
}

// AddConstants adds a Value to Constants, returning the index that the Value was placed in Constants
func AddConstant(chunk *Chunk, v value.Value) int {
	chunk.Constants = append(chunk.Constants, v)
	return len(chunk.Constants) - 1
}
