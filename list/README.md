### List

All solutions are in [main.go](/list/main.go) file

Test for solutions:

* P01. [Find the last element of a list.](/list/p01_test.go#L3)
* P02. [Find the last but one element of a list.](/list/p02_test.go#L3)
* P03. [Find the K'th element of a list.](/list/p03_test.go#L3)
* P04. [Find the number of elements of a list.](/list/p04_test.go#L3)
* P05. [Reverse a list.](/list/p05_test.go#L3)
* P06. [Find out whether a list is a palindrome.](/list/p06_test.go#L3)
```
A palindrome can be read forward or backward; e.g. [x,a,m,a,x].
```

* P07. [Flatten a nested list structure.](/list/p07_test.go#L3)
```
Transform a list, possibly holding lists as elements into a `flat' list by replacing each list with its elements (recursively).

Example:
?- my_flatten([a, [b, [c, d], e]], X).
X = [a, b, c, d, e]

Hint: Use the predefined predicates is_list/1 and append/3
```

* P08. [Eliminate consecutive duplicates of list elements.](/list/p08_test.go#L3)
```
If a list contains repeated elements they should be replaced with a single copy of the element. The order of the elements should not be changed.

Example:
?- compress([a,a,a,a,b,c,c,a,a,d,e,e,e,e],X).
X = [a,b,c,a,d,e]
```