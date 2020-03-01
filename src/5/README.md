# Problem 5

Smallest multiple
   
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

**What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?**

---

Trying to figure out a way to reduce number of operations needed. Instead of dividing each number under test by all numbers {1:20}, instead just divide by 20, and all primes that are not factors of 20. i.e. 7, 11, 13, 17, and 19

Or, perhaps, increment number to test by 20, then test against 7, 11, 13, 17, and 19.

No, this didn't work for some reason, but iterating through the entire range [2:19] did.

To make this faster, there's a concept I'm missing. Something like "greatest common product bounded by some ceiling". Not sure if that's actually a thing. For instance, if number divisible by 3, not necessarily divisible by 9. But, if number divisible by 18, then definitely divisible by 9, 6, 3 also.
