// /home/krylon/go/src/github.com/blicero/capybara/parser/parser.go
// -*- mode: go; coding: utf-8; -*-
// Created on 20. 05. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-05-20 16:37:38 krylon>

package parser

import "github.com/alecthomas/participle/v2/lexer"

var lex = lexer.MustSimple([]lexer.SimpleRule{
	{Name: `Symbol`, Pattern: `[-+*/%:a-zA-Z<>][-+*/%:a-zA-Z\d<>]*`},
	{Name: `Integer`, Pattern: `\d+`},
	{Name: `String`, Pattern: `"(?:[^\"]*)"`},
	{Name: `OpenParen`, Pattern: `\(`},
	{Name: `CloseParen`, Pattern: `\)`},
	{Name: `OpenBrace`, Pattern: `\{`},
	{Name: `CloseBrace`, Pattern: `\}`},
	{Name: `Semicolon`, Pattern: `;`},
	{Name: `Blank`, Pattern: `\s+`},
})
