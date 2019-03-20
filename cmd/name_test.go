package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {

	s := "* [asciigraph](https://github.com/guptarohit/asciigraph) - Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies."
	na := getName(s)
	assert.Equal(t, "asciigraph", na)

	s = "* [dummy](https://github.com/user/repo) - Description containing https://github.com/user/repo"
	na = getName(s)
	assert.Equal(t, "dummy", na)

}
