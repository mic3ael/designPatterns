package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) getPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}
	}

	return d.dummyData[name]
}

func Test1(t *testing.T) {
	names := []string{"a", "b"}
	expected := 3
	actual := getTotalPopulation(&DummyDatabase{}, names)
	assert.Equal(t, actual, expected, "they should be equal")
}
