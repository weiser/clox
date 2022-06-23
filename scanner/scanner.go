package scanner

type TokenType int

const (
	// SINGLE CHAR TOKENS
	TOKEN_LEFT_PAREN TokenType = iota
	TOKEN_RIGHT_PAREN
	TOKEN_LEFT_BRACE
	TOKEN_RIGHT_BRACE
	TOKEN_COMMA
	TOKEN_DOT
	TOKEN_MINUS
	TOKEN_PLUS
	TOKEN_SEMICOLON
	TOKEN_SLASH
	TOKEN_STAR
	// ONE OR TWO CHAR TOKENS
	TOKEN_BANG
	TOKEN_BANG_EQUAL
	TOKEN_EQUAL
	TOKEN_EQUAL_EQUAL
	TOKEN_GREATER
	TOKEN_GREATER_EQUAL
	TOKEN_LESS
	TOKEN_LESS_EQUAL
	//LITERALS
	TOKEN_IDENTIFIER
	TOKEN_STRING
	TOKEN_NUMBER
	//KEYWORDS
	TOKEN_AND
	TOKEN_CLASS
	TOKEN_ELSE
	TOKEN_FALSE
	TOKEN_FOR
	TOKEN_FUN
	TOKEN_IF
	TOKEN_NIL
	TOKEN_OR
	TOKEN_PRINT
	TOKEN_RETURN
	TOKEN_SUPER
	TOKEN_THIS
	TOKEN_TRUE
	TOKEN_VAR
	TOKEN_WHILE

	TOKEN_ERROR
	TOKEN_EOF
)

type Token struct {
	Type TokenType
	// Yes, I know this should prolly be a pointer to a string,
	// but i'm biasing towards speed of implementation and not
	// worrying about speed of execution
	Value string
	Line  int
}

type Scanner struct {
	Source            string
	SourceStart_idx   int
	SourceCurrent_idx int
	Line              int
}

var _scanner Scanner

func InitScanner(source string) {
	_scanner = Scanner{Source: source, SourceStart_idx: 0, SourceCurrent_idx: 0, Line: 1}
}

func ScanToken() Token {
	skipWhitespace()
	_scanner.SourceStart_idx = _scanner.SourceCurrent_idx
	if isAtEnd() {
		return makeToken(TOKEN_EOF)
	}

	c := advance()
	if isDigit(c) {
		return number()
	}
	if isAlpha(c) {
		return identifier()
	}

	switch c {
	case '(':
		return makeToken(TOKEN_LEFT_PAREN)
	case ')':
		return makeToken(TOKEN_RIGHT_PAREN)
	case '{':
		return makeToken(TOKEN_LEFT_BRACE)
	case '}':
		return makeToken(TOKEN_RIGHT_BRACE)
	case ',':
		return makeToken(TOKEN_COMMA)
	case '.':
		return makeToken(TOKEN_DOT)
	case '-':
		return makeToken(TOKEN_MINUS)
	case '+':
		return makeToken(TOKEN_PLUS)
	case ';':
		return makeToken(TOKEN_SEMICOLON)
	case '/':
		return makeToken(TOKEN_SLASH)
	case '*':
		return makeToken(TOKEN_STAR)
	case '!':
		return makeToken(matchIfElse('=', TOKEN_BANG_EQUAL, TOKEN_BANG))
	case '=':
		return makeToken(matchIfElse('=', TOKEN_EQUAL_EQUAL, TOKEN_EQUAL))
	case '<':
		return makeToken(matchIfElse('=', TOKEN_LESS_EQUAL, TOKEN_LESS))
	case '>':
		return makeToken(matchIfElse('=', TOKEN_GREATER_EQUAL, TOKEN_GREATER))
	}

	return errorToken("unexpected character")
}

func skipWhitespace() {
	for {
		r := peek()
		if isDigit(r) {
			number()
			return
		}
		switch r {
		case ' ':
			advance()
		case '\r':
			advance()
		case '\t':
			advance()
		case '\n':
			_scanner.Line += 1
			advance()
		case '/':
			if peekNext() == '/' {
				for peek() != '\n' && !isAtEnd() {
					advance()
				}
			}
		case '"':
			makeString()
			return
		default:
			return
		}
	}
}

func isAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
}

func identifier() Token {
	for !isAtEnd() && (isAlpha(peek()) || isDigit(peek())) {
		advance()
	}
	return makeToken(identifierType())
}

func identifierType() TokenType {
	switch _scanner.Source[_scanner.SourceStart_idx] {
	case 'a':
		return checkKeyword("and", TOKEN_AND)
	case 'c':
		return checkKeyword("class", TOKEN_CLASS)
	case 'e':
		return checkKeyword("else", TOKEN_ELSE)
	case 'i':
		return checkKeyword("if", TOKEN_IF)
	case 'n':
		return checkKeyword("nil", TOKEN_NIL)
	case 'o':
		return checkKeyword("or", TOKEN_OR)
	case 'p':
		return checkKeyword("print", TOKEN_PRINT)
	case 'r':
		return checkKeyword("return", TOKEN_RETURN)
	case 's':
		return checkKeyword("super", TOKEN_SUPER)
	case 'v':
		return checkKeyword("var", TOKEN_VAR)
	case 'w':
		return checkKeyword("while", TOKEN_WHILE)
	case 'f':
		if _scanner.SourceCurrent_idx-_scanner.SourceStart_idx > 1 {
			switch _scanner.Source[_scanner.SourceStart_idx+1] {
			case 'a':
				return checkKeyword("false", TOKEN_FALSE)
			case 'o':
				return checkKeyword("for", TOKEN_FOR)
			case 'u':
				return checkKeyword("fun", TOKEN_FUN)
			}
		}
	case 't':
		if _scanner.SourceCurrent_idx-_scanner.SourceStart_idx > 1 {
			switch _scanner.Source[_scanner.SourceStart_idx+1] {
			case 'h':
				return checkKeyword("this", TOKEN_THIS)
			case 'r':
				return checkKeyword("true", TOKEN_TRUE)
			}
		}
	}
	return TOKEN_IDENTIFIER
}

func checkKeyword(word string, tok TokenType) TokenType {
	if _scanner.Source[_scanner.SourceStart_idx:(_scanner.SourceCurrent_idx)] == word {
		return tok
	}
	return TOKEN_IDENTIFIER
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func number() Token {
	for isDigit(peek()) {
		advance()
	}

	if peek() == '.' && isDigit(peekNext()) {
		advance()

		for isDigit(peek()) {
			advance()
		}
	}
	return makeToken(TOKEN_NUMBER)
}

func makeString() Token {
	for peek() != '"' && !isAtEnd() {
		if peek() == '\n' {
			_scanner.Line += 1
		}
		advance()
	}
	if isAtEnd() {
		return errorToken("unterminated string")
	}
	advance()
	return makeToken(TOKEN_STRING)
}

func peekNext() rune {
	if isAtEnd() {
		return 0
	}
	return rune(_scanner.Source[_scanner.SourceCurrent_idx+1])
}

func peek() rune {
	return rune(_scanner.Source[_scanner.SourceCurrent_idx])
}

func matchIfElse(expected rune, ttIf TokenType, ttElse TokenType) TokenType {
	if isAtEnd() {
		return ttElse
	}
	if rune(_scanner.Source[_scanner.SourceCurrent_idx]) != expected {
		return ttElse
	}
	_scanner.SourceCurrent_idx += 1
	return ttIf
}

func advance() rune {
	if isAtEnd() {
		return 0
	}
	_scanner.SourceCurrent_idx += 1
	return rune(_scanner.Source[_scanner.SourceCurrent_idx-1])
}

func isAtEnd() bool {
	// calling `len` each time this is run, instead of storing the length of the Source (which never changes)
	// isn't as efficient, but i'm trying to keep this code as close to the c iimplementation as i can
	return len(_scanner.Source) == _scanner.SourceCurrent_idx+1
}

func makeToken(_type TokenType) Token {
	return Token{Type: _type,
		Value: _scanner.Source[_scanner.SourceStart_idx : _scanner.SourceCurrent_idx+1],
		Line:  _scanner.Line}
}

func errorToken(msg string) Token {
	return Token{Type: TOKEN_ERROR, Value: msg, Line: _scanner.Line}
}
