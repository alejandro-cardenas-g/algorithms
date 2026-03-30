class TreeNode {
  public left: TreeNode | null = null;
  public right: TreeNode | null = null;

  constructor(public value: number) {}
}

const root = new TreeNode(1);
const node2 = new TreeNode(2);
const node3 = new TreeNode(3);
const node4 = new TreeNode(4);
const node5 = new TreeNode(5);
const node6 = new TreeNode(6);
const node7 = new TreeNode(7);

root.left = node3;
root.right = node6;

node3.left = node2;
node3.right = node4;

node6.left = node5;
node6.right = node7;

const reversex = (node: TreeNode | null) => {
  if (node === null) return node;

  const auxLeft = node.left;

  node.left = reversex(node.right);
  node.right = reversex(auxLeft);
  return node;
};

console.log(reversex(root));
