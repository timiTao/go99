### Binary Tree

* P54. [~~Check whether a given term represents a binary tree~~](/binarytree/p54_test.go#L3) 
```
Write a predicate istree/1 which succeeds if and only if its argument is a Prolog term representing a binary tree.
Example:
?- istree(t(a,t(b,nil,nil),nil)).
Yes
?- istree(t(a,t(b,nil,nil))).
No
```

As I don't find other way to move with `struct`. This will force always have valid tree.

* P55. [Construct completely balanced binary trees](/binarytree/p55_test.go#L3)
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

* P61A [Collect the leaves of a binary tree in a list](/binarytree/main.go) / [test](/binarytree/p61a_test.go#L3)
```
A leaf is a node with no successors. Write a predicate leaves/2 to collect them in a list. 

% leaves(T,S) :- S is the list of all leaves of the binary tree T
```