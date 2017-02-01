package main

import "fmt"
import "os"

func (t *RBtree) Initialize() int {
	t.root = new(RBnode)
	t.root.InitializeN(new(RBnode))
	return 0
}

func (n *RBnode) InitializeN(zTree *RBnode) int {
	n.value = -1
	n.color = "black"
	n.parent = zTree
	return 0
}

func (t *RBtree) FindNode(z int) *RBnode {
	return t.root.FindNodeR(z)
}

func (n *RBnode) FindNodeR(z int) *RBnode {
	if z == n.value {
		return n
	} else if z < n.value {
		return n.left.FindNodeR(z)
	} else if z > n.value {
		return n.right.FindNodeR(z)
	} else {
		fmt.Println("There was an error in FindNode")
		return n
	}
}

func (t *RBtree) FindMin(z int) *RBnode {
	return t.FindNode(z).FindMinR()
}

func (n *RBnode) FindMinR() *RBnode {
	if n.left.value == -1 {
		return n
	} else {
		return n.left.FindMinR()
	}
}

func (n *RBnode) IsRoot() bool {
	if n.parent.value == 0 {
		return true
	} else {
		return false
	}
}

func (t *RBtree) LeftRotate(z int) int {
	// comments are identical to comments in psuedocode on pg 313 (13.2) of
	//   Introduction to Algorithms -- Third Edition, by Cormen et. al.
	//fmt.Printf("Let me L rotate this guy: %v\n", z)

	xTree := t.FindNode(z)
	if xTree.right.value == -1 {
		fmt.Println("LeftRotate was called on a node with a nil right child.")
		os.Exit(1)
	}

	// set y
	yTree := xTree.right

	// turn y's left subtree into x's right subtree
	xTree.right = yTree.left
	xTree.right.parent = xTree

	// link x's parent to y
	yTree.parent = xTree.parent
	if xTree.parent.value == 0 {
	} else if xTree.parent.left.value == xTree.value {
		xTree.parent.left = yTree
	} else {
		xTree.parent.right = yTree
	}

	// put x on y's left
	yTree.left = xTree
	yTree.left.parent = yTree

	if !t.root.IsRoot() {
		t.root = yTree
	}

	return 0
}

func (t *RBtree) RightRotate(z int) int {
	// comments are from the comments in the psuedocode on pg 313 (13.2) of
	//   Introduction to Algorithms -- Third Edition, by Cormen et. al.
	//fmt.Printf("Let me R rotate this guy: %v\n", z)

	xTree := t.FindNode(z)
	if xTree.left.value == -1 {
		fmt.Println("RightRotate was called on a node with a nil left child.")
		os.Exit(1)
	}

	// set y
	yTree := xTree.left

	// turn y's right subtree into x's left subtree
	xTree.left = yTree.right
	xTree.left.parent = xTree

	// link x's parent to y
	yTree.parent = xTree.parent
	if xTree.parent.value == 0 {
	} else if xTree.parent.left.value == xTree.value {
		xTree.parent.left = yTree
	} else {
		xTree.parent.right = yTree
	}

	// put x on y's right
	yTree.right = xTree
	yTree.right.parent = yTree

	if !t.root.IsRoot() {
		t.root = yTree
	}

	return 0
}

func (t *RBtree) Transplant(u *RBnode, v *RBnode) {
	if u.parent.value == -1 {
		t.root = v
	} else if u.value == u.parent.left.value {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func (t *RBtree) Display() string {
	return t.root.DisplayN()
}

func (n *RBnode) DisplayN() string {
	var s string = fmt.Sprintf("(%v/%v/%v", n.value, n.color, n.parent.value)
	if n.left != nil {
		s += ", "
		s += n.left.DisplayN()
	}
	if n.right != nil {
		if n.left == nil {
			s += ", , "
		} else {
			s += ", "
		}
		s += n.right.DisplayN()
	}
	s += ")"
	return s
}

func (t *RBtree) PrettyDisplay() string {
	return t.root.PrettyDisplayN()
}

func (n *RBnode) PrettyDisplayN() string {
	if n.value == -1 {
		return ""
	}
	var s string = fmt.Sprintf("(%v/%v", n.value, n.color)
	if n.left != nil {
		s += ", "
		s += n.left.PrettyDisplayN()
	}
	if n.right != nil {
		if n.left == nil {
			s += ", , "
		} else {
			s += ", "
		}
		s += n.right.PrettyDisplayN()
	}
	s += ")"
	return s
}
