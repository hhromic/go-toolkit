// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// WaitAndShutdown waits for a [context.Context] to be done and shuts down an [http.Server] with a timeout.
func WaitAndShutdown(ctx context.Context, srv *http.Server, timeout time.Duration) error {
	<-ctx.Done()
	errs := []error{ctx.Err()}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil { //nolint:contextcheck
		errs = append(errs, fmt.Errorf("shutdown: %w", err))
	}

	return errors.Join(errs...)
}
