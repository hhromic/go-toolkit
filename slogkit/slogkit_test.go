// SPDX-FileCopyrightText: Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package slogkit_test

import (
	"bytes"
	"context"
	"log/slog"
	"regexp"
	"testing"

	"github.com/hhromic/go-toolkit/slogkit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlerString(t *testing.T) {
	testCases := []struct {
		name    string
		handler slogkit.Handler
		want    string
	}{
		{
			name:    "HandlerText",
			handler: slogkit.HandlerText,
			want:    "text",
		},
		{
			name:    "HandlerJSON",
			handler: slogkit.HandlerJSON,
			want:    "json",
		},
		{
			name:    "HandlerTint",
			handler: slogkit.HandlerTint,
			want:    "tint",
		},
		{
			name:    "HandlerAuto",
			handler: slogkit.HandlerAuto,
			want:    "auto",
		},
		{
			name:    "InvalidHandler",
			handler: -1,
			want:    "",
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			assert.Equal(t, tCase.want, tCase.handler.String())
		})
	}
}

func TestHandlerMarshalText(t *testing.T) {
	testCases := []struct {
		name    string
		handler slogkit.Handler
		want    []byte
		wantErr error
	}{
		{
			name:    "HandlerText",
			handler: slogkit.HandlerText,
			want:    []byte("text"),
			wantErr: nil,
		},
		{
			name:    "HandlerJSON",
			handler: slogkit.HandlerJSON,
			want:    []byte("json"),
			wantErr: nil,
		},
		{
			name:    "HandlerTint",
			handler: slogkit.HandlerTint,
			want:    []byte("tint"),
			wantErr: nil,
		},
		{
			name:    "HandlerAuto",
			handler: slogkit.HandlerAuto,
			want:    []byte("auto"),
			wantErr: nil,
		},
		{
			name:    "InvalidHandler",
			handler: -1,
			want:    []byte(""),
			wantErr: nil,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			b, err := tCase.handler.MarshalText()
			require.ErrorIs(t, err, tCase.wantErr)

			if tCase.wantErr == nil {
				assert.Equal(t, tCase.want, b)
			}
		})
	}
}

func TestHandlerUnmarshalText(t *testing.T) {
	testCases := []struct {
		name    string
		b       []byte
		want    slogkit.Handler
		wantErr error
	}{
		{
			name:    "HandlerText",
			b:       []byte("text"),
			want:    slogkit.HandlerText,
			wantErr: nil,
		},
		{
			name:    "HandlerJSON",
			b:       []byte("json"),
			want:    slogkit.HandlerJSON,
			wantErr: nil,
		},
		{
			name:    "HandlerTint",
			b:       []byte("tint"),
			want:    slogkit.HandlerTint,
			wantErr: nil,
		},
		{
			name:    "HandlerAuto",
			b:       []byte("auto"),
			want:    slogkit.HandlerAuto,
			wantErr: nil,
		},
		{
			name:    "InvalidHandler",
			b:       []byte("foobar"),
			want:    slogkit.HandlerText,
			wantErr: slogkit.ErrUnknownHandlerName,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			var hdl slogkit.Handler

			err := hdl.UnmarshalText(tCase.b)
			require.ErrorIs(t, err, tCase.wantErr)

			if tCase.wantErr == nil {
				assert.Equal(t, tCase.want, hdl)
			}
		})
	}
}

//nolint:funlen
func TestNewLogger(t *testing.T) {
	testCases := []struct {
		name     string
		handler  slogkit.Handler
		leveler  slog.Level
		logLevel slog.Level
		logMsg   string
		logArgs  []any
		want     *regexp.Regexp
	}{
		{
			name:     "HandlerText-Debug-Debug",
			handler:  slogkit.HandlerText,
			leveler:  slog.LevelDebug,
			logLevel: slog.LevelDebug,
			logMsg:   "message",
			logArgs:  []any{"key1", "val1", "key2", "val2"},
			want:     regexp.MustCompile(`^ts=.+ level=DEBUG msg=message key1=val1 key2=val2\n$`),
		},
		{
			name:     "HandlerJSON-Debug-Info",
			handler:  slogkit.HandlerJSON,
			leveler:  slog.LevelDebug,
			logLevel: slog.LevelInfo,
			logMsg:   "message",
			logArgs:  []any{"key1", "val1", "key2", "val2"},
			want: regexp.MustCompile(
				`^{"ts":".+","level":"INFO","msg":"message","key1":"val1","key2":"val2"}\n$`,
			),
		},
		{
			name:     "HandlerTint-Debug-Warn",
			handler:  slogkit.HandlerTint,
			leveler:  slog.LevelDebug,
			logLevel: slog.LevelWarn,
			logMsg:   "message",
			logArgs:  []any{"key1", "val1", "key2", "val2"},
			want: regexp.MustCompile(
				`^\x1b\[2m.+\x1b\[0m \x1b\[93mWRN\x1b\[0m message \x1b\[2mkey1=\x1b\[0mval1 \x1b\[2mkey2=\x1b\[0mval2\n$`,
			),
		},
		{
			name:     "HandlerAuto-Debug-Error",
			handler:  slogkit.HandlerAuto,
			leveler:  slog.LevelDebug,
			logLevel: slog.LevelError,
			logMsg:   "message",
			logArgs:  []any{"key1", "val1", "key2", "val2"},
			want:     regexp.MustCompile(`^ts=.+ level=ERROR msg=message key1=val1 key2=val2\n$`),
		},
		{
			name:     "HandlerText-Warn-Info",
			handler:  slogkit.HandlerText,
			leveler:  slog.LevelWarn,
			logLevel: slog.LevelInfo,
			logMsg:   "message",
			logArgs:  []any{"key1", "val1", "key2", "val2"},
			want:     regexp.MustCompile(`^$`),
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			var buf bytes.Buffer

			l := slogkit.NewLogger(&buf, tCase.handler, tCase.leveler)
			require.NotNil(t, l)

			l.Log(context.Background(), tCase.logLevel, tCase.logMsg, tCase.logArgs...)
			assert.Regexp(t, tCase.want, buf.String())
		})
	}

	t.Run("InvalidHandler", func(t *testing.T) {
		var buf bytes.Buffer

		l := slogkit.NewLogger(&buf, -1, slog.LevelDebug)
		assert.Nil(t, l)
	})
}
