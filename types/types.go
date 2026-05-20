// /home/krylon/go/src/github.com/blicero/capybara/types/types.go
// -*- mode: go; coding: utf-8; -*-
// Created on 20. 05. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-05-20 16:45:49 krylon>

// Package types provides symbolic constants to identify the types
// used in the Capybara language.
package types

//go:generate stringer -type=ID

type Type uint8

const (
	None Type = iota
	Integer
	Float
	String
	List
	Map
	Object
)
