package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortingValue(t *testing.T) {
	r := Response{name: "aaa", Star: 101, Watch: 102, Fork: 103, Issues: 104}
	i := r.sortingValue(keyStar)
	assert.Equal(t, i, 101)
	i = r.sortingValue(keyWatch)
	assert.Equal(t, i, 102)
	i = r.sortingValue(keyFork)
	assert.Equal(t, i, 103)
	i = r.sortingValue(keyIssues)
	assert.Equal(t, i, 104)
}

func TestGitReadWholeList(t *testing.T) {
	md, err := readMD(masterURL)
	assert.Nil(t, err)
	assert.True(t, len(md) > 0)
}

func TestSortKey(t *testing.T) {

	r := Response{
		Star:   101,
		Watch:  102,
		Fork:   103,
		Issues: 104,
	}

	assert.Equal(t, r.sortingValue(keyStar), 101)
	assert.Equal(t, r.sortingValue(keyWatch), 102)
	assert.Equal(t, r.sortingValue(keyFork), 103)
	assert.Equal(t, r.sortingValue(keyIssues), 104)

}
