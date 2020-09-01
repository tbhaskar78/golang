/****************************************

* File Name : bst.go

* Creation Date : 28-08-2020

* Last Modified : Tuesday 01 September 2020 02:42:23 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
    "fmt"
)

/* LEFT and RIGHT Dir in the tree */
var LEFT = 0
var RIGHT = 1

/* define the node struct */
type node struct {
    data int
    link[2] *node
}

/* find max value in the tree */
func findmax(root *node) *node {
    if root == nil {
        return nil
    } else if root.link[RIGHT] == nil {
        return root
    } else {
        return findmax(root.link[RIGHT])
    }
}

/* find min value in the tree */
func findmin(root *node) *node {
    if root == nil {
        return nil
    } else if root.link[LEFT] == nil {
        return root
    } else {
        return findmin(root.link[LEFT])
    }
}

/* delete a node from the tree and adjust the tree */
func deletenode(root *node, val int) *node {
    if root == nil {
        return root
    }
    if root.data > val {
        root.link[LEFT] = deletenode(root.link[LEFT], val)
    } else if root.data < val {
        root.link[RIGHT] = deletenode(root.link[RIGHT], val)
    } else {                                                          // matched the data
        if root.link[LEFT] == nil && root.link[RIGHT] == nil {        // case 1, no children
            root = nil;
        } else if root.link[LEFT] != nil && root.link[RIGHT] == nil { // case 2.a, one child - left
            root = root.link[LEFT]
        } else if root.link[LEFT] == nil && root.link[RIGHT] != nil { // case 2.b, one child - right
            root = root.link[RIGHT]
        } else {                                                      // case 3, both children
            root.data = findmin(root.link[RIGHT]).data
            root.link[RIGHT] = deletenode(root.link[RIGHT], root.data)
        }
    }

    return root
}

/* insert the newly created node in the right place in BST */
func insert(root *node, val int) *node {
    if root == nil {
        root = &node{data: val,  }  // initialize the node
    } else {
        if root.data > val {
            root.link[LEFT] = insert(root.link[LEFT], val)
        } else if (root.data < val) {
            root.link[RIGHT] = insert(root.link[RIGHT], val)
        } else if (root.data == val) {
            defer fmt.Printf("\nDuplicate data detected while inserting value [%v]\n", val)
        }
    }
    return root
}

/* finds the maxdepth or height of the tree */
func maxdepth(root *node) int {
    if root == nil {
        return 0
    }
    ldepth := maxdepth(root.link[LEFT])
    rdepth := maxdepth(root.link[RIGHT])

    if ldepth > rdepth {
        return (ldepth+1)
    } else {
        return (rdepth+1)
    }
}

/* print the tree inorder */
func inorder(root *node) {
    if root == nil {
        return
    }
    inorder(root.link[LEFT])
    fmt.Printf("%v|", root.data)
    inorder(root.link[RIGHT])
}

/* print the tree postorder */
func postorder(root *node) {
    if root == nil {
        return
    }
    postorder(root.link[LEFT])
    postorder(root.link[RIGHT])
    fmt.Printf("%v|", root.data)
}

/* recursively print in preorder */
func preorder(root *node) {
    if root == nil {
        return
    }
    fmt.Printf("%v|", root.data)
    preorder(root.link[LEFT])
    preorder(root.link[RIGHT])
}

/* recursively print in levelorder */
func levelorder(root *node, level int) {
    if root == nil {
        return
    }
    if level == 1 {
        fmt.Printf("%v|", root.data)
    }
    levelorder(root.link[LEFT], level)
    levelorder(root.link[RIGHT], level)
}

/* recursively print in spiralorder */
func spiralorder(root *node, level, toggle int) {
    if root == nil {
        return
    }
    if level == 1 {
        fmt.Printf("%v|", root.data)
    }

   if toggle == 1 {
       spiralorder(root.link[LEFT], level-1, toggle)
       spiralorder(root.link[RIGHT], level-1, toggle)
   } else if toggle == -1 {
       spiralorder(root.link[RIGHT], level-1, toggle)
       spiralorder(root.link[LEFT], level-1, toggle)
   }
}

func main() {
    var toggle int = 1
    var tree_root *node
    array := []int{1, 9, 8, 6, 100, 150, 20, 25, 9, 16, -1, 3, 5, 7, 11, 13, 210, 23, 45, 38, 93, 0, }

    //fmt.Printf("Inserting the following in tree %v\n", cap(array))
    for i := 0 ; i < cap(array); i++ {
        //fmt.Printf("%v, ", array[i])
        tree_root = insert(tree_root, array[i])
    }

    /* print the tree in preorder, post order, level and spiral order */
    fmt.Printf("\nPreOrder : \n")
    preorder(tree_root)
    fmt.Printf("\n\nPostOrder : \n")
    postorder(tree_root)
    fmt.Printf("\n\nLevel order: \n")
    for i := 0; i <= maxdepth(tree_root); i++ {
        levelorder(tree_root, i)
    }
    fmt.Printf("\n\nSpiral Order: \n")
    for i := 0; i <= maxdepth(tree_root); i++ {
        toggle = toggle * (-1)
        spiralorder(tree_root, i, toggle)
    }

    fmt.Printf("\n\nInOrder : \n")
    inorder(tree_root)
    //fmt.Printf("Deleting 1  and 93 from the nodes \n")
    deletenode(tree_root, 1)
    deletenode(tree_root, 93)
    fmt.Printf("\n\nInorder after deleting 1 and 93: \n")
    inorder(tree_root)

    fmt.Printf("\n\nheight of tree is %v\n", maxdepth(tree_root))
    fmt.Printf("min in the tree is %v\n", findmin(tree_root).data)
    fmt.Printf("max in the tree is %v\n", findmax(tree_root).data)
}
