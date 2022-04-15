package debug

import (
	"fmt"

	"github.com/weiser/clox/chunk"
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

	switch instruction {
	case chunk.OP_RETURN:
		return SimpleInstruction("OP_RETURN", offset)
	default:
		fmt.Println("Unknown OpCode: ", instruction)
	}
	return offset + 1
}

func SimpleInstruction(name string, offset int) int {
	fmt.Println(name)
	return offset + 1
}
