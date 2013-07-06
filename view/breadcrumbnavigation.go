// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package view

type BreadCrumbNavigation struct {
	Entries []*BreadCrumb
}

func (navigation *BreadCrumbNavigation) IsAvailable() bool {
	return len(navigation.Entries) > 0
}

type BreadCrumb struct {
	Level int
	Title string
	Path  string
}