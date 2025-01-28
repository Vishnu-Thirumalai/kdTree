package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func printNodeChildren(curr *node[string]) {
// 	v := curr.GetValue()
// 	var l, r []string
// 	if curr.GetLeft() != nil {
// 		l = curr.GetLeft().GetValue()
// 	}
// 	if curr.GetRight() != nil {
// 		r = curr.GetRight().GetValue()
// 	}
// 	fmt.Println(v, l, r)
// }

// returns tree with format:
//
//	  a
//	b   c
//
// d e f g
func getTestTree() *node[string] {
	return NewTreeFromList([][]string{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}, {"f"}, {"g"}})
}

func Test_Preorder(t *testing.T) {
	tree := getTestTree()

	expected := []string{"a", "b", "d", "e", "c", "f", "g"}
	preOrder := PreOrder(tree)

	assert.Len(t, preOrder, len(expected))

	for idx, val := range expected {
		assert.Equal(t, val, preOrder[idx].GetValue()[0])
	}
}

func Test_Inorder(t *testing.T) {
	tree := getTestTree()

	expected := []string{"d", "b", "e", "a", "f", "c", "g"}
	inOrder := InOrder(tree)

	assert.Len(t, inOrder, len(expected))

	for idx, val := range expected {
		assert.Equal(t, val, inOrder[idx].GetValue()[0])
	}
}

func Test_Postorder(t *testing.T) {
	tree := getTestTree()

	expected := []string{"d", "e", "b", "f", "g", "c", "a"}
	postOrder := PostOrder(tree)

	assert.Len(t, postOrder, len(expected))

	for idx, val := range expected {
		assert.Equal(t, val, postOrder[idx].GetValue()[0])
	}
}
