package event

import (
	"testing"
)

func TestGenerateEvents(t *testing.T) {
	es := generateEvents()
	t.Logf("%v\n", es)
}
