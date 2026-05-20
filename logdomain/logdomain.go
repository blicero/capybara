// /home/krylon/go/src/github.com/blicero/capybara/logdomain/logdomain.go
// -*- mode: go; coding: utf-8; -*-
// Created on 19. 05. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-05-20 11:34:16 krylon>

// Package logdomain defines symbolic constants to label the different pieces
// of the entire program that might want to write something to the log.
package logdomain

//go:generate stringer -type=ID

// ID identifies a part of the application that could emit log messages.
type ID uint8

const (
	Parser ID = iota
	Interpreter
)
