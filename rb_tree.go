package main

import (
	"fmt"
)

// created by Matt Weeden
// 9/5/16
//
// This program implements a Red-Black tree type

type RBtree struct {
	root *RBnode
}

type RBnode struct {
	value  int
	color  string
	parent *RBnode
	left   *RBnode
	right  *RBnode
}

func (t *RBtree) Insert(z int) int {
	t.root.InsertR(z)
	t.InsertFix(z)
	return 0
}

func (t *RBnode) InsertR(z int) int {
	if t.value == -1 {
		t.value = z
		t.color = "red"
		t.left = new(RBnode)
		t.left.InitializeN(t)
		t.right = new(RBnode)
		t.right.InitializeN(t)
	} else if z < t.value {
		t.left.InsertR(z)
	} else if z > t.value {
		t.right.InsertR(z)
	} else {
		fmt.Println("Insert was called in an incorrect way.")
		return 1
	}

	return 0
}

func (t *RBtree) InsertFix(z int) int {
	zTree := t.FindNode(z)

	for zTree.parent.color == "red" {
		if zTree.parent.value == zTree.parent.parent.left.value {
			yTree := zTree.parent.parent.right
			if yTree.color == "red" {
				zTree.parent.color = "black"
				yTree.color = "black"
				zTree.parent.parent.color = "red"
				zTree = zTree.parent.parent
			} else {
				if zTree.value == zTree.parent.right.value {
					zTree = zTree.parent
					t.LeftRotate(zTree.value)
					zTree = t.FindNode(zTree.value)
				}
				zTree.parent.color = "black"
				zTree.parent.parent.color = "red"
				t.RightRotate(zTree.parent.parent.value)
			}
		} else {
			yTree := zTree.parent.parent.left
			if yTree.color == "red" {
				zTree.parent.color = "black"
				yTree.color = "black"
				zTree.parent.parent.color = "red"
				zTree = zTree.parent.parent
			} else {
				if zTree.value == zTree.parent.left.value {
					zTree = zTree.parent
					t.RightRotate(zTree.value)
					zTree = t.FindNode(zTree.value)
				}
				zTree.parent.color = "black"
				zTree.parent.parent.color = "red"
				t.LeftRotate(zTree.parent.parent.value)
			}
		}
	}
	t.root.color = "black"
	return 0
}

func (t *RBtree) Delete(z int) int {
	zTree := t.FindNode(z)
	yTree := zTree
	yColor := yTree.color
	var xTree *RBnode
	if zTree.left.value == -1 {
		xTree = zTree.right
		t.Transplant(zTree, zTree.right)
	} else if zTree.right.value == -1 {
		xTree = zTree.left
		t.Transplant(zTree, zTree.left)
	} else {
		yTree = t.FindMin(zTree.right.value)
		yColor = yTree.color
		xTree = yTree.right
		if yTree.parent == zTree {
			xTree.parent = yTree
		} else {
			t.Transplant(yTree, yTree.right)
			yTree.right = zTree.right
			yTree.right.parent = yTree
		}
		t.Transplant(zTree, yTree)
		yTree.left = zTree.left
		yTree.left.parent = yTree
		yTree.color = zTree.color
	}
	if yColor == "black" {
		t.DeleteFix(xTree)
	}

	return 0
}

func (t *RBtree) DeleteFix(xTree *RBnode) int {
	for xTree.parent.value != 0 && xTree.color == "black" {
		if xTree == xTree.parent.left {
			wTree := xTree.parent.right
			if wTree.color == "red" {
				wTree.color = "black"
				xTree.parent.color = "red"
				t.LeftRotate(xTree.parent.value)
				wTree = xTree.parent.right
			}
			if wTree.left.color == "black" && wTree.right.color == "black" {
				wTree.color = "red"
				xTree = xTree.parent
			} else {
				if wTree.right.color == "black" {
					wTree.left.color = "black"
					wTree.color = "red"
					t.RightRotate(wTree.value)
					wTree = xTree.parent.right
				}
				wTree.color = xTree.parent.color
				xTree.parent.color = "black"
				wTree.right.color = "black"
				t.LeftRotate(xTree.parent.value)
				xTree = t.root
			}
		} else {
			wTree := xTree.parent.left
			if wTree.color == "red" {
				wTree.color = "black"
				xTree.parent.color = "red"
				t.RightRotate(xTree.parent.value)
				wTree = xTree.parent.left
			}
			if wTree.right.color == "black" && wTree.left.color == "black" {
				wTree.color = "red"
				xTree = xTree.parent
			} else {
				if wTree.left.color == "black" {
					wTree.right.color = "black"
					wTree.color = "red"
					t.LeftRotate(wTree.value)
					wTree = xTree.parent.left
				}
				wTree.color = xTree.parent.color
				xTree.parent.color = "black"
				wTree.left.color = "black"
				t.RightRotate(xTree.parent.value)
				xTree = t.root
			}
		}
	}
	xTree.color = "black"
	return 0
}

func main() {
	fmt.Printf("1(b)\n====================================\n")

	var t RBtree
	t.Initialize()

	t.Insert(41)
	fmt.Printf("Insert(41): ")
	fmt.Println(t.PrettyDisplay())
	t.Insert(38)
	fmt.Printf("Insert(38): ")
	fmt.Println(t.PrettyDisplay())
	t.Insert(31)
	fmt.Printf("Insert(31): ")
	fmt.Println(t.PrettyDisplay())
	t.Insert(12)
	fmt.Printf("Insert(12): ")
	fmt.Println(t.PrettyDisplay())
	t.Insert(19)
	fmt.Printf("Insert(19): ")
	fmt.Println(t.PrettyDisplay())
	t.Insert(8)
	fmt.Printf("Insert(8): ")
	fmt.Println(t.PrettyDisplay())
	t.Delete(12)
	fmt.Printf("Delete(12):")
	fmt.Println(t.PrettyDisplay())
	t.Insert(32)
	fmt.Printf("Insert(32): ")
	fmt.Println(t.PrettyDisplay())
	t.Delete(41)
	fmt.Printf("Delete(41):")
	fmt.Println(t.PrettyDisplay())

	fmt.Printf("\n1(c)\n====================================\n")

	t.Initialize()

	t.Insert(834)
	t.Insert(807)
	t.Insert(512)
	t.Insert(882)
	t.Insert(127)
	t.Insert(675)
	t.Insert(75)
	t.Insert(216)
	t.Insert(822)
	t.Insert(249)
	t.Insert(114)
	t.Insert(689)
	t.Insert(625)
	t.Insert(974)
	t.Insert(221)
	t.Insert(92)
	t.Insert(374)
	t.Insert(123)
	t.Insert(838)
	t.Insert(930)
	t.Insert(654)
	t.Insert(806)
	t.Insert(234)
	t.Insert(381)
	t.Delete(127)
	t.Delete(221)
	fmt.Println(t.PrettyDisplay())

	fmt.Printf("\n2(a)\n====================================\n")

	t.Initialize()

	t.Insert(14)
	t.Insert(24)
	t.Insert(46)
	t.Insert(26)
	t.Insert(75)
	fmt.Println(t.PrettyDisplay())
	t.Insert(27)
	fmt.Printf("Insert(27)\n")
	fmt.Println(t.PrettyDisplay())
	t.Delete(27)
	fmt.Printf("Delete(27)\n")
	fmt.Println(t.PrettyDisplay())
}
