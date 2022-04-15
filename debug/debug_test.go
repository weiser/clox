package debug

import (
	"github.com/weiser/clox/chunk"
)

func ExampleDisassembleChunk() {

	var chnk chunk.Chunk
	chunk.InitChunk(&chnk)
	chunk.WriteChunk(&chnk, chunk.OP_RETURN)
	DisassembleChunk(&chnk, "test chunk")
	// Output: == test chunk ==
	// 0000: OP_RETURN
}
