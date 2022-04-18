package chunk

// import "common"
import "github.com/weiser/clox/value"

const (
	OP_RETURN = iota
	OP_CONSTANT
)

type Chunk struct {
	Code      []uint8
	Constants value.ValueArray
}

func InitChunk(chunk *Chunk) {
	chunk = &Chunk{Code: make([]uint8, 0), Constants: make(value.ValueArray, 0)}
}

func WriteChunk(chunk *Chunk, blob uint8) {
	chunk.Code = append(chunk.Code, blob)
}

// AddConstants adds a Value to Constants, returning the index that the Value was placed in Constants
func AddConstant(chunk *Chunk, v value.Value) int {
	chunk.Constants = append(chunk.Constants, v)
	return len(chunk.Constants) - 1
}
