// +build !debug

// Package spage ...
// Copyright (c) 2018 pirakansa
package spage

import (
	"net/http"
)

// Serve http
func Serve() {
	http.Handle("/", http.FileServer(Assets))
}
