# Problem 5

Smallest multiple
   
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
---

Trying to figure out a way to reduce number of operations needed. Instead of dividing each number under test by all numbers {1:20}, instead just divide by 20, and all primes that are not factors of 20. i.e. 7, 11, 13, 17, and 19

Or, perhaps, increment number to test by 20, then test against 7, 11, 13, 17, and 19.
