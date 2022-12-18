# golang-leetcode-practices

## Introduction

  - try to summarize the solutions of leetcode problems in golang I have done.
  - you can run test in each problem folder

## Problems

### easy

 title

 - Roman to Integer

description

- Given a roman numeral, convert it to an integer.

hint
Since Roman numerals must be arranged from largest to smallest, if an arrangement is produced in the opposite order, the situation of the latter minus the former will occur. Therefore, we just need to process it in reverse order. If number[n-1] < number [n], then number[n-1] will be added, otherwise both will be added together.

what you need to know in golang

- create byte index map for roman numerals
 for loop from the end of the string
 add/subtract the number
