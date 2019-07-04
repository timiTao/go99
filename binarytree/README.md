### Binary Tree

* P54. ~~Check whether a given term represents a binary tree~~
```
Write a predicate istree/1 which succeeds if and only if its argument is a Prolog term representing a binary tree.
Example:
?- istree(t(a,t(b,nil,nil),nil)).
Yes
?- istree(t(a,t(b,nil,nil))).
No
```

As I don't find other way to move with `struct`. This will force always have valid tree.

* P55. [Construct completely balanced binary trees](/binarytree/p55_test.go#L3) / [test](/binarytree/p55_test.go#L3)
```
In a completely balanced binary tree, the following property holds for every node: The number of nodes in its left subtree and the number of nodes in its right subtree are almost equal, which means their difference is not greater than one.

Write a predicate cbal_tree/2 to construct completely balanced binary trees for a given number of nodes. The predicate should generate all solutions via backtracking. Put the letter 'x' as information into all nodes of the tree.
Example:
?- cbal_tree(4,T).
T = t(x, t(x, nil, nil), t(x, nil, t(x, nil, nil))) ;
T = t(x, t(x, nil, nil), t(x, t(x, nil, nil), nil)) ;
etc......No
```

* P61. [Count the leaves of a binary tree](/binarytree/main.go) / [test](/binarytree/p61_test.go#L3)
```
A leaf is a node with no successors. Write a predicate count_leaves/2 to count them. 

% count_leaves(T,N) :- the binary tree T has N leaves
```

* P61A. [Collect the leaves of a binary tree in a list](/binarytree/main.go) / [test](/binarytree/p61a_test.go#L3)
```
A leaf is a node with no successors. Write a predicate leaves/2 to collect them in a list. 

% leaves(T,S) :- S is the list of all leaves of the binary tree T
```

* P62. [Collect the internal nodes of a binary tree in a list](/binarytree/main.go) / [test](/binarytree/p62_test.go#L3)
```
An internal node of a binary tree has either one or two non-empty successors. Write a predicate internals/2 to collect them in a list. 

% internals(T,S) :- S is the list of internal nodes of the binary tree T.
```

* P62B. [Collect the nodes at a given level in a list](/binarytree/main.go) / [test](/binarytree/p62b_test.go#L3)
```
A node of a binary tree is at level N if the path from the root to the node has length N-1. The root node is at level 1. Write a predicate atlevel/3 to collect all nodes at a given level in a list. 

% atlevel(T,L,S) :- S is the list of nodes of the binary tree T at level L

Using atlevel/3 it is easy to construct a predicate levelorder/2 which creates the level-order sequence of the nodes. However, there are more efficient ways to do that.
```