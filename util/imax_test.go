package util_test

import (
	"testing"

	"github.com/kooltuoehias/gtree/util"
)

func TestMathMax(t *testing.T) {
	if util.Max(2, 3) != 3 {
		t.Errorf("3 instead of %d shall be max value from 2 and 3", util.Max(2, 3))
	}

}
