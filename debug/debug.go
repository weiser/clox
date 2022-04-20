package debug

import (
	"fmt"

	"github.com/weiser/clox/chunk"
	"github.com/weiser/clox/value"
)

func DisassembleChunk(chnk *chunk.Chunk, name string) {
	fmt.Println("==", name, "==")
	for offset := 0; offset < len(chnk.Code); {
		offset = DisassembleInstruction(chnk, offset)
	}
}

func DisassembleInstruction(chnk *chunk.Chunk, offset int) int {
	fmt.Printf("%04d: ", offset)
	instruction := chnk.Code[offset]
	if offset > 0 && chnk.Lines[offset] == chnk.Lines[offset-1] {
		fmt.Printf("   | ")
	} else {
		fmt.Printf("%04d ", chnk.Lines[offset])
	}

	switch instruction {
	case chunk.OP_RETURN:
		return SimpleInstruction("OP_RETURN", offset)
	case chunk.OP_CONSTANT:
		return ConstantInstruction("OP_CONSTANT", chnk, offset)
	default:
		fmt.Println("Unknown OpCode: ", instruction)
	}
	return offset + 1
}

func ConstantInstruction(name string, chnk *chunk.Chunk, offset int) int {
	constant := chnk.Code[offset+1]
	fmt.Printf("%-16s %4d ", name, constant)
	value.PrintValue(chnk.Constants[constant])
	fmt.Println()
	// why offset + 2 ?  offset points at the opcode initially.
	// The constant portion of this instruction is at offset + 1.
	//  We want to return the offset to the next instruction which
	//   is one after the constant.
	return offset + 2

}

func SimpleInstruction(name string, offset int) int {
	fmt.Println(name)
	return offset + 1
}
