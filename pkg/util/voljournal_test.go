/*
Copyright 2020 Ceph-CSI authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"testing"
)

func TestCSIJournalGetNameForUUID(t *testing.T) {
	var name string
	uuid := "646f69f2-2302-11ea-b32b-e27a2df59041"

	// define test cases
	tests := []struct {
		input  string
		output string
		isSnap bool
	}{
		// journal without prefix should fallback to csi-vol-
		{"", "csi-vol-646f69f2-2302-11ea-b32b-e27a2df59041", false},
		{"", "csi-snap-646f69f2-2302-11ea-b32b-e27a2df59041", true},
		// we override the static naming prefix
		{"foo-bar-", "foo-bar-646f69f2-2302-11ea-b32b-e27a2df59041", false},
		// we use a template naming prefix
		{"pvc-{{ printf \"%.8s\" .Hash}}-{{.UUID}}", "pvc-233066e7-646f69f2-2302-11ea-b32b-e27a2df59041", false},
	}

	// run tests
	for _, test := range tests {
		journal := NewCSIVolumeJournal()
		name = journal.GetNameForUUID(test.input, uuid, test.isSnap)
		if name != test.output {
			t.Errorf("expected %s, got %s", test.output, name)
		}
	}
}
