package solve

type WordDictionary struct {
	child  [26] *WordDictionary
	isLeaf bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

func NewWordDictionary() *WordDictionary {
	return &WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	var c uint8
	var index int
	cur := this
	for i := range word {
		c = word[i]
		index = char2Int(c)
		if cur.child[index] == nil {
			cur.child[index] = NewWordDictionary()
		}

		cur = cur.child[index]
	}

	cur.isLeaf = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	cur := this
	var index int
	var c uint8
	for i := range word {
		c = word[i]
		if c == '.' {
			for _, cd := range cur.child {
				if cd != nil {
					if cd.Search(word[i+1:]) {
						return true
					}
				}
			}
			return false
		} else {
			index = char2Int(c)
			if cur.child[index] == nil {
				return false
			}
			cur = cur.child[index]
		}
	}

	return cur.isLeaf
}

func char2Int(c uint8) int {
	return int(c - 'a')
}
