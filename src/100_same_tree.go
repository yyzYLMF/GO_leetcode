/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p==nil || q==nil {
        if p != q {
            return false
        } else {
            return true
        }
    }
    if p.Val != q.Val {
        return false
    }
    if isSameTree(p.Left, q.Left) == false {
        return false
    }
    if isSameTree(p.Right, q.Right) == false {
        return false
    }
    return true
}
