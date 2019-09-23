### Arithmetic

* P31. [Determine whether a given integer number is prime.](/arithmetic/p31.go#L3) / [test](/arithmetic/p31_test.go)
* P32. [Determine the greatest common divisor of two positive integer numbers.](/arithmetic/p32.go#L3) / [test](/arithmetic/p32_test.go) / [WIKI](https://en.wikipedia.org/wiki/Euclidean_algorithm)
* P33. [Determine whether two positive integer numbers are coprime.](/arithmetic/p33.go#L3) [test](/arithmetic/p33_test.go#L3)
```
Two numbers are coprime if their greatest common divisor equals 1.
Example:
?- coprime(35, 64).
Yes
```

* P34. [Calculate Euler's totient function phi(m).](/arithmetic/p34.go#L3) [test](/arithmetic/p34_test.go#L3)
```
Euler's so-called totient function phi(m) is defined as the number of positive integers r (1 <= r < m) that are coprime to m.
Example: m = 10: r = 1,3,7,9; thus phi(m) = 4. Note the special case: phi(1) = 1.

?- Phi is totient_phi(10).
Phi = 4

Find out what the value of phi(m) is if m is a prime number. Euler's totient function plays an important role in one of the most widely used public key cryptography methods (RSA). In this exercise you should use the most primitive method to calculate this function (there are smarter ways that we shall discuss later).
```