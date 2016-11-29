package string_h

import (
	"testing"
)

func Test_Include(t *testing.T) {
	ary := []string{"a", "b"}

	indexMap := map[string]bool{
		"a": true,
		"b": true,
		"x": false,
	}

	for item, found := range indexMap {
		if Include(item, ary) != found {
			t.Error("Item(%v, ary) != %v", item, found)
		}
	}
}
