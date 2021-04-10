# My notes on binary search trees

A binary search tree is a rooted binary tree. What does *rooted* mean? It means that we designate one node as the root. What does *binary* mean? It means that each node has at most two children, called the left and right nodes. Since a binary search tree is ordered, we can use one for (relatively) quick searching or sorting.

Each node should store at least three pieces of information: some value, a pointer to its left child, and a pointer to its right value. The left and right pointers can be empty, in which case the node has no children. (A node with no children is a *leaf*.

In order to count as a binary search tree, the nodes must be ordered. Every node’s value should be greater than (or equal to) any value from a node to its left and less than (or equal to) any value from a node to its right. (Some binary search trees allow for duplicate values, and some do not. Hence, the fudging about "(or equal to).")

The nodes can store complex values, and even ones that are not intrinsically sortable, but in that case they should also have an associated simpler (and sortable) key to that value. In what follows, I am going to pretend that all binary search trees store integer values since that simplifies things.

In order to create a binary search tree, you need to think about several operations. Let’s start with insertion, deletion, and lookup.

## Lookup

In this case, we want to find a specific value. We begin with the root of the tree. If the root is empty, then we have no tree, and we return some appropriate value (probably nil, maybe an error). If the root of the tree is our value, then we return the root as our found node. If our value is less than the value of the root, we search again, starting with the root’s left node. If our value is greater than the root’s value, we search again, starting with the right node of the root. We continue doing that until either we find our value or we reach an empty subtree.

In the worst case scenario, you will have to search from the root to the leaf furthest from the root. Therefore, the time of the worst case is a function of the height of the tree. (The height of a tree is the number of steps from its root to the leaf furthest from the root.)

If you allow duplicates, then searching is slightly more complicated. First you have to decide where to insert duplicates: left or right? It doesn’t matter, but you have to pick. Once you choose, then you will know what to do when searching. Imagine we choose to put duplicates on the right. In that case, if the item we want is smaller than the value of our current node, we go left. Otherwise, we go right. That is, we keep going right until we hit the furthest rightward matching value. (Apparently, if we proceed like this, then we guarantee that searches are stable. We won’t reorder elements that are equal (in one sense).)

## Insertion

In order to insert a new node, we do something similar to lookup. If the root is empty, we create a new node, add the value to that, and return it. If the root is not empty, we travel left or right depending on the key in the root and what we want to insert. If we allow for duplicates, we need to continue left or right, depending on which side we want duplicates to go.

In general, I think I would insert in these steps:

+ If the root is empty, make a new node, add the key there, and that new node is our root.
+ If the root is not empty, and our new key is less than the key of the root, call insert again with the root’s left-child as our new starting point. 
+ If the root is not empty, and our new key is equal to our greater than the key of the root, call insert again with the root’s right-child as our new starting point.
