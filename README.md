# Задачи с экзамена по программированию в Т-Банк

## 1 задача
```go
package main

import "fmt"

func main() {
    var s string
    fmt.Scan(&s)

    flag := false

    for i := 0; i < len(s); i++ {
        subString := s[:i] + s[i+1:]
        if isPalindrome(subString) {
            flag = true
            break
        }
    }

    if flag {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

func isPalindrome(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}
```

## 2 задача
```go
package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	branches := make([][2]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&branches[i][0], &branches[i][1])
	}

	var req int
	fmt.Scan(&req)
	
	for i := 0; i < req; i++ {
		var t, d int
		fmt.Scan(&t, &d)
		t--

		a, b := branches[t][0], branches[t][1]
		if d < a {
			fmt.Println(a)
		} else {
			interval := (d - a + b - 1) / b
			next_train := a + interval*b
			fmt.Println(next_train)
		}
	}
}
```

## 3 задача
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	used := make(map[int]bool)
	count_unique := 0

	for _, num := range a {
		current := num
		for current > 0 {
			if !used[current] {
				used[current] = true
				count_unique++
				break
			}
			current /= 2
		}

		if current == 0 && !used[0] {
			used[0] = true
			count_unique++
		}
	}

	fmt.Println(count_unique)
}
```

## 4 задача
```go
package main

import (
	"fmt"
)

func countArithmeticSubarrays(n int, arr []int) int {
	count := 0

	for l := 0; l < n; l++ {
		for r := l + 2; r < n; r++ {
			subarray := arr[l : r+1]

			for i := 0; i < len(subarray)-2; i++ {
				a, b, c := subarray[i], subarray[i+1], subarray[i+2]
				if b-a == c-b {
					count++
					break
				}
			}
		}
	}

	return count
}

func main() {

	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(countArithmeticSubarrays(n, arr))
}
```

## 5 задача
```go
package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	var s string
	fmt.Scan(&s)

	chars := []rune(s)

	open, close := 0, 0
	for _, c := range chars {
		if c == '(' {
			open++
		} else {
			close++
		}
	}

	cost := 0

	if open > n {
		diff := open - n
		cost += diff * b

		changed := 0
		for i := len(chars) - 1; i >= 0 && changed < diff; i-- {
			if chars[i] == '(' {
				chars[i] = ')'
				changed++
			}
		}
		close += diff
		open = n
	}

	if close > n {
		diff := close - n
		cost += diff * b

		changed := 0
		for i := 0; i < len(chars) && changed < diff; i++ {
			if chars[i] == ')' {
				chars[i] = '('
				changed++
			}
		}
		open += diff
		close = n
	}

	balance := 0
	stack := make([]int, 0)
	totalCost := cost

	for i := 0; i < len(chars); i++ {
		if chars[i] == '(' {
			balance++
			stack = append(stack, i)
		} else {
			balance--
		}

		if balance < 0 {
			if len(stack) > 0 && a < 2*b {

				lastOpen := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				chars[i], chars[lastOpen] = chars[lastOpen], chars[i]
				totalCost += a
				balance += 2
				i = lastOpen
			} else {
				chars[i] = '('
				totalCost += b
				balance += 2
			}
		}

		if balance == 0 && i == len(chars)-1 {
			break
		}
	}

	fmt.Println(totalCost)
}
```

## 6 задача
```go
package main

import (
	"fmt"
	"sort"
)

func maxHeightDifference(n int, heights []int) int {
	sort.Ints(heights)

	totalDifference := 0

	for i := 0; i < n/2; i++ {
		totalDifference += abs(heights[i] - heights[n-1-i])
	}

	return totalDifference
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n int
	fmt.Scan(&n)

	heights := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&heights[i])
	}

	fmt.Println(maxHeightDifference(n, heights))
}
```

## 7 задача
```go

```