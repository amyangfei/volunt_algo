# coding: utf-8

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def oddEvenList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if head is None or head.next is None:
            return head
        pre, cur = head, head.next
        while cur and cur.next:
            tmp = pre.next
            pre.next = cur.next
            cur.next = cur.next.next
            pre.next.next = tmp
            pre = pre.next
            cur = cur.next
        return head
