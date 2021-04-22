
// @Title: 删除链表中的节点 (Delete Node in a Linked List)
// @Author: 846188037@qq.com
// @Date: 2021-02-02 21:08:36
// @Runtime: 4 ms
// @Memory: 2.8 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    node.Val=node.Next.Val
    node.Next=node.Next.Next
}
