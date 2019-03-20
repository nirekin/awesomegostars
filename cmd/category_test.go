package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategory(t *testing.T) {

	c := "   category   "
	ca := getCategory(c)
	assert.Equal(t, "category", ca)

	c = "CATEGORY"
	ca = getCategory(c)
	assert.Equal(t, "category", ca)

	c = "CATEGORY1 CATEGORY2 CATEGORY3"
	ca = getCategory(c)
	assert.Equal(t, "category1-category2-category3", ca)

}
