package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepo(t *testing.T) {

	s := "* [asciigraph](https://github.com/guptarohit/asciigraph) - Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies."
	re := getRepo(s)
	assert.Equal(t, "guptarohit/asciigraph", re)

	s = "* [dummy](https://github.com/user/repo) - Description containing https://github.com/user/repo"
	re = getRepo(s)
	assert.Equal(t, "user/repo", re)

	s = "* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) - A golang package to communicate with HipChat over XMPP."
	re = getRepo(s)
	assert.Equal(t, "daneharrigan/hipchat", re)

	s = "* [dummy](https://github.com/user/repo/) - Description containing https://github.com/user/repo"
	re = getRepo(s)
	assert.Equal(t, "user/repo", re)

}
