package solve

type Trie struct {
	//values map[string]struct{}
	next [26]*Trie
	isLeaf bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
	}
}

func NewTrie() *Trie {
	return &Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	cur := this
	var c uint8
	var idx int
	for i := range word {
		c = word[i]
		idx = char2Int(c)
		if cur.next[idx] == nil {
			cur.next[idx] = NewTrie()
		}

		cur = cur.next[idx]
	}
	cur.isLeaf = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {

	cur := this
	var c uint8
	var idx int
	for i := range word {
		c = word[i]
		idx = char2Int(c)
		if cur.next[idx] == nil {
			return false
		}
		cur = cur.next[idx]
	}

	return cur.isLeaf
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	var c uint8
	var idx int
	for i := range prefix {
		c = prefix[i]
		idx = char2Int(c)
		if cur.next[idx] == nil {
			return false
		}
		cur = cur.next[idx]
	}

	return true
}


func char2Int(c uint8) int {
	return int(c - 'a')
}