package trie

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockPhraseTrieNode() *PhraseTrieNode {
	return &PhraseTrieNode{key: "test", value: 1, children: []*PhraseTrieNode{}}
}

func mockTrieFull() *PhraseTrieNode {
	m := map[string]int{
		"break":            1,
		"shooting":         2,
		"break out":        3,
		"break up":         4,
		"shooting up":      5,
		"break out nicely": 6,
		"r/g":              7,
		"breaking double bottom": 8,
		"double bottom":          9,
	}

	return NewPhraseTrie(m)
}

func TestIsLeaf(t *testing.T) {
	n := mockPhraseTrieNode()
	assert.True(t, n.IsLeaf())

	n.children = nil
	assert.True(t, n.IsLeaf())

	n.children = []*PhraseTrieNode{mockPhraseTrieNode()}
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
	n1 := mockPhraseTrieNode()
	n1.key = p1[0]
	trie.children = append(trie.children, n1)

	member, value = trie.IsMember(p1)
	assert.True(t, member)
	assert.Equal(t, 1, value)

	// add p2
	n2 := mockPhraseTrieNode()
	n2.key = p2[0]
	n2.value = 2
	trie.children = append(trie.children, n2)

	member, value = trie.IsMember(p2)
	assert.True(t, member)
	assert.Equal(t, 2, value)

	// reset
	trie = NewPhraseTrie(nil)

	// add 2 word phrase p3
	n3 := mockPhraseTrieNode()
	n3.key = p3[0]
	n3.children = append(n3.children, mockPhraseTrieNode())
	n3.children[0].key = p3[1]
	n3.children[0].value = 3
	trie.children = append(trie.children, n3)

	member, value = trie.IsMember(p3)
	assert.True(t, member)
	assert.Equal(t, 3, value)

	// add same prefix new 2nd word phrase p4 to existing trie
	n4 := mockPhraseTrieNode()
	n4.key = p4[1]
	n4.value = 4
	trie.children[0].children = append(trie.children[0].children, n4)

	// check p3 just to make sure
	member, value = trie.IsMember(p3)
	assert.True(t, member)
	assert.Equal(t, 3, value)

	// check p4
	member, value = trie.IsMember(p4)
	assert.True(t, member)
	assert.Equal(t, 4, value)

	// add second 2nd layer phrase to existing trie
	n5 := mockPhraseTrieNode()
	n5.key = p5[0]
	n5.children = append(n5.children, mockPhraseTrieNode())
	n5.children[0].key = p5[1]
	n5.children[0].value = 5
	trie.children = append(trie.children, n5)

	member, value = trie.IsMember(p5)
	assert.True(t, member)
	assert.Equal(t, 5, value)

	// reset
	trie = NewPhraseTrie(nil)

	// add 3 word phrase
	n6 := mockPhraseTrieNode()
	n6.key = p6[0]
	n6.children = append(n6.children, mockPhraseTrieNode())
	n6.children[0].key = p6[1]
	n6.children[0].children = append(n6.children[0].children, mockPhraseTrieNode())
	n6.children[0].children[0].key = p6[2]
	n6.children[0].children[0].value = 6
	trie.children = append(trie.children, n6)

	// test partial phrase, should be false
	member, value = trie.IsMember([]string{"break"})
	assert.False(t, member)

	member, value = trie.IsMember([]string{"break", "out"})
	assert.False(t, member)

	// full phrase
	member, value = trie.IsMember(p6)
	assert.True(t, member)
	assert.Equal(t, 6, value)

	// reset
	trie = NewPhraseTrie(nil)

	// weird phrase check
	n7 := mockPhraseTrieNode()
	n7.key = p7[0]
	trie.children = append(trie.children, n7)

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

func TestFindMember(t *testing.T) {
	// make empty root trie first
	trie := NewPhraseTrie(nil)

	// test phrases
	p5 := []string{"shooting", "up"}
	p6 := []string{"break", "out", "nicely"}
	p7 := []string{"r/g"}

	// add
	trie.Add(p5, 5)
	trie.Add(p6, 6)
	trie.Add(p7, 7)

	// junk shit
	bad := "1@#%*!#)@% $AAPL"
	valid, phrase, value := trie.FindMember(strings.Split(bad, " "))
	assert.False(t, valid)
	assert.Equal(t, 0, len(phrase))
	assert.Equal(t, 0, value)

	// false
	s1 := "$AAPL is shooting up"
	valid, phrase, value = trie.FindMember(strings.Split(s1, " "))
	assert.False(t, valid)
	assert.Equal(t, 0, len(phrase))
	assert.Equal(t, 0, value)

	// true
	s1 = "shooting up $AAPL is!"
	valid, phrase, value = trie.FindMember(strings.Split(s1, " "))
	assert.True(t, valid)
	assert.Equal(t, 2, len(phrase))
	assert.Equal(t, 5, value)

	// false, partial prefix
	s2 := "break"
	valid, phrase, value = trie.FindMember(strings.Split(s2, " "))
	assert.False(t, valid)
	assert.Equal(t, 0, len(phrase))
	assert.Equal(t, 0, value)

	// true, full phrase
	s2 = "break out nicely today $AAPL will??"
	valid, phrase, value = trie.FindMember(strings.Split(s2, " "))
	assert.True(t, valid)
	assert.Equal(t, 3, len(phrase))
	assert.Equal(t, 6, value)

	// true
	s3 := "r/g"
	valid, phrase, value = trie.FindMember(strings.Split(s3, " "))
	assert.True(t, valid)
	assert.Equal(t, 1, len(phrase))
	assert.Equal(t, 7, value)
}

func TestFindAllMember(t *testing.T) {
	// full phrase trie with all test phrases
	trie := mockTrieFull()

	// nothing
	s := "$AAPL isn't doing anything today"
	sSplit := strings.Split(s, " ")
	phrases := trie.FindAllMembers(sSplit)
	assert.Equal(t, 0, len(phrases))

	// prefix only
	s = "$AAPL isn't gonna break today"
	sSplit = strings.Split(s, " ")
	phrases = trie.FindAllMembers(sSplit)
	assert.Equal(t, 0, len(phrases))

	// one beginning phrase
	s = "shooting up $AAPL is today!"
	sSplit = strings.Split(s, " ")
	phrases = trie.FindAllMembers(sSplit)
	assert.Equal(t, 1, len(phrases))
	assert.Equal(t, []int{0, 1}, phrases[0].Indices)
	assert.Equal(t, "shooting up", phrases[0].PhraseStr())
	assert.Equal(t, 5, phrases[0].Value)

	// one middle phrase
	s = "$AAPL might break up today!"
	sSplit = strings.Split(s, " ")
	phrases = trie.FindAllMembers(sSplit)
	assert.Equal(t, 1, len(phrases))
	assert.Equal(t, []int{2, 3}, phrases[0].Indices)
	assert.Equal(t, "break up", phrases[0].PhraseStr())
	assert.Equal(t, 4, phrases[0].Value)

	// one end phrase
	s = "$AAPL will break out nicely"
	sSplit = strings.Split(s, " ")
	phrases = trie.FindAllMembers(sSplit)
	assert.Equal(t, 1, len(phrases))
	assert.Equal(t, []int{2, 4}, phrases[0].Indices)
	assert.Equal(t, "break out nicely", phrases[0].PhraseStr())
	assert.Equal(t, 6, phrases[0].Value)

	// mult phrases
	s = "its shooting up it might even break up i bet $100 $AAPL will break out nicely"
	sSplit = strings.Split(s, " ")
	phrases = trie.FindAllMembers(sSplit)
	assert.Equal(t, 3, len(phrases))
	assert.Equal(t, []int{1, 2}, phrases[0].Indices)
	assert.Equal(t, "shooting up", phrases[0].PhraseStr())
	assert.Equal(t, 5, phrases[0].Value)
	assert.Equal(t, []int{6, 7}, phrases[1].Indices)
	assert.Equal(t, "break up", phrases[1].PhraseStr())
	assert.Equal(t, 4, phrases[1].Value)
	assert.Equal(t, []int{13, 15}, phrases[2].Indices)
	assert.Equal(t, "break out nicely", phrases[2].PhraseStr())
	assert.Equal(t, 6, phrases[2].Value)

	// super phrases with sub phrase
	s = "its breaking double bottom $100 $AAPL will break out"
	sSplit = strings.Split(s, " ")
	phrases = trie.FindAllMembers(sSplit)
	assert.Equal(t, 2, len(phrases))
	assert.Equal(t, []int{1, 3}, phrases[0].Indices)
	assert.Equal(t, "breaking double bottom", phrases[0].PhraseStr())
	assert.Equal(t, 8, phrases[0].Value)
	assert.Equal(t, []int{2, 3}, phrases[1].Indices)
	assert.Equal(t, "double bottom", phrases[1].PhraseStr())
	assert.Equal(t, 9, phrases[1].Value)
}

func BenchmarkAdd(b *testing.B) {
	// benchmark by new method with a map

	for i := 0; i < b.N; i++ {
		_ = mockTrieFull()
	}
}

func BenchmarkIsMember(b *testing.B) {
	trie := mockTrieFull()
	p := []string{"break", "out", "nicely"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = trie.IsMember(p)
	}
}

func BenchmarkFindAllMembers(b *testing.B) {
	trie := mockTrieFull()
	s := "its shooting up it might even break up i bet $100 $AAPL will break out nicely"
	sSplit := strings.Split(s, " ")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = trie.FindAllMembers(sSplit)
	}
}
