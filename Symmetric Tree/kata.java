/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public boolean isSymmetric(TreeNode root) {
        return verify(root.left, root.right);
    }
    private boolean verify(TreeNode nodeLeft, TreeNode nodeRight) {
        if (nodeLeft == null && nodeRight == null) {
            return true;
        } else if (nodeLeft == null || nodeRight == null) {
            return false;
        }else if (nodeLeft.val != nodeRight.val) {
            return false;
        }
        var a = verify(nodeLeft.left, nodeRight.right);
        var b = verify(nodeLeft.right, nodeRight.left);
        return a & b;
    }
}