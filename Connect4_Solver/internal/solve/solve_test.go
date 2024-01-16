package solve

import (
	"testing"
)

func TestSolve(t *testing.T) {
	got := GetBestColumns("4444443265")
	if got != "5" {
		t.Errorf("GetBestColumns(\"4444443265\") = %s; want \"5\"", got)
	}
}
