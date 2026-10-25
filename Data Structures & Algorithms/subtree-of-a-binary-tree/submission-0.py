# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:   
    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        def sameTree(
            tree1: Optional[TreeNode],
            tree2: Optional[TreeNode]
        ) -> bool:
            if not tree1 and not tree2:
                return True

            if not tree1 or not tree2:
                return False

            return (
                tree1.val == tree2.val
                and sameTree(tree1.left, tree2.left)
                and sameTree(tree1.right, tree2.right)
            )
        
        if not subRoot:
            return True

        if not root:
            return False

        return (
            sameTree(root, subRoot)
            or self.isSubtree(root.left, subRoot)
            or self.isSubtree(root.right, subRoot)
        )