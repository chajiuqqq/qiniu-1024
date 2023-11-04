package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testMapStruct struct {
	ID   int
	Name string
}

func (m testMapStruct) Key() any {
	return m.ID
}
func TestToMap(t *testing.T) {
	d := []testMapStruct{
		{
			ID:   1,
			Name: "test1",
		},
		{
			ID:   2,
			Name: "test2",
		},
	}
	m := ToMap(d)
	assert.Equal(t, 2, len(m))
	assert.Equal(t, 1, m[1].ID)
	assert.Equal(t, "test1", m[1].Name)
	assert.Equal(t, 2, m[2].ID)
	assert.Equal(t, "test2", m[2].Name)
}
