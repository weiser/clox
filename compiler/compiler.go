package compiler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/weiser/clox/chunk"
	"github.com/weiser/clox/debug"
	"github.com/weiser/clox/scanner"
	"github.com/weiser/clox/value"
)

type Precedence int

const (
	PREC_NONE       Precedence = iota
	PREC_ASSIGNMENT            // =
	PREC_OR                    // or
	PREC_AND                   // and
	PREC_EQUALITY              // == !=
	PREC_COMPARISON            // < > <= >=
	PREC_TERM                  // + -
	PREC_FACTOR                // * /
	PREC_UNARY                 // ! -
	PREC_CALL                  // . ()
	PREC_PRIMARY
)

type Parser struct {
	Current   scanner.Token
	Previous  scanner.Token
	HadError  bool
	PanicMode bool
}

type ParseFn func()

type ParseRule struct {
	Prefix ParseFn
	Infix  ParseFn
	Prec   Precedence
}

var _rules map[scanner.TokenType]ParseRule

func init() {
	_rules = map[scanner.TokenType]ParseRule{
		scanner.TOKEN_LEFT_PAREN:    {Prefix: grouping, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_RIGHT_PAREN:   ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_LEFT_BRACE:    ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_RIGHT_BRACE:   ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_COMMA:         ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_DOT:           ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_MINUS:         ParseRule{Prefix: unary, Infix: binary, Prec: PREC_TERM},
		scanner.TOKEN_PLUS:          ParseRule{Prefix: nil, Infix: binary, Prec: PREC_TERM},
		scanner.TOKEN_SEMICOLON:     ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_SLASH:         ParseRule{Prefix: nil, Infix: binary, Prec: PREC_FACTOR},
		scanner.TOKEN_STAR:          ParseRule{Prefix: nil, Infix: binary, Prec: PREC_FACTOR},
		scanner.TOKEN_BANG:          ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_BANG_EQUAL:    ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_GREATER:       ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_GREATER_EQUAL: ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_LESS:          ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_LESS_EQUAL:    ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_IDENTIFIER:    ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_STRING:        ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_NUMBER:        ParseRule{Prefix: number, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_AND:           ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_CLASS:         ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_ELSE:          ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_FALSE:         ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_FOR:           ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_FUN:           ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_IF:            ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_NIL:           ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_OR:            ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_PRINT:         ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_RETURN:        ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_SUPER:         ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_THIS:          ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_TRUE:          ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_VAR:           ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_WHILE:         ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_ERROR:         ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
		scanner.TOKEN_EOF:           ParseRule{Prefix: nil, Infix: nil, Prec: PREC_NONE},
	}
}

var _parser Parser
var _compilingChunk *chunk.Chunk
var _isDebug bool

func currentChunk() *chunk.Chunk {
	return _compilingChunk
}

func Compile(source string, _chunk *chunk.Chunk, isDebug bool) bool {
	_isDebug = isDebug
	scanner.InitScanner(source)
	_compilingChunk = _chunk

	_parser.PanicMode = false
	_parser.HadError = false

	advance()
	expression()
	consume(scanner.TOKEN_EOF, "expect end of expression")
	endCompiler()
	return !_parser.HadError
}

func advance() {
	_parser.Previous = _parser.Current
	for {
		_parser.Current = scanner.ScanToken()
		if _parser.Current.Type != scanner.TOKEN_ERROR {
			break
		}

		errorAtCurrent(_parser.Current.Value)
	}
}

func consume(_type scanner.TokenType, msg string) {
	if _parser.Current.Type == _type {
		advance()
		return
	}
	errorAtCurrent(msg)
}

func emitByte(_byte byte) {
	chunk.WriteChunk(currentChunk(), _byte, _parser.Previous.Line)
}

func emitBytes(b1 byte, b2 byte) {
	emitByte(b1)
	emitByte(b2)
}

func endCompiler() {
	emitReturn()
	if _isDebug && !_parser.HadError {
		debug.DisassembleChunk(currentChunk(), "code")
	}
}

func binary() {
	operatorType := _parser.Previous.Type
	rule := getRule(operatorType)
	parsePrecedence(rule.Prec + 1)
	switch operatorType {
	case scanner.TOKEN_PLUS:
		emitByte(chunk.OP_ADD)
	case scanner.TOKEN_MINUS:
		emitByte(chunk.OP_SUBTRACT)
	case scanner.TOKEN_STAR:
		emitByte(chunk.OP_MULTIPLY)
	case scanner.TOKEN_SLASH:
		emitByte(chunk.OP_DIVIDE)
	default:
		return
	}
}

func number() {
	_value, _ := strconv.ParseFloat(_parser.Previous.Value, 64)
	emitConstant(value.NumberVal(_value))
}

func emitConstant(_value value.Value) {
	emitBytes(chunk.OP_CONSTANT, makeConstant(_value))
}

func makeConstant(value value.Value) byte {
	constant := chunk.AddConstant(currentChunk(), value)
	return byte(constant)
}

func expression() {
	parsePrecedence(PREC_ASSIGNMENT)
}

func grouping() {
	expression()
	consume(scanner.TOKEN_RIGHT_PAREN, "Expected ')' after expression")
}

func unary() {
	operatorType := _parser.Previous.Type

	// compile the operand
	parsePrecedence(PREC_UNARY)

	switch operatorType {
	case scanner.TOKEN_MINUS:
		// functionally, this has us compute the expression, then negate it
		emitByte(chunk.OP_NEGATE)
	default:
		return
	}
}

func parsePrecedence(precedence Precedence) {
	advance()
	prefixRule := getRule(_parser.Previous.Type).Prefix
	if prefixRule == nil {
		errorAtCurrent("Expect expression")
	}

	prefixRule()

	for precedence <= getRule(_parser.Current.Type).Prec {
		advance()
		infixRule := getRule(_parser.Previous.Type).Infix
		infixRule()
	}

}

func getRule(_type scanner.TokenType) ParseRule {
	return _rules[_type]
}

func emitReturn() {
	emitByte(chunk.OP_RETURN)
}

func errorAtCurrent(msg string) {
	errorAt(&_parser.Previous, msg)
}

func errorAt(token *scanner.Token, msg string) {
	if _parser.PanicMode {
		return
	}
	_parser.PanicMode = true
	os.Stderr.Write([]byte(fmt.Sprintf("[line %v] Error", token.Line)))
	if token.Type == scanner.TOKEN_EOF {
		os.Stderr.Write([]byte(" at end"))
	} else {
		if token.Type == scanner.TOKEN_ERROR {
			// do nothing
		} else {
			os.Stderr.Write([]byte(fmt.Sprintf(" at '%v'", token.Value)))
		}
	}

	os.Stderr.Write([]byte(fmt.Sprintf(": %v\n", msg)))
	_parser.HadError = true
}
