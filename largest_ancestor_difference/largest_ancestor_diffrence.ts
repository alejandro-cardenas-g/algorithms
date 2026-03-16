class TreeNode {
  val: number;
  left: TreeNode | null;
  right: TreeNode | null;
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = val === undefined ? 0 : val;
    this.left = left === undefined ? null : left;
    this.right = right === undefined ? null : right;
  }
}

function maxAncestorDiff(root: TreeNode | null): number {
  return getDifference([], root);
}

function getDifference(
  listOfAncestors: TreeNode[],
  node: TreeNode | null,
): number {
  if (node === null) return 0;

  let difference = 0;
  for (const ancestor of listOfAncestors) {
    difference = Math.max(difference, Math.abs(ancestor.val - node.val));
  }

  const leftDifference = getDifference([...listOfAncestors, node], node.left);
  const rigthDifference = getDifference([...listOfAncestors, node], node.right);

  return Math.max(difference, leftDifference, rigthDifference);
}
