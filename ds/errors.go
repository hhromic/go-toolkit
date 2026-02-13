// SPDX-FileCopyrightText: Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package ds

import "errors"

// Errors used by the ds package.
var (
	// ErrUnknownFormat is returned when an unknown format is used.
	ErrUnknownFormat = errors.New("unknown format")
)
