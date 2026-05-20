// /home/krylon/go/src/github.com/blicero/capybara/main.go
// -*- mode: go; coding: utf-8; -*-
// Created on 20. 05. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-05-20 11:35:38 krylon>

package main

import (
	"fmt"

	"github.com/blicero/capybara/common"
)

func main() {
	fmt.Printf("%s %s\n",
		common.AppName,
		common.Version)
	fmt.Println("Nothing to see here, move along...")
}
