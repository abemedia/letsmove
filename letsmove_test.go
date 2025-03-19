//go:build darwin
// +build darwin

package letsmove_test

import (
	"testing"

	"github.com/abemedia/letsmove"
)

func TestIsMoveInProgress(t *testing.T) {
	ok := letsmove.IsMoveInProgress()
	if ok {
		t.Error("should not be in progress")
	}
}
