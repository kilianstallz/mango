package cmd

import (
	"testing"
)

func TestSorting(t *testing.T) {

	list := []string{"go1.20.2", "go1.20.1", "go1.20", "go1.21.1", "go1.21.2", "go1.21.3", "go1.22.1", "go1.22.3", "go1.22.2"}

	lastExpected := "go1.22.3"

	last := getLatestVersion(list)
	if last != lastExpected {
		t.Errorf("Expected %s, got %s", lastExpected, last)
	}
}

func TestQuery(t *testing.T) {
	queryLatestVersion()
}
