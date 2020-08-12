package containers

import "fmt"

// 数据结构---二叉树的节点
type CBinaryTreeNode struct {

	value 	interface{}			// 数据

	left 	*CBinaryTreeNode	// 左子节点

	right 	*CBinaryTreeNode	// 又子节点

}

// 数据结构---二叉树
type CBinaryTree struct {

	depth 	int 				// 树的深度

	size 	int					// 节点数

	root 	*CBinaryTreeNode	// 根节点

}

func NewCBinaryTreeNode(v interface{}) *CBinaryTreeNode {
	return &CBinaryTreeNode{
				value: v,
				left: nil,
				right: nil,
			}
}

func NewCBinaryTree() *CBinaryTree {
	return new(CBinaryTree).Init()
}

func (self *CBinaryTree)Init() *CBinaryTree {
	self.size = 0
	self.root = nil

	return self
}

func (self *CBinaryTree) Size() int {
	return self.size
}

func (self *CBinaryTree) Root() *CBinaryTreeNode {
	return self.root
}

func (self *CBinaryTree) IsEob(node *CBinaryTreeNode) bool {
	if node == nil {
		return true
	}

	return false
}

func (self *CBinaryTree) IsLeaf(node *CBinaryTreeNode) bool {
	if node.left == nil && node.right == nil {
		return true
	}

	return false
}

func (self *CBinaryTree) Data(node *CBinaryTreeNode) interface{} {
	return node.value
}

func (self *CBinaryTree) Left(node *CBinaryTreeNode) *CBinaryTreeNode {
	return node.left
}

func (self *CBinaryTree) Right(node *CBinaryTreeNode) *CBinaryTreeNode {
	return node.right
}

func (self *CBinaryTree) Insert(v interface{}) int {
	var node *CBinaryTreeNode
	var prev *CBinaryTreeNode
	var direction int

	node = self.root
	direction = 0

	for !self.IsEob(node) {
		prev = node

		if v.(int) == self.Data(node).(int) {
			return -1
		}else if v.(int) < self.Data(node).(int){
			node = self.Left(node)
			direction = 1
		}else {
			node = self.Right(node)
			direction = 2
		}
	}

	data := v

	if direction == 0 {
		return self.LeftInsert(nil, data)
	}

	if direction == 1 {
		return self.LeftInsert(prev, data)
	}

	if direction == 2 {
		return self.RightInsert(prev, data)
	}

	return -1


}

func (self *CBinaryTree) LeftInsert(node *CBinaryTreeNode, v interface{}) int {
	var newNode *CBinaryTreeNode
	var position **CBinaryTreeNode

	if node == nil {
		if self.Size() > 0 {
			return -1
		}

		position = &self.root
	}else {
		if self.Left(node) != nil {
			return -1
		}

		position = &node.left
	}

	newNode = NewCBinaryTreeNode(v)
	*position = newNode

	self.size++

	return 0
}

func (self *CBinaryTree) RightInsert(node *CBinaryTreeNode, v interface{}) int {
	var newNode *CBinaryTreeNode
	var position **CBinaryTreeNode

	if node == nil {
		if self.size > 0 {
			return -1
		}

		position = &self.root
	}else {
		if self.Right(node) != nil {
			return -1
		}

		position = &node.right
	}

	newNode = NewCBinaryTreeNode(v)
	*position = newNode

	self.size++

	return 0
}

func (self *CBinaryTree) LeftRemove(node *CBinaryTreeNode) {
	var position **CBinaryTreeNode

	if self.Size() == 0 {
		return
	}

	if node == nil {
		position = &self.root
	}else {
		position = &node.left
	}

	if *position != nil {
		self.LeftRemove(*position)
		self.RightRemove(*position)

		*position = nil

		self.size--
	}

	return
}

func (self *CBinaryTree) RightRemove(node *CBinaryTreeNode) {
	var position **CBinaryTreeNode

	if self.Size() == 0 {
		return
	}

	if node == nil {
		position = &self.root
	}else {
		position = &node.right
	}

	if *position != nil {
		self.LeftRemove(*position)
		self.RightRemove(*position)

		*position = nil

		self.size--
	}

	return
}

func (self *CBinaryTree) PreorderPrint(node *CBinaryTreeNode) {
	if node == nil {
		return
	}

	fmt.Println(node.value)

	if self.Left(node) != nil {
		self.PreorderPrint(node.left)
	}

	if self.Right(node) != nil {
		self.PreorderPrint(node.right)
	}

	return
}

func (self *CBinaryTree) InorderPrint(node *CBinaryTreeNode) {
	if node == nil {
		return
	}

	if self.Left(node) != nil {
		self.InorderPrint(node.left)
	}

	fmt.Println(node.value)

	if self.Right(node) != nil {
		self.InorderPrint(node.right)
	}

	return
}

func (self *CBinaryTree) PostorderPrint(node *CBinaryTreeNode) {
	if node == nil {
		return
	}

	if self.Left(node) != nil {
		self.PostorderPrint(node.left)
	}

	if self.Right(node) != nil {
		self.PostorderPrint(node.right)
	}

	fmt.Println(node.value)

	return
}
