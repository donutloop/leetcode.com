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
    public List<Integer> getLonelyNodes(TreeNode root) {
         var nodes = new ArrayList<Integer>();
         if (root == null) {
             return nodes;
         }

         countNodes(nodes, root.left, root.right);
         return nodes;
    }
    private void countNodes(ArrayList<Integer> nodes, TreeNode leftNode, TreeNode rightNode) {
        if (leftNode == null && rightNode == null) {
            return;
        } else if (leftNode == null && rightNode != null) {
            nodes.add(rightNode.val);
        } else if (leftNode != null && rightNode == null) {
            nodes.add(leftNode.val);
        }

        if (leftNode != null) countNodes(nodes, leftNode.left, leftNode.right);
        if (rightNode != null) countNodes(nodes, rightNode.left, rightNode.right);

        return;
    }
}