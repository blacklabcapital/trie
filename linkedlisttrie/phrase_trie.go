package linkedlisttrie

import (
	"strings"
)

/* LINKED LIST IMPLEMENTAITON */

// A Trie is a kind of search tree, an ordered data structure that stores a dynamic set or associative array
// using a string key, where the position in the trie denotes the value of the key

// A PhraseTrie is an implementation of a Trie tree but for single/multi word phrases
// where a part of a phrase is a full word or expression

// A PhraseTrieNode is a Trie element that stores a key/value pair, a pointer to next Node
// and child nodes if any
type PhraseTrieNode struct {
	key      string
	value    int
	next     *PhraseTrieNode
	children *PhraseTrieNode
}

// NewPhraseTrie creates a new Trie tree by returning a pointer to a
// root PhraseTrieNode with the empty string as the key
// If phrases key/value map is supplied, adds all the given phrases to the Trie
// to create the full phrase tree
func NewPhraseTrie(phrases map[string]int) *PhraseTrieNode {
	root := &PhraseTrieNode{}

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
	if n.key != "" {
		if len(phrase) != 1 {
			if n.key == phrase[0] {
				if n.IsLeaf() { // no children yet
					n.children = &PhraseTrieNode{key: phrase[1]}
					n.children.Add(phrase[1:], value)
				} else {
					n.children.Add(phrase[1:], value)
				}
			} else if n.HasNext() {
				n.next.Add(phrase, value)
			} else { // add next
				n.next = &PhraseTrieNode{key: phrase[0]}
				n.next.Add(phrase, value)
			}
		} else if n.key == phrase[0] { // this is leaf
			n.value = value
			return
		} else if n.HasNext() {
			n.next.Add(phrase, value)
		} else { // add next leaf
			n.next = &PhraseTrieNode{key: phrase[0], value: value}
			return
		}
	} else if n.IsLeaf() { // root no children
		n.children = &PhraseTrieNode{key: phrase[0]}
		n.children.Add(phrase, value)
	} else {
		n.children.Add(phrase, value)
	}
}

// Remove recursively removes a phrase from this Trie
// Preserves other phrases if other nodes use the same prefixes
func (n *PhraseTrieNode) Remove(phrase []string) {
}

// IsMember checks if the given phrase is a member of this Phrase Trie tree
// and returns the phrase value if true
func (n *PhraseTrieNode) IsMember(phrase []string) (bool, int) {
	if n.key != "" {
		if len(phrase) != 1 {
			if n.key == phrase[0] {
				if n.IsLeaf() { // no children yet
					return false, 0
				} else {
					return n.children.IsMember(phrase[1:])
				}
			} else if n.HasNext() {
				return n.next.IsMember(phrase)
			} else { // add next
				return false, 0
			}
		} else if n.key == phrase[0] { // match
			if n.IsLeaf() { // must be leaf, cant match partial
				return true, n.value
			} else {
				return false, 0
			}
		} else if n.HasNext() {
			return n.next.IsMember(phrase)
		} else {
			return false, 0
		}
	} else if n.IsLeaf() { // root no children
		return false, 0
	} else {
		return n.children.IsMember(phrase)
	}
}

// IsLeaf returns true if this node is a leaf
// A node is a leaf when it has no children
func (n *PhraseTrieNode) IsLeaf() bool {
	return n.children == nil
}

// HasNext returns true if this node has an adjacent next node
func (n *PhraseTrieNode) HasNext() bool {
	return n.next != nil
}
