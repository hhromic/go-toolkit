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

// RunServer calls srv.ListenAndServe and waits for the context to be done.
// When the context is done, it gracefully shuts down the server with a timeout.
func RunServer(ctx context.Context, srv *http.Server, timeout time.Duration) error {
	done := make(chan struct{}, 1)
	err := make(chan error, 1)

	go waitAndShutdown(ctx, srv, timeout, done, err)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve: %w", err)
	}

	<-done

	if err := <-err; err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	return nil
}

// RunServerTLS calls srv.ListenAndServeTLS and waits for the context to be done.
// When the context is done, it gracefully shuts down the server with a timeout.
func RunServerTLS(ctx context.Context, srv *http.Server, certFile, keyFile string, timeout time.Duration) error {
	done := make(chan struct{}, 1)
	err := make(chan error, 1)

	go waitAndShutdown(ctx, srv, timeout, done, err)

	if err := srv.ListenAndServeTLS(certFile, keyFile); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve TLS: %w", err)
	}

	<-done

	if err := <-err; err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	return nil
}

func waitAndShutdown(
	ctx context.Context,
	srv *http.Server,
	timeout time.Duration,
	done chan<- struct{},
	err chan<- error,
) {
	defer close(err)
	defer close(done)

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err <- srv.Shutdown(ctx) //nolint:contextcheck
}
