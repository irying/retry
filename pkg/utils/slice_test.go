package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type HasClone struct {
	Filed1 string
	Filed2 string
}

func (h *HasClone) Clone() (clone HasClone) {
	clone = *h
	return
}

func TestArrayShouldBeCloned(t *testing.T) {
	src := []HasClone{
		HasClone{
			"field1_1",
			"field1_2",
		},
		HasClone{
			"field2_1",
			"field2_2",
		},
	}

	dest := ArrayClone(src)
	destHasClone, ok := dest.([]HasClone)
	if !ok {
		t.Fatalf("dest could not conversion to []HasClone")
	}

	destHasClone[0].Filed1 = "field1_1 has Modified"
	assert.NotEqual(t, destHasClone, src,
		"DesHasClone %+v should not deep equal with srcHasClone +%v",
		destHasClone, src)
}
