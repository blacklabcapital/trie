package linkedlisttrie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testPhraseTrieNode() *PhraseTrieNode {
	return &PhraseTrieNode{key: "test", value: 1}
}

func testTrieFull() *PhraseTrieNode {
	m := map[string]int{
		"break":            1,
		"shooting":         2,
		"break out":        3,
		"break up":         4,
		"shooting up":      5,
		"break out nicely": 6,
		"r/g":              7,
	}

	return NewPhraseTrie(m)
}

func TestHasNext(t *testing.T) {
	n := testPhraseTrieNode()

	assert.False(t, n.HasNext())

	n.next = testPhraseTrieNode()

	assert.True(t, n.HasNext())
}

func TestIsLeaf(t *testing.T) {
	n := testPhraseTrieNode()

	assert.True(t, n.IsLeaf())

	n.children = testPhraseTrieNode()

	assert.False(t, n.IsLeaf())
}

func TestIsMember(t *testing.T) {
	// make empty root trie first
	trie := NewPhraseTrie(nil)

	// test phrases
	p1 := []string{"break"}
	p2 := []string{"shooting"}
	p3 := []string{"break", "out"}
	p4 := []string{"break", "up"}
	p5 := []string{"shooting", "up"}
	p6 := []string{"break", "out", "nicely"}
	p7 := []string{"r/g"}

	// root case
	member, value := trie.IsMember(p1)
	assert.False(t, member)

	// add p1
	n1 := testPhraseTrieNode()
	n1.key = p1[0]
	trie.children = n1

	member, value = trie.IsMember(p1)
	assert.True(t, member)
	assert.Equal(t, 1, value)

	// add p2
	n2 := testPhraseTrieNode()
	n2.key = p2[0]
	trie.children.next = n2

	member, value = trie.IsMember(p2)
	assert.True(t, member)
	assert.Equal(t, 1, value)

	// reset
	trie = NewPhraseTrie(nil)

	// add 2 word phrase p3
	n3 := testPhraseTrieNode()
	n3.key = p3[0]
	n3.children = testPhraseTrieNode()
	n3.children.key = p3[1]
	trie.children = n3

	member, value = trie.IsMember(p3)
	assert.True(t, member)
	assert.Equal(t, 1, value)

	// add same prefix new 2nd word phrase p4 to existing trie
	n4 := testPhraseTrieNode()
	n4.key = p4[1]
	trie.children.children.next = n4

	// check p3 just to make sure
	member, value = trie.IsMember(p3)
	assert.True(t, member)
	assert.Equal(t, 1, value)

	// check p4
	member, value = trie.IsMember(p4)
	assert.True(t, member)
	assert.Equal(t, 1, value)

	// add third entirely diff phrase to existing trie
	n5 := testPhraseTrieNode()
	n5.key = p5[0]
	n5.children = testPhraseTrieNode()
	n5.children.key = p5[1]
	trie.children.next = n5

	member, value = trie.IsMember(p5)
	assert.True(t, member)
	assert.Equal(t, 1, value)

	// reset
	trie = NewPhraseTrie(nil)

	// add 3 word phrase
	n6 := testPhraseTrieNode()
	n6.key = p6[0]
	n6.children = testPhraseTrieNode()
	n6.children.key = p6[1]
	n6.children.children = testPhraseTrieNode()
	n6.children.children.key = p6[2]
	trie.children = n6

	// test partial phrase, should be false
	member, value = trie.IsMember([]string{"break"})
	assert.False(t, member)

	member, value = trie.IsMember([]string{"break", "out"})
	assert.False(t, member)

	// full phrase
	member, value = trie.IsMember(p6)
	assert.True(t, member)
	assert.Equal(t, 1, value)

	// reset
	trie = NewPhraseTrie(nil)

	// weird phrase check
	n7 := testPhraseTrieNode()
	n7.key = p7[0]
	trie.children = n7

	member, value = trie.IsMember(p7)
	assert.True(t, member)
	assert.Equal(t, 1, value)
}

func TestAdd(t *testing.T) {
	// make empty root trie first
	trie := NewPhraseTrie(nil)

	// test phrases
	p1 := []string{"break"}
	p2 := []string{"shooting"}
	p3 := []string{"break", "out"}
	p4 := []string{"break", "up"}
	p5 := []string{"shooting", "up"}
	p6 := []string{"break", "out", "nicely"}
	p7 := []string{"r/g"}

	// add p1
	trie.Add(p1, 1)

	// check
	member, value := trie.IsMember(p1)
	assert.True(t, member)
	assert.Equal(t, 1, value)

	// add p2
	trie.Add(p2, 2)

	// check
	member, value = trie.IsMember(p2)
	assert.True(t, member)
	assert.Equal(t, 2, value)

	// add p3
	trie.Add(p3, 3)

	// check p3
	member, value = trie.IsMember(p3)
	assert.True(t, member)
	assert.Equal(t, 3, value)

	// check p1, should fail
	member, value = trie.IsMember(p1)
	assert.False(t, member)

	// add p4
	trie.Add(p4, 4)

	// check p3, should still be true
	member, value = trie.IsMember(p3)
	assert.True(t, member)
	assert.Equal(t, 3, value)

	// check p4
	member, value = trie.IsMember(p4)
	assert.True(t, member)
	assert.Equal(t, 4, value)

	// add p5
	trie.Add(p5, 5)

	// check p3, should still be true
	member, value = trie.IsMember(p3)
	assert.True(t, member)
	assert.Equal(t, 3, value)

	// check p4, should still be true
	member, value = trie.IsMember(p4)
	assert.True(t, member)
	assert.Equal(t, 4, value)

	// check p5
	member, value = trie.IsMember(p5)
	assert.True(t, member)
	assert.Equal(t, 5, value)

	// add p6, 3 level phrase
	trie.Add(p6, 6)

	// check p3, should be false now bc break out if a prefix to p6
	member, value = trie.IsMember(p3)
	assert.False(t, member)

	// check p4, should still be true
	member, value = trie.IsMember(p4)
	assert.True(t, member)
	assert.Equal(t, 4, value)

	// check p5, should still be true
	member, value = trie.IsMember(p5)
	assert.True(t, member)
	assert.Equal(t, 5, value)

	// check p6
	member, value = trie.IsMember(p6)
	assert.True(t, member)
	assert.Equal(t, 6, value)

	// add weird phrase p7
	trie.Add(p7, 7)

	member, value = trie.IsMember(p7)
	assert.True(t, member)
	assert.Equal(t, 7, value)
}

func TestRemove(t *testing.T) {
}

func BenchmarkAdd(b *testing.B) {
	// benchmark by new method with a map

	for i := 0; i < b.N; i++ {
		_ = testTrieFull()
	}
}

func BenchmarkIsMember(b *testing.B) {
	trie := testTrieFull()
	p := []string{"break", "out", "nicely"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = trie.IsMember(p)
	}
}
