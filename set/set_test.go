package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	assert := assert.New(t)

	var set = make(Set)
	set.Add("Australia")
	assert.Equal(1, len(set))
	assert.True(set.Contains("Australia"))
}

func TestAddSameElementTwice(t *testing.T) {
	assert := assert.New(t)

	var set = make(Set)
	set.Add("Australia")
	set.Add("New Zealand")
	set.Add("Australia")

	assert.Equal(2, len(set))
	assert.True(set.Contains("Australia"))
	assert.True(set.Contains("New Zealand"))
}
