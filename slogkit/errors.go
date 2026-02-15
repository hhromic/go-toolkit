// SPDX-FileCopyrightText: Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package slogkit

import "errors"

// Errors used by the slogkit package.
var (
	// ErrUnknownHandlerName is returned when an unknown slogkit handler name is used.
	ErrUnknownHandlerName = errors.New("unknown handler name")
)
