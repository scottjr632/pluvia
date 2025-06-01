package testutils

import (
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func AssertStringEquals(t *testing.T, expected string, actual string) {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(expected, actual, false)

	diffPretty := dmp.DiffPrettyText(diffs)

	if expected != actual {
		t.Errorf("Strings did not match")

		t.Log(diffPretty)
		t.Logf("Diffs:\n %v", diffs)
	}
}
