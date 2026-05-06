# Pair Comparison — Find Max and Min

> Find the largest and smallest element in an array using the fewest comparisons possible.

---

## Problem Description

Given an array of `n` elements, find its **maximum** and **minimum** values while minimizing the total number of comparisons performed.

---

## Input

```
Input(Array[n], n)
```

| Parameter | Description           |
| --------- | --------------------- |
| `Array`   | Array of `n` integers |
| `n`       | Number of elements    |

---

## Expected Output

```
Output(Maximum, Minimum)
```

| Value     | Description                        |
| --------- | ---------------------------------- |
| `Maximum` | `max(Array[n])` — largest element  |
| `Minimum` | `min(Array[n])` — smallest element |

---

## Brute Force Solution

The naive approach iterates through the entire array once, performing **two comparisons per element** — one for the maximum and one for the minimum.

```go
func FindMaxMin(array []int, n int) (int, int) {
    if n == 1 {
        return array[0], array[0]
    }

    maximum := array[0]
    minimum := array[0]

    for i := 1; i < n; i++ {
        if maximum < array[i] {
            maximum = array[i]
        }
        if minimum > array[i] {
            minimum = array[i]
        }
    }

    return maximum, minimum
}
```

### Complexity

| Metric      | Value      |
| ----------- | ---------- |
| Time        | `O(n)`     |
| Space       | `O(1)`     |
| Comparisons | `2(n - 1)` |

---

## Optimized Solution — Pair Comparison

Instead of evaluating each element individually, this approach processes the array **two elements at a time**. The key insight is:

> Before comparing against the global max/min, first compare the two elements of the pair against **each other**. This way, only the **larger** candidate is tested against `maximum`, and only the **smaller** candidate is tested against `minimum` — eliminating one unnecessary comparison per pair.

```go
func FindMaxMin(array []int, n int) (int, int) {
    var maximum int
    var minimum int
    var i int

    if n%2 == 0 {
        if array[0] > array[1] {
            maximum = array[0]
            minimum = array[1]
        } else {
            maximum = array[1]
            minimum = array[0]
        }
        i = 2
    } else {
        maximum = array[0]
        minimum = array[0]
        i = 1
    }

    for i < n-1 {
        if array[i] > array[i+1] {
            if array[i] > maximum {
                maximum = array[i]
            }
            if array[i+1] < minimum {
                minimum = array[i+1]
            }
        } else {
            if array[i+1] > maximum {
                maximum = array[i+1]
            }
            if array[i] < minimum {
                minimum = array[i]
            }
        }
        i += 2
    }

    return maximum, minimum
}
```

### Why is this better?

In the brute force approach, each element requires **2 comparisons** (one against `maximum`, one against `minimum`), totaling `2(n-1)` comparisons.

In the pair comparison approach, every **2 elements** require only **3 comparisons**:

1. Compare the two elements against each other → `1 comparison`
2. Compare the larger one against `maximum` → `1 comparison`
3. Compare the smaller one against `minimum` → `1 comparison`
   This gives `3` comparisons per pair of `2` elements, or **1.5 comparisons per element** — a 25% reduction over the brute force.

### Concrete example with n = 8

| Approach        | Comparisons          |
| --------------- | -------------------- |
| Brute Force     | `2(8 - 1) = 14`      |
| Pair Comparison | `⌈3 × 8/2⌉ - 2 = 10` |

### Complexity

| Metric      | Brute Force | Pair Comparison |
| ----------- | ----------- | --------------- |
| Time        | `O(n)`      | `O(n)`          |
| Space       | `O(1)`      | `O(1)`          |
| Comparisons | `2(n - 1)`  | `⌈3n/2⌉ - 2`    |

Both algorithms have the same asymptotic time and space complexity, but the pair comparison approach makes **~25% fewer comparisons** in practice — a meaningful gain for large arrays or expensive comparison operations.
