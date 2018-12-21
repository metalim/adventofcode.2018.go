# Advent of Code 2018 in Go

These are my soluions for [Advent of Code 2018](https://adventofcode.com/2018/) written in Go.

The goal is to learn quirks & tricks of Go, which are new to me, compared to concise dynamically typed languages like Coffeescript, that I'm familiar with, while keeping the code concise, and not go the wordy Java-way.

## "Go gotchas"

Quirks found so far:

* ### Strings are **arrays of bytes**, encoded in **UTF-8**

  Iterating over string with index variable works for ASCII characters, but leads to bugs with Unicode characters. For the moment I try to write Unicode-safe code even in time-restricted conditions, by using one of the following:
  * iterating over string with `for i,r := range s {}`;
  * converting strings to slices of runes (`sr := []rune(s)`), so each character can be indexed easily;
  * or by extracting substrings with regexes.

* ### Arrays are **values**

  And accessing a map returns copy of a mapped value. Example from day 4:

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

* ### There are not Sets

  In Go there are not Set structures, but they are easy to emulate with maps:

  ```go
  type sset map[string]bool // "set" of strings

  s := sset{}
  s["foo"] = true
  s["bar"] = true
  // set of "foo" and "bar"

  delete(s, "foo")
  s["baz"] = true
  // set of "bar" and "baz"

  if s["bar"] {/* "bar" is in set */}

  for w := range s {/* iterate set */}
  ```

  There is only one restriction: data types have to be comparable (`a==b`). So, no slices or maps allowed, but pointers to them are ok.

* ### Everything is passed by value

  There is no pass-by-reference in Go. The only hardcoded exception are maps. Maps are literally pointers to hmap structure, and in early Go versions they were created as pointers (`var m *map[int]int`), that were missing dereference mechanics. For that reason, in later versions the `*` was removed from map syntax.

  The rest, as with arrays, is passed by value.

  ```go
  type Some struct {n int}
  func inc(a Some){
    a.n++
  }

  b := Some{0}
  inc(b)
  fmt.Println(b.n) // still 0
  ```

  This has to be taken into account, when you iterate:

  ```go
  slice := []struct{n int}{/*...*/}

  for _, obj := range slice {
    obj.n++ // woops, we're changing a copy !
  }

  // proper way
  for i := range slice {
    slice[i].n++
  }

  ```

* ### Imports work with folders, not single files

  For that reason you should not have several files in same folder, containing same function definition. For example, golint complains that I have `func main()` redefined, because I keep all Advent of Code files in single folder. While this is not acceptable in larger projects, Advent of Code solutions are just single files, so I decided to ignore this rule, instead of creating separate folder for each day. It works, and is more accessible.

* ### Type names and variable names do not conflict (with rare exceptions)

  While this is not reader-friendly, same name can mean type and variable:

  ```go
  type some [2]int
  some := some{1, 2}
  some[0] = 3
  ```

* ### Selectors can work with pointers to struct

  Comes in handy for looped list operations in day 9:

  ```go
  type list struct {
    val int
    prev, next *list
  }
  l := &list{val: 0}
  l.prev = l // equal to (*l).prev = l
  l.next = l // equal to (*l).next = l
  ```

* ### Embedding works similar to inheritance

  That is: all fields and methods are inherited.

  Day 13:

  ```go
  type point struct {
    x, y int
  }

  func (a point) collides(b point) bool {
    return a==b
  }

  type cart struct {
    point
    id int
  }

  a := cart{point{5, 7}, 1}
  b := cart{point{5, 7}, 4}

  if a.x == b.x && a.y == b.y {}
  // or
  if a.point == b.point {}
  // are equivalent to
  if a.point.x == b.point.x && a.point.y == b.point.y {}

  crashed = a.collides(b)
  // is same as
  crashed = a.point.collides(b.point)
  ```

* ### Type definition, on the other hand, hides methods "inherited" from another type

  ```go
    type point struct {
      x, y int
    }
    func (a point) collides(b point) bool {
      return a==b
    }

    type pos point

    a := pos{1,2}
    b := pos{3,4}
    a.collides(b) // woops: a.collides undefined (type pos has no field or method collides)
  ```

  Solution is type aliasing, which just creates a new name for the same type:

  ```go
    type pos = point // note the `=`

    a := pos{1,2}
    b := pos{3,4}
    a.collides(b) // works

  ```

## Advent of Code specifics

* One has to remember what Advent of Code is: tasks with quickly implemented solutions. Solution submissions in leaderboard take anywhere from 1 to tens of minutes after task unlock. With that in mind, you have to watch for clues in task description, that allow you to get away from general case solution, that would require significantly more time to implement.

  * Day 4 example

    Timestamps are written in `YYYY-MM-DD HH:mm` format. But in task description it is said that all sleeping periods are in single hour: from 0:00 to 0:59, even though guard shift starts at 23:**. Hence it doesn't make sense to parse the whole timestamp, but only last 2 digits. For same reason there's no point in calculating and storing time ranges, as a single 60 minute array would be sufficient.

* Day 6 solution **checker** had a bug, that caused 20% of part1 solutions and 33% of part2 solutions to be detected as invalid ones. Bug was fixed 1 hour 42 minutes after unlock. My github-linked account had no issues, but my google-linked account was rejecting both parts. First I "cheated", by manually selecting the next largest area (3144 instead of 3358), thinking my solver had an issue detecting infinity of the area. But trying to submit solution of part2 confirmed Advent of Code has an issue, as the solution is trivial and does not depend on edge detection. So I added `_print()` function, to visualize the map and confirm my sanity for part1 solution. After the bug was discovered and fixed, Advent of Code staff removed Day 6 scores from affecting global Leaderboard.

* Day 19 task part 2 requires to analyze assembly code, rewrite it in higher-level language and refactor. This has to be done manually. In that specific case code was calculating sum of all (not only prime) factors of a number. Rewritting nested loop makes program run in <100ms instead of 10 years, refactoring outer loop brings it down to <100µs.

## Input

Multiple inputs are provided in global `ins` variable:

```go
ins := map[string]string{
  "a":`1`,
  "b":`2`,
}
```

That's because I submit solutions to 2 different accounts: one linked to github, another - to google. I would stick to github one, but google account has private leaderboard from previous year, that I want to track, but have no joining link for it.
