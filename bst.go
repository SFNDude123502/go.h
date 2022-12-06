package main

import "fmt"

type Value interface {
	~int |
		~uint |
		~float32 |
		~float64 |
		~string |
		~rune |
		~byte
}

type bstNode[V Value] struct {
	value  V
	left   *bstNode[V]
	right  *bstNode[V]
	ignore bool
}

type bst[V Value] struct {
	root     *bstNode[V]
	len      int
	iterable []*bstNode[V]
}

func mkBst[V Value](value V) *bst[V] {
	var root *bstNode[V] = &bstNode[V]{value: value}
	return &bst[V]{root: root, iterable: []*bstNode[V]{root}, len: 1}
}

func (t *bst[V]) set(value V) {
	pos := t.root
	for {
		if value == pos.value {
			pos.value = value
			pos.ignore = false
			return
		} else if value < pos.value {
			if pos.left != nil {
				pos = pos.left
				continue
			}
			pos.left = &bstNode[V]{value: value}
			t.iterable = append(t.iterable, pos.left)
			break
		} else {
			if pos.right != nil {
				pos = pos.right
				continue
			}
			pos.right = &bstNode[V]{value: value}
			t.iterable = append(t.iterable, pos.right)
			break
		}
	}
	t.len++
}

func (t *bst[V]) exists(value V) (out bool) {
	pos := t.root
	for pos != nil {
		if value == pos.value {
			if !pos.ignore {
				out = true
			}
			return
		}
		if value < pos.value {
			pos = pos.left
		} else {
			pos = pos.right
		}
	}
	return false
}

func (t *bst[V]) delete(value V) {
	pos := t.root
	for pos != nil {
		if value == pos.value {
			pos.ignore = true
			return
		} else if value < pos.value {
			pos = pos.left
		} else {
			pos = pos.right
		}
	}
}

func (t *bst[V]) print() {
	for _, ival := range t.iterable {
		fmt.Print("(", ival.value, ")")
	}
	fmt.Print("\n")
}

func (t *bst[V]) realPrint() {
	var out [][]*bstNode[V]
	t.reRealPrint(t.root, 0, &out)
	for _, ival := range out {
		for _, jval := range ival {
			fmt.Print("(", jval.value, ")")
		}
		fmt.Print("\n")
	}
}

func (t *bst[V]) reRealPrint(pos *bstNode[V], lvl int, out *[][]*bstNode[V]) {
	if pos == nil {
		return
	}
	ouTmp := *out
	if len(ouTmp) == lvl {
		ouTmp = append(ouTmp, []*bstNode[V]{})
	}
	ouTmp[lvl] = append(ouTmp[lvl], pos)
	*out = ouTmp
	t.reRealPrint(pos.left, lvl+1, out)
	t.reRealPrint(pos.right, lvl+1, out)
}
