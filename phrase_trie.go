package trie

import (
	"strings"
)

/* ARRAY BASED VECTOR IMPLEMENTAITON */

// A Trie is a kind of search tree, an ordered data structure that stores a dynamic set or associative array
// using a string key, where the position in the trie denotes the value of the key

// A PhraseTrie is an implementation of a Trie tree but for single/multi word phrases
// where a part of a phrase is a full word or expression

// A PhraseTrieNode is a Trie element that stores its key/value pair
// and a list of children nodes
type PhraseTrieNode struct {
	key      string
	value    int
	children []*PhraseTrieNode
}

// NewPhraseTrie creates a new Trie tree by initializing and returning a root Node
// as the base of the Trie.
// If phrases key/value map is supplied, adds all the given phrases to the Trie
// to create the full phrase tree
func NewPhraseTrie(phrases map[string]int) *PhraseTrieNode {
	root := &PhraseTrieNode{children: []*PhraseTrieNode{}} // init children to 0 len slice

	for k, v := range phrases {
		root.Add(strings.Split(k, " "), v)
	}

	return root
}

// Add recursively adds a phrase key/value to this Trie
// Note: if adding a multi word phrase with a prefix that
// already exists in the Trie, that prefix will no longer
// be a valid phrase member of the Trie. Only full phrases
// that end in a leaf are valid members
func (n *PhraseTrieNode) Add(phrase []string, value int) {
	if n.IsLeaf() {
		n.children = []*PhraseTrieNode{&PhraseTrieNode{key: phrase[0]}}
		if len(phrase) != 1 {
			n.children[0].Add(phrase[1:], value)
		} else { // leaf, set value
			n.children[0].value = value
		}
	} else { // has children
		for i, child := range n.children {
			if phrase[0] == child.key {
				if len(phrase) != 1 {
					n.children[i].Add(phrase[1:], value)
				} // else already exists, return from function

				return
			}
		}

		// add new node
		n.children = append(n.children, &PhraseTrieNode{key: phrase[0]})
		if len(phrase) != 1 {
			n.children[len(n.children)-1].Add(phrase[1:], value)
		} else { // leaf, set value
			n.children[len(n.children)-1].value = value
		}
	}
}

// Remove recursively removes a phrase from this Trie
// Preserves other phrases if other nodes use the same prefixes
// TODO: implement
func (n *PhraseTrieNode) Remove(phrase []string) {
}

// IsMember checks if the given phrase is a member of this Phrase Trie tree
// and returns the phrase value if true
func (n *PhraseTrieNode) IsMember(phrase []string) (bool, int) {
	for _, child := range n.children {
		if phrase[0] == child.key {
			if len(phrase) != 1 {
				if child.IsLeaf() { // full phrase not member
					return false, 0
				}

				return child.IsMember(phrase[1:])
			} else if child.IsLeaf() { // match
				return true, child.value
			}

			// not a leaf, no match
			return false, 0
		}
	}

	return false, 0
}

// FindMember recursively traverses this Trie to find if the given
// sequence begins with a member phrase
//
// Returns the a bool valid if the parts found were a valid
// full member phrase, the found phrase, and its value
//
// A valid member phrase ends its search on a leaf node, i.e. a full phrase
//
// If there are multiple member phrases in the sequence FindMember only
// finds and returns the FIRST found phrase
func (n *PhraseTrieNode) FindMember(sequence []string) (bool, []string, int) {
	var (
		phrase []string
		valid  bool
		value  int
	)

	phrase = make([]string, 0)

	for _, child := range n.children {
		if child.key == sequence[0] {
			if child.IsLeaf() { // found phrase
				valid = true
				value = child.value
				phrase = append(phrase, child.key)
				break
			} else if len(sequence) != 1 { // recur down trie
				// first add this child's key to the phrase
				phrase = append(phrase, child.key)

				// recur down matched child
				childValid, childPhrase, childValue := child.FindMember(sequence[1:])

				valid = childValid
				value = childValue

				for _, p := range childPhrase {
					phrase = append(phrase, p)
				}

				break
			} else { // not valid phrase
				break
			}
		}
	}

	return valid, phrase, value
}

// FindAllMembers iterates in a linear sequential fashion through a sentence
// array and finds all potential phrases in the sentence that are members
// of this Trie.
// Returns a map containing the phrase string and its value
// An empty (len == 0) map consitutes no valid member phrases found in the given sentence
func (n *PhraseTrieNode) FindAllMembers(sentence []string) PCtxList {
	foundMembers := make(PCtxList, 0)

	for i := 0; i < len(sentence); i++ {
		if n.IsLeaf() { // no children to match
			return nil
		}

		valid, p, v := n.FindMember(sentence[i:])

		if valid { // valid phrase was found
			foundMembers = append(foundMembers, NewPhraseContext(p, sentence, []int{i, i + len(p) - 1}, v))
		}
	}

	return foundMembers
}

// IsLeaf returns true if this node is a leaf
// A node is a leaf when it has no children
func (n *PhraseTrieNode) IsLeaf() bool {
	return n.children == nil || len(n.children) == 0
}
