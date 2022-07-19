package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/weiser/clox/vm"
)

// import "common"

func main() {

	debugPtr := flag.Bool("debug", false, "a bool. true if in debug mode")

	flag.Parse()

	vm.InitVM(*debugPtr)

	programArgs := flag.Args()

	fmt.Println("debug is: %v", *debugPtr)

	if len(programArgs) == 0 {
		Repl(*debugPtr)
	} else if len(programArgs) == 1 {
		RunFile(programArgs[0], *debugPtr)
	} else {
		fmt.Println("Usage: clox [-debug] [path]")
	}

	vm.FreeVM()
}

// Provides a REPL for clox. doesn't provide for multiline statements
func Repl(isDebug bool) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()
		if text == "\n" {
			break
		}
		vm.InterpretSource(text, isDebug)
	}
}

func RunFile(path string, isDebug bool) {
	source, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Couldn't read %v", path))
	}

	result := vm.InterpretSource(string(source), isDebug)

	if result == vm.INTERPRET_COMPILE_ERROR {
		panic("INTERPRET COMPILE ERROR")
	}
	if result == vm.INTERPRET_RUNTIME_ERROR {
		panic("INTERPRET RUNTIME ERROR")
	}

}
