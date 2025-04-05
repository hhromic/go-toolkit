// SPDX-FileCopyrightText: Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package slog_test

import (
	"fmt"
	"log/slog"
	"os"

	tkslog "github.com/hhromic/go-toolkit/slog"
)

func ExampleHandler_String() {
	h := tkslog.HandlerText
	fmt.Println(h.String())
	// Output: text
}

//nolint:staticcheck
func ExampleHandler_MarshalText() {
	h := tkslog.HandlerText

	b, err := h.MarshalText()
	if err != nil {
		panic(err)
	}

	fmt.Println(b)
	// Output: [116 101 120 116]
}

func ExampleHandler_UnmarshalText() {
	b := []byte{116, 101, 120, 116}

	var h tkslog.Handler
	if err := h.UnmarshalText(b); err != nil {
		panic(err)
	}

	fmt.Println(h.String())
	// Output: text
}

//nolint:testableexamples
func ExampleNewSlogLogger() {
	h := tkslog.HandlerText
	l := slog.LevelDebug
	logger := tkslog.NewSlogLogger(os.Stdout, h, l)

	version := "1.2.3"
	logger.Info("application started", "version", version)
}

//nolint:testableexamples
func ExampleNewSlogLogger_setDefault() {
	h := tkslog.HandlerText
	l := slog.LevelDebug
	logger := tkslog.NewSlogLogger(os.Stdout, h, l)

	slog.SetDefault(logger)

	version := "1.2.3"
	slog.Info("application started", "version", version)
}
