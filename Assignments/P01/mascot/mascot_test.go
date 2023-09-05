package mascot_test

import (
	"testing"

	"github.com/aelious/4143-PLC-Nagel/tree/main/Assignments/P01/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Smokey Bear" {
		t.Fatal("Beep boop, test failed. Fix your stuff >:-(")
	}
}
