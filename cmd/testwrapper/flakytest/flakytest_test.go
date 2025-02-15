// Copyright (c) 2023 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flakytest

import "testing"

func TestIssueFormat(t *testing.T) {
	testCases := []struct {
		issue string
		want  bool
	}{
		{"https://github.com/tailscale/cOrp/issues/1234", true},
		{"https://github.com/otherproject/corp/issues/1234", false},
		{"https://github.com/tailscale/corp/issues/", false},
	}
	for _, testCase := range testCases {
		if issueRegexp.MatchString(testCase.issue) != testCase.want {
			ss := ""
			if !testCase.want {
				ss = " not"
			}
			t.Errorf("expected issueRegexp to%s match %q", ss, testCase.issue)
		}
	}
}
