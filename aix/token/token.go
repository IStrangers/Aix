package token

import (
	"strconv"
	"unicode/utf8"
)

type Token int

const (
	_ Token = iota

	ILLEGAL
	EOF
	COMMENT
	IDENTIFIER
	KEYWORD

	//literal_beg
	STRING
	NUMBER
	BOOLEAN
	NULL
	//literal_end

	//operator_beg
	PLUS      // +
	MINUS     // -
	MULTIPLY  // *
	EXPONENT  // **
	SLASH     // /
	REMAINDER // %

	AND                  // &
	OR                   // |
	EXCLUSIVE_OR         // ^
	SHIFT_LEFT           // <<
	SHIFT_RIGHT          // >>
	UNSIGNED_SHIFT_RIGHT // >>>

	ADD_ASSIGN       // +=
	SUBTRACT_ASSIGN  // -=
	MULTIPLY_ASSIGN  // *=
	EXPONENT_ASSIGN  // **=
	QUOTIENT_ASSIGN  // /=
	REMAINDER_ASSIGN // %=

	AND_ASSIGN                  // &=
	OR_ASSIGN                   // |=
	EXCLUSIVE_OR_ASSIGN         // ^=
	SHIFT_LEFT_ASSIGN           // <<=
	SHIFT_RIGHT_ASSIGN          // >>=
	UNSIGNED_SHIFT_RIGHT_ASSIGN // >>>=

	LOGICAL_AND // &&
	LOGICAL_OR  // ||
	COALESCE    // ??
	INCREMENT   // ++
	DECREMENT   // --

	EQUAL        // ==
	STRICT_EQUAL // ===
	LESS         // <
	GREATER      // >
	ASSIGN       // =
	NOT          // !

	BITWISE_NOT // ~

	NOT_EQUAL        // !=
	STRICT_NOT_EQUAL // !==
	LESS_OR_EQUAL    // <=
	GREATER_OR_EQUAL // >=

	LEFT_PARENTHESIS // (
	LEFT_BRACKET     // [
	LEFT_BRACE       // {
	COMMA            // ,
	PERIOD           // .

	RIGHT_PARENTHESIS // )
	RIGHT_BRACKET     // ]
	RIGHT_BRACE       // }
	SEMICOLON         // ;
	COLON             // :
	QUESTION_MARK     // ?
	QUESTION_DOT      // ?.
	ARROW             // =>
	ELLIPSIS          // ...
	BACKTICK          // `
	//operator_end

	//keyword_beg
	IF
	IN
	OF
	DO
	VAR
	FOR
	NEW
	TRY
	THIS
	ELSE
	CASE
	CONST
	WHILE
	BREAK
	CATCH
	THROW
	CLASS
	SUPER
	RETURN
	SWITCH
	DEFAULT
	FINALLY
	EXTENDS
	FUN
	CONTINUE
	INSTANCEOF
	STATIC
	//keyword_end
)

var token2string = [...]string{
	ILLEGAL:                     "ILLEGAL",
	EOF:                         "EOF",
	COMMENT:                     "COMMENT",
	KEYWORD:                     "KEYWORD",
	STRING:                      "STRING",
	BOOLEAN:                     "BOOLEAN",
	NULL:                        "NULL",
	NUMBER:                      "NUMBER",
	IDENTIFIER:                  "IDENTIFIER",
	PLUS:                        "+",
	MINUS:                       "-",
	EXPONENT:                    "**",
	MULTIPLY:                    "*",
	SLASH:                       "/",
	REMAINDER:                   "%",
	AND:                         "&",
	OR:                          "|",
	EXCLUSIVE_OR:                "^",
	SHIFT_LEFT:                  "<<",
	SHIFT_RIGHT:                 ">>",
	UNSIGNED_SHIFT_RIGHT:        ">>>",
	ADD_ASSIGN:                  "+=",
	SUBTRACT_ASSIGN:             "-=",
	MULTIPLY_ASSIGN:             "*=",
	EXPONENT_ASSIGN:             "**=",
	QUOTIENT_ASSIGN:             "/=",
	REMAINDER_ASSIGN:            "%=",
	AND_ASSIGN:                  "&=",
	OR_ASSIGN:                   "|=",
	EXCLUSIVE_OR_ASSIGN:         "^=",
	SHIFT_LEFT_ASSIGN:           "<<=",
	SHIFT_RIGHT_ASSIGN:          ">>=",
	UNSIGNED_SHIFT_RIGHT_ASSIGN: ">>>=",
	LOGICAL_AND:                 "&&",
	LOGICAL_OR:                  "||",
	COALESCE:                    "??",
	INCREMENT:                   "++",
	DECREMENT:                   "--",
	EQUAL:                       "==",
	STRICT_EQUAL:                "===",
	LESS:                        "<",
	GREATER:                     ">",
	ASSIGN:                      "=",
	NOT:                         "!",
	BITWISE_NOT:                 "~",
	NOT_EQUAL:                   "!=",
	STRICT_NOT_EQUAL:            "!==",
	LESS_OR_EQUAL:               "<=",
	GREATER_OR_EQUAL:            ">=",
	LEFT_PARENTHESIS:            "(",
	LEFT_BRACKET:                "[",
	LEFT_BRACE:                  "{",
	COMMA:                       ",",
	PERIOD:                      ".",
	RIGHT_PARENTHESIS:           ")",
	RIGHT_BRACKET:               "]",
	RIGHT_BRACE:                 "}",
	SEMICOLON:                   ";",
	COLON:                       ":",
	QUESTION_MARK:               "?",
	QUESTION_DOT:                "?.",
	ARROW:                       "=>",
	ELLIPSIS:                    "...",
	BACKTICK:                    "`",
	IF:                          "if",
	IN:                          "in",
	OF:                          "of",
	DO:                          "do",
	VAR:                         "var",
	FOR:                         "for",
	NEW:                         "new",
	TRY:                         "try",
	THIS:                        "this",
	ELSE:                        "else",
	CASE:                        "case",
	CONST:                       "const",
	WHILE:                       "while",
	BREAK:                       "break",
	CATCH:                       "catch",
	THROW:                       "throw",
	CLASS:                       "class",
	SUPER:                       "super",
	RETURN:                      "return",
	SWITCH:                      "switch",
	STATIC:                      "static",
	DEFAULT:                     "default",
	FINALLY:                     "finally",
	EXTENDS:                     "extends",
	FUN:                         "fun",
	CONTINUE:                    "continue",
	INSTANCEOF:                  "instanceof",
}

