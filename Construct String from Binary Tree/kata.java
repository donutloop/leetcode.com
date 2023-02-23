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
    public String tree2str(TreeNode root) {
        var pattern = new StringBuilder("");
        generatePattern(root, pattern);
        return pattern.toString();
    }
    private void generatePattern(TreeNode node, StringBuilder pattern) {
           pattern.append(Integer.valueOf(node.val));

           if (node.left == null && node.right != null) {
              pattern.append('(');
              pattern.append(')');
           }else if (node.left != null) {
              pattern.append('(');
              generatePattern(node.left, pattern);
              pattern.append(')');
           }

            if (node.right != null) {
              pattern.append('(');
              generatePattern(node.right, pattern);
              pattern.append(')');
           }

        return;
    }
}