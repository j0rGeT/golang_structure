package main

import "fmt"

const Size = 26

type TrieNode struct {
	Data uint8
	Flag bool
	Next []*TrieNode
}

type TrieTree struct {
	Len int
	Next []*TrieNode
}

func NewTrieTree() *TrieTree {
	return &TrieTree{
		Len: 0,
		Next: nil,
	}
}

func (t *TrieTree) Insert(str string) {
	if t.Next == nil {
		t.Next = make([]*TrieNode, Size, Size)
	}

	cur := t.Next
	index := uint8(0)
	for i := 0; i < len(str); i ++ {
		index = str[i] - 'a'
		if cur[index] == nil {
			newNode := &TrieNode{
				Data: str[i],
				Flag: false,
				Next: make([]*TrieNode, Size, Size),
			}
			cur[index] = newNode
		}else {
			cur[index].Data = str[i]
		}

		if i == len(str) - 1 {
			cur[index].Flag = true
		}
		cur = cur[index].Next
	}
}

func (t *TrieTree) IsExist(str string) bool {
	if t.Next == nil {
		return false
	}

	cur := t.Next
	index := uint8(0)
	for i:=0;i<len(str);i++ {
		index = str[i] - 'a'
		if cur[index] == nil {
			return false
		}
		if cur[index].Data != str[i] {
			return false
		}
		if i == len(str) - 1 && cur[index].Flag == true {
			return true
		}
		cur = cur[index].Next
		if cur == nil {
			return false
		}
	}
	return false
}

func main() {
	trieTree := NewTrieTree()

	trieTree.Insert("jeee")
	trieTree.Insert("je")
	fmt.Println(trieTree.IsExist("je"))
	fmt.Println(trieTree.IsExist("jee"))
}