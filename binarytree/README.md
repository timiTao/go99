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

* P55. [Construct completely balanced binary trees](/binarytree/main.go#L53) / [test](/binarytree/p55_test.go#L3)
```
In a completely balanced binary tree, the following property holds for every node: The number of nodes in its left subtree and the number of nodes in its right subtree are almost equal, which means their difference is not greater than one.

Write a predicate cbal_tree/2 to construct completely balanced binary trees for a given number of nodes. The predicate should generate all solutions via backtracking. Put the letter 'x' as information into all nodes of the tree.
Example:
?- cbal_tree(4,T).
T = t(x, t(x, nil, nil), t(x, nil, t(x, nil, nil))) ;
T = t(x, t(x, nil, nil), t(x, t(x, nil, nil), nil)) ;
etc......No
```

* P56. [Symmetric binary trees](/binarytree/main.go#L120) / [test](/binarytree/p56_test.go#L3) 