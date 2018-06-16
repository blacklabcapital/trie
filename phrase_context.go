package trie

import (
	"sort"
	"strings"
)

// PhraseContext contains the found phrase, the sentence in which the phrase was found,
// the word indices of found phrase in the sentence, and the sentiment value of the phrase
type PhraseContext struct {
	Phrase   []string
	Indices  []int
	Value    int
	Sentence []string
}

// NewPhraseContext constructs and initializes a new PhraseContext
// NOTE: Always use this constructor when creating a new PhraseContext
func NewPhraseContext(phrase []string, sentence []string, indices []int, value int) *PhraseContext {
	pc := PhraseContext{
		Phrase:   phrase,
		Indices:  indices,
		Value:    value,
		Sentence: sentence,
	}

	return &pc
}

// SentenceStr returns this PhraseContext's sentence as a string
func (p *PhraseContext) SentenceStr() string {
	return strings.Join(p.Sentence, " ")
}

// PhraseStr returns this PhraseContext's phrase as a string
func (p *PhraseContext) PhraseStr() string {
	return strings.Join(p.Phrase, " ")
}

// PCtxList is a list of PhraseContext pointers
// Implements sort.Interface for []*PhraseContext based on
// lower bound indices first then upper bound
type PCtxList []*PhraseContext

func (pcl PCtxList) Len() int {
	return len(pcl)
}

func (pcl PCtxList) Swap(i, j int) {
	pcl[i], pcl[j] = pcl[j], pcl[i]
}

func (pcl PCtxList) Less(i, j int) bool {
	iIndices := pcl[i].Indices
	jIndices := pcl[j].Indices

	// bounds check
	if len(iIndices) > 0 && len(jIndices) > 0 {
		if iIndices[0] == jIndices[0] { // same first idx, use second
			return iIndices[1] < jIndices[1]
		}

		return iIndices[0] < jIndices[0]
	}

	return true
}

// SuperOnly returns a PCtxList with filtered out post super subphrases
// I.e. removes subphrases that have forward overlapping indices with a superphrase
// example:
//		super phrase: breaking double floor
//		sub phrase: double floor
// 	SuperOnly will remove 'double floor' from, as it is contained
// in the word superset of another found phrase
// Note: This will NOT remove pre super subphrases. that is complicated and not desirable
func (pcl PCtxList) SuperOnly() PCtxList {
	if len(pcl) == 0 {
		return PCtxList{}
	}

	// sort first by indices
	sort.Sort(pcl)

	// create sorted list index lookup array
	lookups := make([]int, len(pcl[0].Sentence))

	// fill with dummy value -1s
	for i := 0; i < len(lookups); i++ {
		lookups[i] = -1
	}

	// filter out post super subphrases using color array algorithm
Outer:
	for i, c := range pcl {
		for j := c.Indices[0]; j < c.Indices[1]+1; j++ {
			if lookups[j] == -1 {
				lookups[j] = i
			} else {
				newIndices := c.Indices
				curIndices := pcl[lookups[j]].Indices

				if newIndices[0] < curIndices[0] {
					lookups[i] = newIndices[0]
				} else if newIndices[0] == curIndices[0] && newIndices[1] > curIndices[1] { // same lower bound, high upper, replace
					lookups[i] = newIndices[0]
				} else { // c is a subphrase, skip
					continue Outer
				}
			}
		}
	}

	// return only supers
	supers := make(PCtxList, 0)
	last := -1
	for i := 0; i < len(lookups); i++ {
		if lookups[i] != -1 && lookups[i] != last {
			supers = append(supers, pcl[lookups[i]])
			last = lookups[i]
		}
	}

	return supers
}
