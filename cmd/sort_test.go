package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadSortingKey(t *testing.T) {
	ts := make([]Response, 0)
	_, err := sortResponses("dummy_key", ts)
	assert.NotNil(t, err)
}

func TestSortOnStars(t *testing.T) {
	r1 := build("aaa", 9, 0, 0, 0)
	r2 := build("aac", 9, 0, 0, 0)
	r3 := build("bbb", 1, 0, 0, 0)
	r4 := build("ccc", 7, 0, 0, 0)
	r5 := build("ddd", 5, 0, 0, 0)
	r6 := build("eee", 2, 0, 0, 0)
	r7 := build("fff", 8, 0, 0, 0)
	r8 := build("hhh", 4, 0, 0, 0)
	r9 := build("iii", 3, 0, 0, 0)
	r10 := build("aab", 9, 0, 0, 0)
	r11 := build("jjj", 6, 0, 0, 0)
	ts := builToSort(r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11)

	r, err := sortResponses(keyStar, ts)
	assert.Nil(t, err)
	assert.Equal(t, len(r), len(ts))
	check(t, r, r1, r10, r2, r7, r4, r11, r5, r8, r9, r6, r3)
}

func TestSortOnWatch(t *testing.T) {
	r1 := build("aaa", 0, 9, 0, 0)
	r2 := build("aac", 0, 9, 0, 0)
	r3 := build("bbb", 0, 1, 0, 0)
	r4 := build("ccc", 0, 7, 0, 0)
	r5 := build("ddd", 0, 5, 0, 0)
	r6 := build("eee", 0, 2, 0, 0)
	r7 := build("fff", 0, 8, 0, 0)
	r8 := build("hhh", 0, 4, 0, 0)
	r9 := build("iii", 0, 3, 0, 0)
	r10 := build("aab", 0, 9, 0, 0)
	r11 := build("jjj", 0, 6, 0, 0)
	ts := builToSort(r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11)

	r, err := sortResponses(keyWatch, ts)
	assert.Nil(t, err)
	assert.Equal(t, len(r), len(ts))
	check(t, r, r1, r10, r2, r7, r4, r11, r5, r8, r9, r6, r3)
}

func TestSortOnFork(t *testing.T) {
	r1 := build("aaa", 0, 0, 9, 0)
	r2 := build("aac", 0, 0, 9, 0)
	r3 := build("bbb", 0, 0, 1, 0)
	r4 := build("ccc", 0, 0, 7, 0)
	r5 := build("ddd", 0, 0, 5, 0)
	r6 := build("eee", 0, 0, 2, 0)
	r7 := build("fff", 0, 0, 8, 0)
	r8 := build("hhh", 0, 0, 4, 0)
	r9 := build("iii", 0, 0, 3, 0)
	r10 := build("aab", 0, 0, 9, 0)
	r11 := build("jjj", 0, 0, 6, 0)
	ts := builToSort(r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11)

	r, err := sortResponses(keyFork, ts)
	assert.Nil(t, err)
	assert.Equal(t, len(r), len(ts))
	check(t, r, r1, r10, r2, r7, r4, r11, r5, r8, r9, r6, r3)
}

func TestSortOnIssues(t *testing.T) {
	r1 := build("aaa", 0, 0, 0, 9)
	r2 := build("aac", 0, 0, 0, 9)
	r3 := build("bbb", 0, 0, 0, 1)
	r4 := build("ccc", 0, 0, 0, 7)
	r5 := build("ddd", 0, 0, 0, 5)
	r6 := build("eee", 0, 0, 0, 2)
	r7 := build("fff", 0, 0, 0, 8)
	r8 := build("hhh", 0, 0, 0, 4)
	r9 := build("iii", 0, 0, 0, 3)
	r10 := build("aab", 0, 0, 0, 9)
	r11 := build("jjj", 0, 0, 0, 6)
	ts := builToSort(r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11)

	r, err := sortResponses(keyIssues, ts)
	assert.Nil(t, err)
	assert.Equal(t, len(r), len(ts))
	check(t, r, r1, r10, r2, r7, r4, r11, r5, r8, r9, r6, r3)
}

func builToSort(s ...Response) []Response {
	r := make([]Response, 0)
	for _, v := range s {
		r = append(r, v)
	}
	return r
}

func check(t *testing.T, resp []Response, wanted ...Response) {
	assert.Equal(t, len(wanted), len(resp))
	for i, w := range wanted {
		r := resp[i]
		assert.Equal(t, w.name, r.name)
		assert.Equal(t, w.Star, r.Star)
		assert.Equal(t, w.Fork, r.Fork)
		assert.Equal(t, w.Watch, r.Watch)
		assert.Equal(t, w.Issues, r.Issues)
	}
}

func build(n string, s, w, f, i int) Response {
	return Response{name: n, Star: s, Watch: w, Fork: f, Issues: i}

}
