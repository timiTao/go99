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