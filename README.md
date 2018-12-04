# Advent of Code 2018 in Go

These are my soluions for [Advent of Code 2018](https://adventofcode.com/2018/) written in Go.

The goal is to learn quirks & tricks of Go, witch are new to me, compared to concise dynamically typed languages like Coffeescript, that I'm familiar with.

## Go gotchas

Quirks found so far:

* ### Strings are **arrays of bytes**, encoded in **UTF-8**

  Iterating over string with index variable works for ASCII characters, but leads to bugs with Unicode characters. For the moment I try to write Unicode-safe code even in time-restricted conditions, by using one of the following:
  * iterating over string with `for i,r := range s {}`;
  * converting strings to slices of runes (`sr := []rune(s)`), so each character can be indexed easily;
  * or by extracting substrings with regexes.

* ### Arrays are **values**

  And accessing a map returns copy of a stored value. Example from day 4:

  ```go
  var guards map[int][60]int
  //...
  guards[id][i]++ // doesn't compile.
  ```

  Even if it would compile, it would produce wrong result, because each access to `guards[id]` returns **copy** of array.

  Solution is either to use slices, like so:

  ```go
  var guards map[int][]int
  guards[id] = make([]int,60)
  guards[id][i]++
  ```

   or use pointers to arrays:

   ```go
   var guards map[int]*[60]int
   guards[id] = &[60]int{}
   (*guards[id])[i]++
   ```

## Advent of Code specifics

One has to remember what Advent of Code is: tasks with quick solutions. Solutions in leaderboard take anywhere from 1 to tens of minutes. With that in mind, you have to watch for clues in task description, that allow you to get away from general case solution, that would require significantly more time to implement.

* ### Day 4 example

  Timestamps are written in `YYYY-MM-DD HH:mm` format. But in task description it is said that all sleeping periods are in single hour: from 0:00 to 0:59, even though guard shift starts at 23:**. Hence it doesn't make sense to parse the whole timestamp, but only last 2 digits. For same reason there's no point in calculating and storing time ranges, as a single 60 minute array would be sufficient.

## Input

Multiple inputs are provided in global variable

```go
var ins map[string]string
```

That's because I submit solutions to 2 different accounts: one linked to github, another - to google. I would stick to github one, but google account has private leaderboard from previous year, that I want to track, but have no joining link for it.
