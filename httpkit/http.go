// SPDX-FileCopyrightText: Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package httpkit

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
	doneCh := make(chan struct{}, 1)
	errCh := make(chan error, 1)

	go waitAndShutdown(ctx, srv, timeout, doneCh, errCh)

	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve: %w", err)
	}

	<-doneCh

	err = <-errCh
	if err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	return nil
}

// RunServerTLS calls srv.ListenAndServeTLS and waits for the context to be done.
// When the context is done, it gracefully shuts down the server with a timeout.
func RunServerTLS(
	ctx context.Context,
	srv *http.Server,
	certFile, keyFile string,
	timeout time.Duration,
) error {
	doneCh := make(chan struct{}, 1)
	errCh := make(chan error, 1)

	go waitAndShutdown(ctx, srv, timeout, doneCh, errCh)

	err := srv.ListenAndServeTLS(certFile, keyFile)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve TLS: %w", err)
	}

	<-doneCh

	err = <-errCh
	if err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	return nil
}

func waitAndShutdown(
	ctx context.Context,
	srv *http.Server,
	timeout time.Duration,
	doneCh chan<- struct{},
	errCh chan<- error,
) {
	defer close(errCh)
	defer close(doneCh)

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	errCh <- srv.Shutdown(ctx) //nolint:contextcheck
}