func (tkn Token) String() string {
	if tkn >= 0 && tkn < Token(len(token2string)) {
		return token2string[tkn]
	}
	return "token(" + strconv.Itoa(int(tkn)) + ")"
}

func isCommonIdentifier(chr rune) bool {
	return chr == '$' || chr == '_' || chr == '\\' ||
		'a' <= chr && chr <= 'z' || 'A' <= chr && chr <= 'Z' ||
		chr >= utf8.RuneSelf
}

func isDecimalNumber(chr rune) bool {
	return '0' <= chr && chr <= '9'
}

func isIdentifierStart(chr rune) bool {
	return isCommonIdentifier(chr)
}

func isIdentifierPart(chr rune) bool {
	return isCommonIdentifier(chr) || isDecimalNumber(chr)
}

func IsIdentifier(name string) bool {
	if name == "" || IsKeyword(name) {
		return false
	}
	for i, c := range name {
		if i == 0 && !isIdentifierStart(c) {
			return false
		} else if !isIdentifierPart(c) {
			return false
		}
	}
	return true
}

func IsKeyword(name string) bool {
	if _, exists := keywordTable[name]; exists {
		return true
	}
	return false
}

type Keyword struct {
	token Token
}

var keywordTable = map[string]Keyword{
	"if": {
		token: IF,
	},
	"in": {
		token: IN,
	},
	"do": {
		token: DO,
	},
	"var": {
		token: VAR,
	},
	"for": {
		token: FOR,
	},
	"new": {
		token: NEW,
	},
	"try": {
		token: TRY,
	},
	"this": {
		token: THIS,
	},
	"else": {
		token: ELSE,
	},
	"case": {
		token: CASE,
	},
	"while": {
		token: WHILE,
	},
	"break": {
		token: BREAK,
	},
	"catch": {
		token: CATCH,
	},
	"throw": {
		token: THROW,
	},
	"return": {
		token: RETURN,
	},
	"switch": {
		token: SWITCH,
	},
	"default": {
		token: DEFAULT,
	},
	"finally": {
		token: FINALLY,
	},
	"fun": {
		token: FUN,
	},
	"continue": {
		token: CONTINUE,
	},
	"instanceof": {
		token: INSTANCEOF,
	},
	"const": {
		token: CONST,
	},
	"class": {
		token: CLASS,
	},
	"enum": {
		token: KEYWORD,
	},
	"extends": {
		token: EXTENDS,
	},
	"super": {
		token: SUPER,
	},
	"implements": {
		token: KEYWORD,
	},
	"interface": {
		token: KEYWORD,
	},
	"private": {
		token: KEYWORD,
	},
	"protected": {
		token: KEYWORD,
	},
	"public": {
		token: KEYWORD,
	},
	"static": {
		token: STATIC,
	},
	"false": {
		token: BOOLEAN,
	},
	"true": {
		token: BOOLEAN,
	},
	"null": {
		token: NULL,
	},
}
