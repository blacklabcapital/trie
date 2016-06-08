package trie

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockPCList1() PCtxList {
	sentence := []string{"AAPL", "breaks", "long", "up", "high", "bound", "and", "short", "down", "hard", "today"}
	pcl := PCtxList{
		NewPhraseContext([]string{"breaks", "long", "up", "high", "bound"}, 1, 5, 1, sentence),
		NewPhraseContext([]string{"long", "up"}, 2, 3, 1, sentence),
		NewPhraseContext([]string{"short", "down", "hard"}, 7, 9, -1, sentence),
		NewPhraseContext([]string{"high", "bound", "and", "short", "down"}, 4, 8, -1, sentence),
	}

	return pcl
}

func mockPCList2() PCtxList {
	sentence := []string{"GOOG", "goes", "up", "and", "up", "then", "down", "turns", "out", "fast"}
	pcl := PCtxList{
		NewPhraseContext([]string{"out", "hard"}, 7, 9, -1, sentence),
		NewPhraseContext([]string{"down", "turns", "out", "fast"}, 6, 9, -1, sentence),
		NewPhraseContext([]string{"up", "and", "up"}, 2, 4, 1, sentence),
	}

	return pcl
}

func mockPCList3() PCtxList {
	sentence := []string{"up", "hard", "TSLA", "make", "money", "i", "will"}
	pcl := PCtxList{
		NewPhraseContext([]string{"up", "hard"}, 0, 1, 1, sentence),
		NewPhraseContext([]string{"make", "money"}, 3, 4, 1, sentence),
	}

	return pcl
}

func TestSuperOnly(t *testing.T) {
	pcl := mockPCList1()
	supers := pcl.SuperOnly()
	sentence1 := []string{"AAPL", "breaks", "long", "up", "high", "bound", "and", "short", "down", "hard", "today"}

	assert.Equal(t, 2, len(supers))
	assert.Equal(t, NewPhraseContext([]string{"breaks", "long", "up", "high", "bound"}, 1, 5, 1, sentence1), supers[0])
	assert.Equal(t, NewPhraseContext([]string{"short", "down", "hard"}, 7, 9, -1, sentence1), supers[1])

	pcl = mockPCList2()
	supers = pcl.SuperOnly()

	sentence2 := []string{"GOOG", "goes", "up", "and", "up", "then", "down", "turns", "out", "fast"}
	assert.Equal(t, 2, len(supers))
	assert.Equal(t, NewPhraseContext([]string{"up", "and", "up"}, 2, 4, 1, sentence2), supers[0])
	assert.Equal(t, NewPhraseContext([]string{"down", "turns", "out", "fast"}, 6, 9, -1, sentence2), supers[1])

	pcl = mockPCList3()
	supers = pcl.SuperOnly()

	sentence3 := []string{"up", "hard", "TSLA", "make", "money", "i", "will"}
	assert.Equal(t, 2, len(supers))
	assert.Equal(t, NewPhraseContext([]string{"up", "hard"}, 0, 1, 1, sentence3), supers[0])
	assert.Equal(t, NewPhraseContext([]string{"make", "money"}, 3, 4, 1, sentence3), supers[1])

	// trie found phrases
	// full phrase trie with all test phrases
	trie := mockTrieFull()
	// nothing
	s := "$AAPL isn't doing anything today"
	sSplit := strings.Split(s, " ")
	supers = trie.FindAllMembers(sSplit).SuperOnly()
	assert.Equal(t, 0, len(supers))

	s = "$AAPL will break out nicely"
	sSplit = strings.Split(s, " ")
	supers = trie.FindAllMembers(sSplit).SuperOnly()
	assert.Equal(t, 1, len(supers))
	assert.Equal(t, []int{2, 4}, supers[0].Indices)
	assert.Equal(t, "break out nicely", supers[0].PhraseStr())
	assert.Equal(t, 6, supers[0].Value)

	// mult phrases
	s = "its shooting up it might even break up i bet $100 $AAPL will break out nicely"
	sSplit = strings.Split(s, " ")
	supers = trie.FindAllMembers(sSplit).SuperOnly()
	assert.Equal(t, 3, len(supers))
	assert.Equal(t, []int{1, 2}, supers[0].Indices)
	assert.Equal(t, "shooting up", supers[0].PhraseStr())
	assert.Equal(t, 5, supers[0].Value)
	assert.Equal(t, []int{6, 7}, supers[1].Indices)
	assert.Equal(t, "break up", supers[1].PhraseStr())
	assert.Equal(t, 4, supers[1].Value)
	assert.Equal(t, []int{13, 15}, supers[2].Indices)
	assert.Equal(t, "break out nicely", supers[2].PhraseStr())
	assert.Equal(t, 6, supers[2].Value)

	// super phrases with sub phrase
	s = "its breaking double bottom $100 $AAPL will break out"
	sSplit = strings.Split(s, " ")
	supers = trie.FindAllMembers(sSplit).SuperOnly()
	assert.Equal(t, 1, len(supers))
	assert.Equal(t, []int{1, 3}, supers[0].Indices)
	assert.Equal(t, "breaking double bottom", supers[0].PhraseStr())
	assert.Equal(t, 8, supers[0].Value)
}

func BenchmarkSuperOnly(b *testing.B) {
	pcl := mockPCList1()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = pcl.SuperOnly()
	}
}
