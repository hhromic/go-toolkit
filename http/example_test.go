// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package http_test

import (
	"context"
	"net/http"
	"time"

	tkhttp "github.com/hhromic/go-toolkit/http"
)

//nolint:exhaustruct,gosec,testableexamples
func ExampleRunServer() {
	ctx := context.Background()
	srv := &http.Server{}
	shutdownTimeout := 30 * time.Second

	if err := tkhttp.RunServer(ctx, srv, shutdownTimeout); err != nil {
		panic(err)
	}
}

//nolint:exhaustruct,gosec,testableexamples
func ExampleRunServerTLS() {
	ctx := context.Background()
	srv := &http.Server{}
	certFile := "/path/to/server.crt"
	keyFile := "/path/to/server.key"
	shutdownTimeout := 30 * time.Second

	if err := tkhttp.RunServerTLS(ctx, srv, certFile, keyFile, shutdownTimeout); err != nil {
		panic(err)
	}
}
