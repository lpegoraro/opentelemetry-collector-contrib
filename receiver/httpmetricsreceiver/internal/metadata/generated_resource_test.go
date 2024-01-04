// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceBuilder(t *testing.T) {
	for _, test := range []string{"default", "all_set", "none_set"} {
		t.Run(test, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, test)
			rb := NewResourceBuilder(cfg)
			rb.SetTags(map[string]any{"key1": "tags-val1", "key2": "tags-val2"})

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return empty Resource

			switch test {
			case "default":
				assert.Equal(t, 0, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 1, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", test)
			}

			val, ok := res.Attributes().Get("tags")
			assert.Equal(t, test == "all_set", ok)
			if ok {
				assert.EqualValues(t, map[string]any{"key1": "tags-val1", "key2": "tags-val2"}, val.Map().AsRaw())
			}
		})
	}
}
