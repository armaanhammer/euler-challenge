Problem 3

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

---

I will need to work through this algorithm a bit.

Two tasks:
- step up through the natural numbers, identifying if that number is a prime,
- divide the original query number (and subsequent dividends) by that prime number to see if it is a factor.

I imagine some psudocode for this would be like"

```
num = original num to test and factor
natural = num to test against

natural starts at 1

WHILE some condition, maybe num itself is a prime? {
	test num against natural
	IF natural is a prime {
		IF natural is factor of num {
			divide num by natural
		}
		increment natural by ...

```

Found this algorithm at: https://www.geeksforgeeks.org/print-all-prime-factors-of-a-given-number/

Following are the steps to find all prime factors.
1) While n is divisible by 2, print 2 and divide n by 2.
2) After step 1, n must be odd. Now start a loop from i = 3 to square root of n. While i divides n, print i and divide n by i. After i fails to divide n, increment i by 2 and continue.
3) If n is a prime number and is greater than 2, then n will not become 1 by above two steps. So print n if it is greater than 2.
