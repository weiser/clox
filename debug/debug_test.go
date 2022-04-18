package debug

import (
	"github.com/weiser/clox/chunk"
)

func ExampleDisassembleChunk_op_return() {

	var chnk chunk.Chunk
	chunk.InitChunk(&chnk)
	chunk.WriteChunk(&chnk, chunk.OP_RETURN)
	DisassembleChunk(&chnk, "test chunk")
	// Output: == test chunk ==
	// 0000: OP_RETURN
}

func ExampleDisassembleChunk_op_constant() {

	var chnk chunk.Chunk
	chunk.InitChunk(&chnk)
	idx := uint8(chunk.AddConstant(&chnk, 1234.0))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT)
	chunk.WriteChunk(&chnk, idx)
	DisassembleChunk(&chnk, "test chunk")
	// Output: == test chunk ==
	// 0000: OP_CONSTANT         0 1234.000000

}
