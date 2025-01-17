// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fsnotify implements file system notification.
package fsnotify

import "fmt"

// Op describes a set of file operations.
type Op uint32

// These are the file operations that can trigger a notification.
const (
	Create Op = 1 << iota
	Write
	Remove
	Rename
	Chmod
)

// Add starts watching for operations on the named file.
func (w *Watcher) Add(path string) error {
	return w.watch(path)
}

// Remove stops watching for operations on the named file.
func (w *Watcher) Remove(path string) error {
	return w.removeWatch(path)
}

// String formats the event e in the form
// "filename: REMOVE|WRITE|..."
func (e *Event) String() string {
	var events string = ""

	if e.Op&Create == Create {
		events += "|" + "CREATE"
	}

	if e.Op&Remove == Remove {
		events += "|" + "REMOVE"
	}

	if e.Op&Write == Write {
		events += "|" + "WRITE"
	}

	if e.Op&Rename == Rename {
		events += "|" + "RENAME"
	}

	if e.Op&Chmod == Chmod {
		events += "|" + "CHMOD"
	}

	if len(events) > 0 {
		events = events[1:]
	}

	return fmt.Sprintf("%q: %s", e.Name, events)
}
