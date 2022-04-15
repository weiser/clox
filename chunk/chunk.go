package chunk

// import "common"

const (
	OP_RETURN = iota
)

type Chunk struct {
	Code []uint8
}

func InitChunk(chunk *Chunk) {
	chunk = &Chunk{Code: make([]uint8, 0)}
}

func WriteChunk(chunk *Chunk, blob uint8) {
	chunk.Code = append(chunk.Code, blob)
}
