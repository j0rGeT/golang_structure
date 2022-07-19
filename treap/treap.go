package treap

import "math/rand"

const Maxn = 100005
var treap []*Node
var cnt int
var root int

type Node struct {
	Lc int
	Rc int
	Value int
	Pri int
	Num int
	Size int
}

func init()  {
	treap = make([]*Node, Maxn)
	cnt = 0
}

func NewNode(val int) int {
	cnt++
	treap[cnt].Value = val
	treap[cnt].Pri = rand.Int()
	treap[cnt].Num = 1
	treap[cnt].Size = 1
	treap[cnt].Rc = 0
	treap[cnt].Lc = 0
	return cnt
}


func Update(p int) {
	treap[p].Size = treap[treap[p].Lc].Size + treap[treap[p].Rc].Size + treap[p].Num
}

func Zig(p int) {
	q := treap[p].Lc
	treap[p].Lc = treap[q].Rc
	treap[q].Rc = p
	treap[q].Size = treap[p].Size
	Update(p)
	p=q
}

func Zag(p int) {
	q := treap[p].Rc
	treap[p].Rc = treap[q].Lc
	treap[q].Lc = p
	treap[q].Size = treap[p].Size
	Update(p)
	p=q
}

func Insert(p int, val int) {
	if p != 0 {
		p = NewNode(val)
		return
	}
	treap[p].Size++
	if val == treap[p].Value {
		treap[p].Num ++
		return
	}
	if val < treap[p].Value {
		Insert(treap[p].Lc, val)
		if treap[p].Pri < treap[treap[p].Lc].Pri {
			Zig(p)
		}
	} else {
		Insert(treap[p].Rc, val)
		if treap[p].Pri < treap[treap[p].Rc].Pri {
			Zag(p)
		}
	}
}

func Delete(p int, val int) {
	if p == 0 {
		return
	}
	treap[p].Size --
	if val == treap[p].Value {
		if treap[p].Num > 1 {
			treap[p].Num --
			return
		}
		if treap[p].Lc == 0 || treap[p].Rc == 0 {
			p = treap[p].Lc + treap[p].Rc
		}else if treap[treap[p].Lc].Pri > treap[treap[p].Rc].Pri {
			Zig(p)
			Delete(treap[p].Rc, val)
		}else {
			Zag(p)
			Delete(treap[p].Lc, val)
		}
		return
	}
	if val < treap[p].Value {
		Delete(treap[p].Lc, val)
	}else {
		Delete(treap[p].Rc, val)
	}
}

func GetPre(val int) int {
	p := root
	res := -1
	for {
		if p > 0 {
			if treap[p].Value < val {
				res = treap[p].Value
				p = treap[p].Rc
			} else {
				p = treap[p].Lc
			}
		} else {
			break
		}
	}
	return res
}

func GetNext(val int) int {
	p := root
	res := -1
	for {
		if p > 0 {
			if treap[p].Value > val {
				res = treap[p].Value
				p = treap[p].Lc
			}else {
				p = treap[p].Rc
			}
		} else {
			break
		}
	}
	return res
}

func BuildTreap() {
	var x int
	var n int
	for i:=1;i<=n;i++ {
		Insert(root, x)
	}
}