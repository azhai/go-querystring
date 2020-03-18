// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package query

import (
	"strings"
)

// Change the type of tagOptions to a string

// tagOptions is the string following a comma in a struct field's "url" tag, or
// the empty string. It does not include the leading comma.
type tagOptions string

// parseTag splits a struct field's url tag into its name and comma-separated
// options.
func parseTag(tag string) (string, tagOptions) {
	s := strings.SplitN(tag, ",", 2)
	s = append(s, "", "") // padding
	return s[0], tagOptions(","+s[1]+",")
}

// Contains checks whether the tagOptions contains the specified option.
func (o tagOptions) Contains(option string) bool {
	return option != "" && strings.Contains(string(o), ","+option+",")
}
