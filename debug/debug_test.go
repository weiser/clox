package debug

import (
	"github.com/weiser/clox/chunk"
	"github.com/weiser/clox/value"
)

func ExampleDisassembleChunk_op_return() {

	var chnk chunk.Chunk
	chunk.InitChunk(&chnk)
	chunk.WriteChunk(&chnk, chunk.OP_RETURN, 0)
	DisassembleChunk(&chnk, "test chunk")
	// Output: == test chunk ==
	// 0000: 0000 OP_RETURN
}

func ExampleDisassembleChunk_op_constant() {

	var chnk chunk.Chunk
	chunk.InitChunk(&chnk)
	idx := uint8(chunk.AddConstant(&chnk, value.NumberVal(1234.0)))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT, 0)
	chunk.WriteChunk(&chnk, idx, 0)
	chunk.WriteChunk(&chnk, chunk.OP_RETURN, 0)
	DisassembleChunk(&chnk, "test chunk")
	// Output: == test chunk ==
	// 0000: 0000 OP_CONSTANT         0 1234.000000
	// 0002:    | OP_RETURN

}
