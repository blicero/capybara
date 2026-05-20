// /home/krylon/go/src/github.com/blicero/capybara/parser/parser_test.go
// -*- mode: go; coding: utf-8; -*-
// Created on 20. 05. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-05-20 12:33:32 krylon>

package parser

import "testing"

func TestSimple(t *testing.T) {
	expr := "abobo 7 255"
	capy := &Parser{Buffer: expr}
	capy.Init()
	if err := capy.Parse(); err != nil {
		t.Fatal(err)
	}

	capy.PrintSyntaxTree()
}
