// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package testenv

import (
	"syscall"
)

// Sigquit is the signal to send to kill a hanging subprocess.
// Send SIGQUIT to get a stack trace.
var Sigquit = syscall.SIGQUIT

func syscallIsNotSupported(err error) bool {
	// Removed by Hugo (not supported in Go 1.20)
	return false
}
