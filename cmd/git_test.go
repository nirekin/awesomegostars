package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortingValue(t *testing.T) {
	r := Response{name: "aaa", Star: 101, Watch: 102, Fork: 103, Issues: 104}
	i := r.sortingValue(keyStart)
	assert.Equal(t, i, 101)
	i = r.sortingValue(keyWatch)
	assert.Equal(t, i, 102)
	i = r.sortingValue(keyFork)
	assert.Equal(t, i, 103)
	i = r.sortingValue(keyIssues)
	assert.Equal(t, i, 104)
}
