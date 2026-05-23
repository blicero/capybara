// /home/krylon/go/src/github.com/blicero/capybara/parser/parser.go
// -*- mode: go; coding: utf-8; -*-
// Created on 20. 05. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-05-23 16:28:40 krylon>

package parser

import (
	"fmt"
	"strconv"

	"github.com/alecthomas/participle/v2/lexer"
	"github.com/blicero/capybara/types"
)

var lex = lexer.MustSimple([]lexer.SimpleRule{
	{Name: `Symbol`, Pattern: `[-+*/%:a-zA-Z<>][-+*/%:a-zA-Z\d<>]*`},
	{Name: `Integer`, Pattern: `\d+`},
	{Name: `String`, Pattern: `"(?:[^\"]*)"`},
	{Name: `OpenParen`, Pattern: `\(`},
	{Name: `CloseParen`, Pattern: `\)`},
	{Name: `OpenBrace`, Pattern: `\{`},
	{Name: `CloseBrace`, Pattern: `\}`},
	{Name: `OpenBracket`, Pattern: `\[`},
	{Name: `CloseBracket`, Pattern: `\]`},
	{Name: `Semicolon`, Pattern: `;`},
	{Name: `Blank`, Pattern: `\s+`},
	{Name: `Comma`, Pattern: `,`},
})

type Value interface {
	fmt.Stringer
	Type() types.Type
}

type Symbol struct {
	Pos lexer.Position
	Sym string `parser:"@Symbol"`
}

func (s Symbol) String() string   { return s.Sym }
func (s Symbol) Type() types.Type { return types.Symbol }
func (s Symbol) IsKeyword() bool  { return s.Sym[0] == ':' }

func (s Symbol) Equal(other Value) bool {
	switch val := other.(type) {
	case Symbol:
		return s.Sym == val.Sym
	default:
		return false
	}
} // func (s Symbol) Equal(other Value)

type Integer struct {
	Pos lexer.Position
	Int int64 `parser:"@Integer"`
}

func (i Integer) Type() types.Type { return types.Integer }
func (i Integer) String() string   { return strconv.FormatInt(i.Int, 10) }
func (i Integer) Equal(other Value) bool {
	switch val := other.(type) {
	case Integer:
		return i.Int == val.Int
	default:
		return false
	}
} // func (i Integer) Equal(other Value) bool

type Float struct {
	Pos lexer.Position
	Num float64 `parser:"@Float"`
}

func (f Float) Type() types.Type { return types.Float }
func (f Float) String() string   { return strconv.FormatFloat(f.Num, 'f', -1, 64) }

func (f Float) Equal(other Value) bool {
	switch val := other.(type) {
	case Float:
		return f.Num == val.Num
	default:
		return false
	}
} // func (f Float) Equal(other Value) bool

type String struct {
	Pos lexer.Position
	Str string
}

func (s String) Type() types.Type { return types.String }
func (s String) String() string   { return s.Str }

func (s String) Equal(other Value) bool {
	switch val := other.(type) {
	case String:
		return s.Str == val.Str
	default:
		return false
	}
} // func (s String) Equal(other Value) bool

type List struct {
	Pos lexer.Position
}
