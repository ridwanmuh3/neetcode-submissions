# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        # visited = set()
        # curr = head
        # while curr != None:
        #     if curr in visited:
        #         return True
        #     else:
        #         visited.add(curr)
        #     curr = curr.next
        # return False
        slow = head
        fast = head
        while fast != None and fast.next != None:
            slow = slow.next
            fast = fast.next.next
            if slow is fast:
                return True
        return False