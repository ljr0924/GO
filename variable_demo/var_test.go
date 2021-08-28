package variable_demo

import (
	"testing"
)

func TestUint32Time(t *testing.T) {

	s := uint64(1627660800)

	t.Logf("%d", s * 1000)

}
