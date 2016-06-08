# Trie

[![CircleCI](https://circleci.com/gh/blacklabcapital/trie.svg?style=svg)](https://circleci.com/gh/blacklabcapital/trie)

## Description

The `trie` package provides simple, performant and unique open-source implementations of the trie data structure.

 A Trie is a kind of search tree, an ordered data structure that stores a dynamic set or associative array
using a string key, where the position in the trie denotes the value of the key.


## Usage

`import "github.com/blacklabcapital/trie"`

You can use `trie` data structures for various types of purposes/environments, such as *natural language processing, sentiment analysis, lexicographic analysis*, or other situations where an efficient, highly performant lookup store is needed.

The package can be and has been used in real time low-latency environments, and performs effectively even in microsecond latency environments.

The library has extensive unit tests which also double as examples for common usage. Please see the godoc for package documentation.



Currently, the main trie implemenation of this package is the `PhraseTrie` type.

#### PhraseTrie

A `PhraseTrie` is an implementation of a trie data structure but for single/multi word phrase keys, where a part of a phrase is a full word or expression, compared to the more commonly implemented "word" trie, where a node value is a single character and a key is a full word.

There are currently two supported implementations of the PhraseTrie:

- an array based vector implementation
- a linked list implementation

The vector based trie has more features and methods available and is the main phrasetrie data structure.



## Contributing

`master` holds the latest current stable version of trie. Commits with a minor version are guaranteed to have no breaking API changes, only feature additions and bug fixes.

`dev` holds the latest commits and is where active developmemnt takes place. If you submit a pull request it should be against the `dev` branch.

`<major.minor>` are version branches. Tested changes from `dev` are staged for a release by merging into the appropriate version branch.
