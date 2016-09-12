/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traversal(root *TreeNode, maxLength *int, height int) {
    if root == nil {
        if height > (*maxLength) {
            (*maxLength) = height
        }
        return
    }
    traversal(root.Left, maxLength, height+1)
    traversal(root.Right, maxLength, height+1)
    return
}

func maxDepth(root *TreeNode) int {
    maxLength := -1
    traversal(root, &maxLength, 0)
    return maxLength
}
