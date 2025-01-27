package node

import (
	"cmp"
)

// Curr -> Left -> Right
func PreOrder[T cmp.Ordered](head *node[T]) []*node[T] {
	if head == nil {
		return []*node[T]{}
	}

	var curr *node[T]
	ret := make([]*node[T], 0, head.TreeDepth()*2+1)
	stk := NewStack[T](head.TreeDepth())
	stk.Push(head)

	for stk.Len() > 0 {
		curr = stk.Pop()

		ret = append(ret, curr)

		if curr.right != nil {
			stk.Push(curr.right)
		}
		if curr.left != nil {
			stk.Push(curr.left)
		}
	}

	return ret
}

// Left -> Curr -> Right
func InOrder[T cmp.Ordered](head *node[T]) []*node[T] {
	if head == nil {
		return []*node[T]{}
	}

	var curr *node[T]
	ret := make([]*node[T], 0, head.TreeDepth()*2+1)
	stk := NewStack[T](head.TreeDepth())
	stk.Push(head)

	for stk.Len() > 0 {
		curr = stk.Peep()

		// Go down left as far as possible
		for curr.left != nil {
			stk.Push(curr.left)
			curr = stk.Peep()
		}

		// Go up until there's a right branch
		for curr != nil && curr.right == nil {
			ret = append(ret, stk.Pop())
			curr = stk.Peep()
		}

		// Only if we got back to the head
		if curr == nil {
			break
		}

		// Get current then move right
		ret = append(ret, curr)
		stk.Pop()
		stk.Push(curr.right)
	}

	return ret
}

// type seenCount[T cmp.Ordered] struct {
// 	n     *node[T]
// 	count int
// }

// // Left -> Curr -> Right
// func InOrder[T cmp.Ordered](head *node[T]) []*node[T] {
// 	if head == nil {
// 		return []*node[T]{}
// 	}

// 	var curr *seenCount[T]
// 	ret := make([]*node[T], 0, head.TreeDepth()*2+1)
// 	stk := []*seenCount[T]{{n: head}}
// 	l := 0

// 	for l > -1 {
// 		curr = stk[l]
// 		fmt.Println(curr.n, stk)

// 		switch curr.count {
// 		case 0:
// 			if curr.n.left != nil {
// 				stk = append(stk, &seenCount[T]{n: curr.n.left})
// 				l += 1
// 			}
// 			curr.count += 1
// 		case 1:
// 			ret = append(ret, curr.n)

// 			stk = stk[:l]
// 			l -= 1

// 			if curr.n.right != nil {
// 				stk = append(stk, &seenCount[T]{n: curr.n.right})
// 				l += 1
// 			}
// 		}

// 	}

// 	return ret
// }
