package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	assert := assert.New(t)

	var set = make(Set)
	set.Add("Australia")
	assert.True(set.Contains("Australia"))
	assert.False(set.Contains("New Zealand"))

	set.Add("Australia", "New Zealand")
	assert.True(set.Contains("Australia"))
	assert.True(set.Contains("New Zealand"))
}

func TestSize(t *testing.T) {
	assert := assert.New(t)

	var set = make(Set)
	set.Add("Australia")
	set.Add("New Zealand")
	assert.Equal(2, len(set))
}
