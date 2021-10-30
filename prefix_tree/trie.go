package trie

const AlphabetSize = 26

type Trie struct {
	children [AlphabetSize]*Trie
	isLeaf   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	var obj = Trie{}
	obj.isLeaf = false
	return obj
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	var cur = this
	for _, ch := range word {
		var idx = ch - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = &Trie{}
		}
		cur = cur.children[idx]
	}
	cur.isLeaf = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	var cur = this
	for _, ch := range word {
		var idx = ch - 'a'
		if cur.children[idx] == nil {
			return false
		}
		cur = cur.children[idx]
	}
	return cur.isLeaf
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	var cur = this
	for _, ch := range prefix {
		var idx = ch - 'a'
		if cur.children[idx] == nil {
			return false
		}
		cur = cur.children[idx]
	}
	return true
}
