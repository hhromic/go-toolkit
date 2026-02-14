// SPDX-FileCopyrightText: Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package slogkit_test

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/hhromic/go-toolkit/slogkit"
)

func ExampleHandler_String() {
	h := slogkit.HandlerText
	fmt.Println(h.String())
	// Output: text
}

//nolint:staticcheck
func ExampleHandler_MarshalText() {
	h := slogkit.HandlerText

	b, err := h.MarshalText()
	if err != nil {
		panic(err)
	}

	fmt.Println(b)
	// Output: [116 101 120 116]
}

func ExampleHandler_UnmarshalText() {
	b := []byte{116, 101, 120, 116}

	var hdl slogkit.Handler

	err := hdl.UnmarshalText(b)
	if err != nil {
		panic(err)
	}

	fmt.Println(hdl.String())
	// Output: text
}

//nolint:testableexamples
func ExampleNewSlogLogger() {
	h := slogkit.HandlerText
	l := slog.LevelDebug
	logger := slogkit.NewSlogLogger(os.Stdout, h, l)

	version := "1.2.3"
	logger.Info("application started", "version", version)
}

//nolint:testableexamples
func ExampleNewSlogLogger_setDefault() {
	h := slogkit.HandlerText
	l := slog.LevelDebug
	logger := slogkit.NewSlogLogger(os.Stdout, h, l)

	slog.SetDefault(logger)

	version := "1.2.3"
	slog.Info("application started", "version", version)
}
