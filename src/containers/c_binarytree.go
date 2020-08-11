package containers

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

func (self *CBinaryTree) Insert(interface{}) {

}
